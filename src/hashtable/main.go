package main

import (
	"fmt"
	"in-memory-hash-table/src/hashtable/pb"
	"log"
	"net"
	"os"

	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	app := &cli.App{
		Name:  "yaks",
		Usage: "Start a single threaded in-memory hash table",
		Action: func(cCtx *cli.Context) error {
			port := cCtx.Args().Get(0)
			listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))

			if err != nil {
				log.Fatalf("Failed to start server %v", err)
			}

			tcpServer := grpc.NewServer()
			reflection.Register(tcpServer)
			pb.RegisterHashTableServer(tcpServer, &server{})

			log.Printf("Starting server on port %v \n", port)

			if err := tcpServer.Serve(listener); err != nil {
				log.Fatalf("Failed to serve: %v", err)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
