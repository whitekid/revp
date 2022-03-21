package server

import (
	"context"

	"github.com/whitekid/revp/pb"
)

type greeterServerImpl struct {
	pb.UnimplementedGreeterServer
}

func (s *greeterServerImpl) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: "hello " + in.Name,
	}, nil
}
