package client

import (
	"bufio"
	"context"
	"io"
	"net"
	"net/http"
	"net/textproto"

	"revp/config"
	"revp/pb"

	"github.com/pkg/errors"
	"github.com/whitekid/goxp/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/keepalive"
)

// Client reverse tunnel client
//
//  (localAddr) ----<grpc>---- (serverAddr) --------<https>---- (remoteAddr)
//  http://127.0.0.1:8080      grpc//recp.woosum.net:49999      https://recp.woosum.net:599999
type Client struct {
	serverAddr string
	localAddr  string

	conn *grpc.ClientConn

	revp pb.RevpClient

	streamExample pb.StreamExampleClient
	greeter       pb.GreeterClient
}

func New(localAddr string, serverAddr string) (*Client, error) {
	client := &Client{
		localAddr:  localAddr,
		serverAddr: serverAddr,
	}

	var kacp = keepalive.ClientParameters{
		Time:                config.Client.KeepaliveTime(),    // send pings every 10 seconds if there is no activity
		Timeout:             config.Client.KeepaliveTimeout(), // wait 1 second for ping ack before considering the connection dead
		PermitWithoutStream: true,                             // send pings even without active streams
	}

	// localhost:49999
	log.Debugf("connecting %s....", serverAddr)
	conn, err := grpc.Dial(serverAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(kacp),
	)
	if err != nil {
		return nil, errors.Wrap(err, "connect failed")
	}
	grpc.UseCompressor(gzip.Name)

	client.conn = conn
	client.revp = pb.NewRevpClient(conn)
	client.greeter = pb.NewGreeterClient(conn)
	client.streamExample = pb.NewStreamExampleClient(conn)

	return client, nil
}

func (c *Client) Close() error { return c.conn.Close() }

const (
	ctxKeySecret = "ctxKeySecret:ab8918d9-fe61-4d7c-ab97-63a332a0fe28"
)

// Run handshake with server and run proxy in goroutine
// run run stream in gorotuine and returns remoteAddr
//
// TODO cancel을 두는 것은 영 맘에 안듬. goroutine으로 proxy가 돌아가고,
// 종료시점을 잡기 위해서 stopped가 필요함
func (c *Client) Run(ctx context.Context, stopped context.CancelFunc) (string, error) {
	// start streaming
	log.Debug("start streaming...")
	stream, err := c.revp.Stream(ctx)
	if err != nil {
		return "", errors.Wrap(err, "fail to start streaming...")
	}
	log.Debug("straming started....")

	secret := config.Secret()
	if v, ok := ctx.Value(ctxKeySecret).(string); ok {
		secret = v
	}

	log.Debug("start handshaking...")
	if err := stream.Send(&pb.StreamData{Secret: &secret}); err != nil {
		log.Fatal(err)
	}
	log.Debug("start handshaking... done")

	// receive server address
	log.Debug("waiting...proxing")
	data, err := stream.Recv()
	if err != nil {
		return "", errors.Wrap(err, "fail to receive remote server address")
	}

	remoteAddr := string(data.Data)
	log.Debugf("get server address: %s", remoteAddr)

	go func() {
		defer stream.CloseSend()

		for {
			select {
			case <-stream.Context().Done():
				log.Debugf("stream disconnected")
				if stopped != nil {
					stopped()
				}
				return
			default:
			}

			req, err := http.ReadRequest(bufio.NewReader(newStreamReader(stream)))
			if err != nil {
				log.Debugf("ReadRequest failed: %v", err)
				continue
			}

			// connect local server
			log.Debugf("connecting local server: %s", c.localAddr)
			localConn, err := net.Dial("tcp", c.localAddr)
			if err != nil {
				log.Errorf("fail to connect local server: %s", err)
				err := pb.StreamData_EOF
				stream.Send(&pb.StreamData{Err: &err}) // bad gateway

				continue
			}

			// headers
			localWriter := textproto.NewWriter(bufio.NewWriter(localConn))
			localWriter.PrintfLine("%s %s HTTP/1.1", req.Method, req.URL.String())
			localWriter.PrintfLine("Host: %s", req.Host)
			for k, headers := range req.Header {
				for _, h := range headers {
					localWriter.PrintfLine("%s: %s", k, h)
				}
			}
			localWriter.PrintfLine("")

			// body
			if req.Body != nil {
				if _, err := io.CopyBuffer(localConn, req.Body, make([]byte, 4096)); err != nil {
					log.Errorf("%+v", err)
					pbErr := pb.StreamData_EOF
					stream.Send(&pb.StreamData{Err: &pbErr})
					continue
				}
				defer req.Body.Close()
			}

			// read response
			remoteSW := newStreamWriter(stream)
			resp, err := http.ReadResponse(bufio.NewReader(localConn), req)
			if err != nil {
				log.Errorf("%+v", err)
				pbErr := pb.StreamData_EOF
				stream.Send(&pb.StreamData{Err: &pbErr})
				continue
			}
			defer resp.Body.Close()

			remoteTW := textproto.NewWriter(bufio.NewWriter(remoteSW))
			remoteTW.PrintfLine("HTTP/1.1 %s", resp.Status)
			for k, headers := range resp.Header {
				for _, h := range headers {
					remoteTW.PrintfLine("%s: %s", k, h)
				}
			}
			remoteTW.PrintfLine("")

			if _, err := io.CopyBuffer(remoteSW, resp.Body, make([]byte, 4096)); err != nil {
				log.Errorf("%+v", err)
				pbErr := pb.StreamData_EOF
				stream.Send(&pb.StreamData{Err: &pbErr})

				continue
			}

			resp.Body.Close()
			localConn.Close()
		}
	}()

	return remoteAddr, nil
}

type streamReader struct {
	stream pb.Revp_StreamClient
}

func newStreamReader(stream pb.Revp_StreamClient) io.Reader {
	return &streamReader{
		stream: stream,
	}
}

func (r *streamReader) Read(p []byte) (n int, err error) {
	data, err := r.stream.Recv()
	if err != nil {
		return 0, errors.Wrapf(err, "stream read failed with %+v", err)
	}

	copy(p, data.Data)
	return len(data.Data), nil
}

type streamWriter struct {
	stream pb.Revp_StreamClient
}

func newStreamWriter(stream pb.Revp_StreamClient) io.Writer {
	return &streamWriter{
		stream: stream,
	}
}

func (r *streamWriter) Write(p []byte) (n int, err error) {
	if err := r.stream.Send(&pb.StreamData{Data: p}); err != nil {
		return 0, errors.Wrap(err, "send failed")
	}

	return len(p), nil
}
