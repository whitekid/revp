package client

import (
	"context"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"github.com/whitekid/goxp/log"

	pb "revp/pb/v1alpha1"
)

func TestServerSideStream(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, client, err := newTestServerAndClient(ctx, t, func(e *echo.Echo) {
	})
	require.NoError(t, err)
	defer func() { client.Close() }()

	stream, err := client.streamExample.ServerStream(ctx, &pb.StreamReq{
		Count: 4,
	})
	require.NoError(t, err)

	resps := []string{}
	for {
		data, err := stream.Recv()
		if err == io.EOF {
			break
		}
		require.NoError(t, err)

		resps = append(resps, data.Data)
	}
	require.Equal(t, "data 0,data 1,data 2,data 3", strings.Join(resps, ","))
}

func TestClientSideStream(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, client, err := newTestServerAndClient(ctx, t, func(e *echo.Echo) {
	})
	require.NoError(t, err)
	defer func() { client.Close() }()

	stream, err := client.streamExample.ClientStream(ctx)
	require.NoError(t, err)

	for i := 0; i < 4; i++ {
		err := stream.Send(&pb.StreamExampleData{Data: fmt.Sprintf("data %d", i)})
		require.NoError(t, err)
	}
	summary, err := stream.CloseAndRecv()
	require.NoError(t, err)
	require.Equal(t, "data 0,data 1,data 2,data 3", summary.Summary)
}

func TestBidirectionalStream(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, client, err := newTestServerAndClient(ctx, t, func(e *echo.Echo) {
	})
	require.NoError(t, err)
	defer func() { client.Close() }()

	stream, err := client.streamExample.BidirectionalStream(ctx)
	require.NoError(t, err)

	results := []string{}
	waitC, waitCancel := context.WithCancel(ctx)
	defer waitCancel()
	go func() {
		defer waitCancel()
		for {
			data, err := stream.Recv()
			if err == io.EOF {
				log.Debugf("EOF")
				return
			}
			if err != nil {
				log.Debugf("client err: %+v", err)
				continue
			}

			results = append(results, data.Data)
			fmt.Printf("%s", data.Data)
		}
	}()

	// send messages
	for i := 0; i < 4; i++ {
		err := stream.Send(&pb.StreamExampleData{Data: fmt.Sprintf("%d", i)})
		require.NoError(t, err)
	}
	require.NoError(t, stream.CloseSend())

	<-waitC.Done()

	require.Equal(t, "data 0,data 1,data 2,data 3", strings.Join(results, ","))
}
