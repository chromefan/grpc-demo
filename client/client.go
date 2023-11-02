package main

import (
	"context"
	"log"

	"grpc-demo/proto/helloworld" // 请根据实际的包路径导入生成的代码

	"google.golang.org/grpc"
)

func main() {
	// 建立到 gRPC 服务端的连接
	conn, err := grpc.Dial("localhost:8890", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("无法连接到服务器: %v", err)
	}
	defer conn.Close()

	// 创建一个 Greeter 客户端
	client := helloworld.NewGreeterClient(conn)

	// 调用 SayHello 方法
	request := &helloworld.HelloRequest{
		Name: "World",
	}
	response, err := client.SayHello(context.Background(), request)
	if err != nil {
		log.Fatalf("调用 SayHello 方法失败: %v", err)
	}

	// 打印服务端的响应
	log.Printf("服务端响应: %s", response.Message)
}
