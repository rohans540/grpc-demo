package main

import (
	"context"

	pb "github.com/rohans540/grpc-demo/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello, World!",
	}, nil
}
