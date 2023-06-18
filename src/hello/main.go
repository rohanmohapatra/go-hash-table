package main

import (
	"in-memory-hash-table/src/hello/pb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}

	tcpServer := grpc.NewServer()
	reflection.Register(tcpServer)
	pb.RegisterHelloServiceServer(tcpServer, &server{})

	log.Println("Starting server on port 8080")

	if err := tcpServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
