package main

import (
	"context"
	"in-memory-hash-table/src/hello/pb"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.GenericMessage) (*pb.GenericMessage, error) {
	return &pb.GenericMessage{Body: "Server says hello!"}, nil
}
