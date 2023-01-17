package client

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
	"github.com/whitekid/goxp"
	"github.com/whitekid/goxp/log"
	"github.com/whitekid/goxp/request"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"revp/config"
	"revp/server"
)

func TestMain(m *testing.M) {
	cmd := &cobra.Command{}
	config.Init("revp", cmd.Flags())

	cmd = &cobra.Command{}
	config.Client.Init("revp", cmd.Flags())

	cmd = &cobra.Command{}
	config.Server.Init("revps", cmd.Flags())

	os.Exit(m.Run())
}

// setup server and returns (localAddr, serverAddr)
func newTestServer(ctx context.Context, t *testing.T, routes ...func(*echo.Echo)) (string, string, io.Closer) {
	t.Parallel()

	// start test server
	remoteServerPort, err := goxp.AvailablePort()
	require.NoError(t, err)
	localPort, err := goxp.AvailablePort()
	require.NoError(t, err)

	remoteServerAddr := fmt.Sprintf("127.0.0.1:%d", remoteServerPort)
	localServerAddr := fmt.Sprintf("127.0.0.1:%d", localPort)

	go func() {
		server.New(remoteServerAddr).Serve(ctx)
	}()

	// start local services
	ln, err := net.Listen("tcp", localServerAddr)
	require.NoError(t, err)

	e := echo.New()
	e.HideBanner = true
	e.Listener = ln
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})
	e.GET("/echo", func(c echo.Context) error {
		return c.String(http.StatusOK, c.QueryParam("message"))
	})
	e.POST("/add", func(c echo.Context) error {
		x, err := strconv.ParseInt(c.FormValue("x"), 10, 64)
		if err != nil {
			return echo.ErrBadRequest
		}
		y, err := strconv.ParseInt(c.FormValue("y"), 10, 64)
		if err != nil {
			return echo.ErrBadRequest
		}

		return c.JSON(http.StatusOK, map[string]int{
			"result": int(x + y),
		})
	})

	for _, route := range routes {
		route(e)
	}

	go func() { e.Start("") }()

	go func() {
		<-ctx.Done()
		ln.Close()
	}()

	log.Debugf("starting server: local=%s, remote=%s", localServerAddr, remoteServerAddr)
	return localServerAddr, remoteServerAddr, ln
}

// setup and returns (local_server_closer, client, err)
func newTestServerAndClient(ctx context.Context, t *testing.T, routes ...func(*echo.Echo)) (io.Closer, *Client, error) {
	localAddr, serverAddr, localCloser := newTestServer(ctx, t, routes...)
	client, err := New(localAddr, serverAddr)
	return localCloser, client, err
}

func testPost(ctx context.Context, t *testing.T, remoteAddr string) {
	x := rand.Intn(100)
	y := rand.Intn(100)

	resp, err := request.Post(remoteAddr+"add").
		Form("x", strconv.FormatInt(int64(x), 10)).
		Form("y", strconv.FormatInt(int64(y), 10)).
		Do(ctx)
	require.NoError(t, err)
	require.True(t, resp.Success())
	response := map[string]int{}
	defer resp.Body.Close()
	require.NoError(t, resp.JSON(&response))
	require.Equal(t, x+y, response["result"])
}

func testGet(ctx context.Context, t *testing.T, remoteAddr string) {
	resp, err := request.Get(remoteAddr+"echo").
		Query("message", t.Name()).
		Do(ctx)
	require.NoError(t, err)
	require.True(t, resp.Success())
	require.Equal(t, t.Name(), resp.String())
}

func TestRun(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, client, err := newTestServerAndClient(ctx, t)
	require.NoError(t, err)
	defer func() { client.Close() }()

	remoteAddr, err := client.Run(ctx, nil)
	require.NoError(t, err)

	t.Run("GET", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			t.Run(goxp.RandomString(10), func(t *testing.T) {
				testGet(ctx, t, remoteAddr)
			})
		}
	})

	t.Run("POST", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			t.Run("", func(t *testing.T) { testPost(ctx, t, remoteAddr) })
		}
	})
}

func TestRunPost(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, client, err := newTestServerAndClient(ctx, t)
	require.NoError(t, err)
	defer func() { client.Close() }()

	remoteAddr, err := client.Run(ctx, nil)
	require.NoError(t, err)

	testPost(ctx, t, remoteAddr)
}

// local server가 닫혔을 대 처리 bad gateway
func TestLocalDisconnect(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	localCloser, client, err := newTestServerAndClient(ctx, t)
	require.NoError(t, err)
	defer func() { client.Close() }()

	remoteAddr, err := client.Run(ctx, nil)
	require.NoError(t, err)

	localCloser.Close()

	resp, err := request.Get(remoteAddr).Do(ctx)
	require.NoError(t, err)
	require.Equal(t, http.StatusBadGateway, resp.StatusCode)
}

func TestSecret(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, client, err := newTestServerAndClient(ctx, t)
	require.NoError(t, err)
	defer func() { client.Close() }()

	t.Run("valid", func(t *testing.T) {
		_, err = client.Run(ctx, nil)
		require.NoError(t, err)
	})

	t.Run("demo", func(t *testing.T) {
		ctx = context.WithValue(ctx, ctxKeySecret, config.Server.DemoSecret())
		_, err = client.Run(ctx, nil)
		require.NoError(t, err)
	})

	t.Run("invalid", func(t *testing.T) {
		ctx = context.WithValue(ctx, ctxKeySecret, "invalid-secret")
		_, err = client.Run(ctx, nil)
		require.Error(t, err)

		e := status.Error(codes.Unauthenticated, "")
		require.ErrorAs(t, err, &e)
	})
}

func TestDemoTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, client, err := newTestServerAndClient(ctx, t)
	require.NoError(t, err)
	defer func() { client.Close() }()

	os.Setenv("RP_DEMO_TIMEOUT", "1s")
	ctx = context.WithValue(ctx, ctxKeySecret, config.Server.DemoSecret())
	ctxClose, cancel := context.WithCancel(context.Background())
	_, err = client.Run(ctx, cancel)

	select {
	case <-ctx.Done():
		require.Fail(t, ctx.Err().Error())
	case <-ctxClose.Done():
		// closeC로 종료되어야지..
	}
}
