package server

import (
	"context"
	"helloworld"
)

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
