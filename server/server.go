package server

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
	"net/http/httputil"
	"net/textproto"
	"time"

	"github.com/flosch/pongo2/v5"
	"github.com/pkg/errors"
	"github.com/whitekid/goxp/log"
	"github.com/whitekid/goxp/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	"revp/config"
	pb "revp/pb/v1alpha1"
)

func New(serverAddr string) service.Interface {
	return &serverImpl{
		serverAddr: serverAddr,
	}
}

type serverImpl struct {
	serverAddr string

	pb.UnimplementedRevpServer

	pb.UnimplementedStreamExampleServer
	pb.UnimplementedGreeterServer
}

func (s *serverImpl) Serve(ctx context.Context) error {
	ln, err := net.Listen("tcp", s.serverAddr)
	if err != nil {
		return errors.Wrapf(err, "listen failed")
	}
	log.Infof("starting server %s", ln.Addr().String())

	// grpc server
	var kaep = keepalive.EnforcementPolicy{
		MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
		PermitWithoutStream: true,            // Allow pings even when there are no active streams
	}

	// MaxConnectionIdle, MaxConnectionAge, MaxConnectionAgeGrace 등등을 설정하는 경우 스트림이 IDLE 상태가 계속되면
	// unavailable 오류가 떨어진다.
	// 주의) stream이 LB의 세션을 계속 잡고 있어서 LB를 사용하는 환경이라면 연결이 끊어질 수 있어서 이거 처리하는 게 필요함
	//
	// MaxConnectionIdle: 15 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
	// MaxConnectionAge:  30 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
	// MaxConnectionAgeGrace: 5 * time.Second,  // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
	var kasp = keepalive.ServerParameters{
		Time:    config.Server.KeepaliveTime(),    // Ping the client if it is idle for 5 seconds to ensure the connection is still active
		Timeout: config.Server.KeepaliveTimeout(), // Wait 1 second for the ping ack before assuming the connection is dead
	}
	log.Debugf("server keepalive time = %s", config.Server.KeepaliveTime())
	log.Debugf("server keepalive timeout = %s", config.Client.KeepaliveTimeout())
	g := grpc.NewServer(grpc.KeepaliveEnforcementPolicy(kaep), grpc.KeepaliveParams(kasp))

	pb.RegisterRevpServer(g, s)
	pb.RegisterGreeterServer(g, &greeterServerImpl{})
	pb.RegisterStreamExampleServer(g, &streamExampleServerImpl{})

	go func() {
		<-ctx.Done()
		ln.Close()
	}()

	if err := g.Serve(ln); err != nil {
		return errors.Wrapf(err, "serve failed")
	}

	return nil
}

func (s *serverImpl) Stream(stream pb.Revp_StreamServer) error {
	peer, ok := peer.FromContext(stream.Context())
	if !ok {
		log.Fatal("fail to get peer info")
	}
	log.Debugf("start streaming: peer=%s", peer.Addr.String())

	// handshaking
	data, err := stream.Recv()
	if err != nil {
		return errors.Wrapf(err, "handshake failed")
	}

	if data.Secret != nil && !(*data.Secret == config.Secret() || *data.Secret == config.Server.DemoSecret()) {
		return status.Error(codes.Unauthenticated, "invalid secret")
	}

	demoTimeoutC := make(<-chan time.Time)
	if *data.Secret == config.Server.DemoSecret() {
		log.Debugf("with demo timeout: %s", config.Server.DemoTimeout())
		demoTimeoutC = time.After(config.Server.DemoTimeout())
	}

	// start proxing
	ln, err := allocatePort()
	if err != nil {
		return errors.Wrap(err, "server listen failed")
	}
	defer ln.Close()

	port := ln.Addr().(*net.TCPAddr).Port

	serverAddr, err := config.Server.RootURL().Execute(pongo2.Context{"port": port})
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	log.Debugf("send remote address as: %s", serverAddr)
	if err := stream.Send(&pb.StreamData{Data: []byte(serverAddr)}); err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := http.Serve(ln, newStreamProxy(stream)); err != http.ErrServerClosed {
			log.Errorf("%s", err) // io.EOF on client disconnect
		}
	}()

	select {
	case <-demoTimeoutC:
		return status.Error(codes.DeadlineExceeded, "demo timeout")
	case <-stream.Context().Done():
	}

	return nil
}

func allocatePort() (net.Listener, error) {
	portRange := config.Server.PortRange()

	for i := 0; i < 1000; {
		port := portRange[0] + rand.Intn(portRange[1]-portRange[0])

		ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			return nil, errors.Wrap(err, "server listen failed")
		}

		if port >= portRange[0] && port <= portRange[1] {
			return ln, nil
		}

		ln.Close()
		i++
	}

	return nil, errors.New("no available port")
}

type proxyRoundtripper struct {
	stream pb.Revp_StreamServer
}

func (rt *proxyRoundtripper) RoundTrip(req *http.Request) (*http.Response, error) {
	sw := newStreamWriter(rt.stream)
	tw := textproto.NewWriter(bufio.NewWriter(sw))

	tw.PrintfLine("%s %s HTTP/1.1", req.Method, req.URL.String())
	tw.PrintfLine("Host: %s", req.Host)
	for k, headers := range req.Header {
		for _, header := range headers {
			log.Debugf("header %s: %s", k, header)
			tw.PrintfLine("%s: %s", k, header)
		}
	}
	tw.PrintfLine("")

	if req.Body != nil {
		if _, err := io.CopyBuffer(sw, req.Body, make([]byte, 4096)); err != nil {
			log.Fatalf("%s: %+v", err, err)
		}
	}

	r := bufio.NewReader(newStreamReader(rt.stream))
	resp, err := http.ReadResponse(r, req)
	if err != nil {
		if errors.Is(err, io.EOF) { // server disconnected, will get bad geteway error
			return nil, io.EOF
		}
		return nil, errors.Wrapf(err, "read failed")
	}

	return resp, nil
}

// newStreamProxy takes target host and creates a reverse proxy
func newStreamProxy(stream pb.Revp_StreamServer) *httputil.ReverseProxy {
	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {},
		Transport: &proxyRoundtripper{
			stream: stream,
		},
		ModifyResponse: func(resp *http.Response) error {
			return nil
		},
		ErrorLog: log.NewGoLogger(log.DebugLevel),
	}

	return proxy
}

type streamReader struct {
	stream pb.Revp_StreamServer
}

func newStreamReader(stream pb.Revp_StreamServer) io.Reader {
	return &streamReader{
		stream: stream,
	}
}

func (r *streamReader) Read(p []byte) (n int, err error) {
	data, err := r.stream.Recv()
	if err != nil {
		return 0, errors.Wrap(err, "recv failed")
	}

	if data.Data == nil && data.Err != nil && *data.Err != pb.StreamData_NO_ERROR {
		switch *data.Err {
		case pb.StreamData_EOF:
			return 0, io.EOF
		default:
			log.Fatal("unhandled error: %s", data.Err.String())
		}
	}

	copy(p, data.Data)
	return len(data.Data), nil
}

type streamWriter struct {
	stream pb.Revp_StreamServer
}

func newStreamWriter(stream pb.Revp_StreamServer) io.Writer {
	return &streamWriter{
		stream: stream,
	}
}

func (r *streamWriter) Write(p []byte) (n int, err error) {
	if err := r.stream.Send(&pb.StreamData{
		Data: p,
	}); err != nil {
		return 0, errors.Wrap(err, "send failed")
	}

	return len(p), nil
}
