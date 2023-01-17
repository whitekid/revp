package server

import (
	"context"

	pb "revp/pb/v1alpha1"
)

type greeterServerImpl struct {
	pb.UnimplementedGreeterServer
}

func (s *greeterServerImpl) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: "hello " + in.Name,
	}, nil
}
