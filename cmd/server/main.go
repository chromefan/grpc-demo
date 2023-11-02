package main

import (
	"grpc-demo/internal/server"
	"log"
)

func main() {
	if err := server.RunGRPCServer(); err != nil {
		log.Fatalf("failed to run gRPC server: %v", err)
	}

	if err := server.RunHTTPServer(); err != nil {
		log.Fatalf("failed to run HTTP server: %v", err)
	}
}
