package server

import (
	"context"
	"fmt"
	"math/rand"

	"grpc-demo/proto/helloworld" // 请根据实际的包路径导入生成的代码
	"time"
)

type GreeterServer struct {
	helloworld.GreeterServer
}

func (s *GreeterServer) SayHello(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	// 获取客户端传递的 name 参数
	name := request.Name

	// 获取当前时间
	currentTime := time.Now().Format(time.RFC3339)

	// 构建响应消息
	message := fmt.Sprintf("hello %s, 现在的时间为 %s", name, currentTime)

	// 返回响应
	return &helloworld.HelloReply{
		Message: message,
	}, nil
}

func (s *GreeterServer) TestEndpoint(ctx context.Context, request *helloworld.TestRequest) (*helloworld.HelloReply, error) {
	// 获取客户端传递的 number 参数
	number := request.Number

	// 生成一个随机数，范围在 10 到 999 之间
	rand.Seed(time.Now().UnixNano())
	randomFactor := rand.Intn(990) + 10

	// 计算结果
	result := int(number) * randomFactor

	// 构建响应消息
	response := &helloworld.HelloReply{
		Message: fmt.Sprintf("Result: %d", result),
	}

	// 返回响应
	return response, nil
}
