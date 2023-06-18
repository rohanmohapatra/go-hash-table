package main

import (
	"context"
	"in-memory-hash-table/src/hashtable/pb"
	"in-memory-hash-table/src/hashtable/storage"
)

type server struct {
	pb.UnimplementedHashTableServer
}

func (s *server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	rc, value := storage.Get(in.GetKey())
	return &pb.GetResponse{Rc: rc, Value: value}, nil
}

func (s *server) Put(ctx context.Context, in *pb.PutRequest) (*pb.PutResponse, error) {
	storage.Put(in.GetKey(), in.GetValue())
	return &pb.PutResponse{Rc: 0}, nil
}
