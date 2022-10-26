package main

import (
	"base_project/grpc2/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// grpc 模拟客户端
func main() {
	// 1 创建客户端连接
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("grpc Dial err: ", err)
		return
	}
	defer conn.Close()
	// 2 创建远程服务的客户端
	grpcClient := pb.NewOrderServiceClient(conn)

	req := pb.OrderListRequest{OrderId: "aaa"}
	// 3 发送 rpc 请求
	res, err := grpcClient.GetList(context.Background(), &req)
	if err != nil {
		fmt.Println("grpcClient GetList err: ", err)
		return
	}
	fmt.Println(res)

}
