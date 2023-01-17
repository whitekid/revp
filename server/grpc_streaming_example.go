package server

import (
	"fmt"
	"io"
	"strings"

	"github.com/whitekid/goxp/log"

	pb "revp/pb/v1alpha1"
)

type streamExampleServerImpl struct {
	pb.UnimplementedStreamExampleServer
}

func (s *streamExampleServerImpl) ClientStream(stream pb.StreamExample_ClientStreamServer) error {
	summary := []string{}

	for {
		data, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.StreamExampleSummary{
				Summary: strings.Join(summary, ","),
			})
		}
		if err != nil {
			log.Errorf("server error: %+v", err)
			continue
		}
		summary = append(summary, string(data.Data))
	}
}

func (s *streamExampleServerImpl) ServerStream(in *pb.StreamReq, stream pb.StreamExample_ServerStreamServer) error {
	for i := 0; i < int(in.Count); i++ {
		stream.Send(&pb.StreamExampleData{
			Data: fmt.Sprintf("data %d", i),
		})
	}

	return nil
}

func (s *streamExampleServerImpl) BidirectionalStream(stream pb.StreamExample_BidirectionalStreamServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			log.Debugf("server EOF")
			return nil
		}
		if err != nil {
			log.Debugf("server error: %+v", err)
			continue
		}

		log.Debugf("server got data: %s", in.Data)
		if err := stream.Send(&pb.StreamExampleData{
			Data: fmt.Sprintf("data %s", in.Data),
		}); err != nil {
			return err
		}
	}
}
