package server

import (
	"context"
	"log"
	"net"
	"net/http"

	"helloworld"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunGRPCServer() error {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	helloworld.RegisterGreeterServer(grpcServer, &GreeterHandler{})
	reflection.Register(grpcServer)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve gRPC server: %v", err)
		}
	}()

	return nil
}

func RunHTTPServer() error {
	mux := http.NewServeMux()
	mux.Handle("/", grpcHandlerFunc(grpcServer))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		return err
	}

	return nil
}

func grpcHandlerFunc(grpcServer *grpc.Server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		grpcServer.ServeHTTP(w, r)
	})
}

type GreeterHandler struct{}

func (h *GreeterHandler) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloResponse, error) {
	name := req.GetName()
	if name == "" {
		name = "World"
	}
	resp := &helloworld.HelloResponse{
		Greeting: "Hello, " + name + "!",
	}
	return resp, nil
}
