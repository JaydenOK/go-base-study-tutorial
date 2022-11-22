package main

import (
	"base_project/grpc2/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

//grpc 服务端

// 实现 Service接口
type OrderService struct {
}

// 使用*pb.路径
func (c *OrderService) GetList(ctx context.Context, req *pb.OrderListRequest) (*pb.OrderList, error) {
	if req.OrderId == "" {
		return nil, nil
	}
	//返回结果列表
	return &pb.OrderList{
		Id:           0,
		OrderId:      req.OrderId,
		ShipName:     "testShipName",
		PlatformCode: "testPlatformCode",
	}, nil
}

func main() {
	// 1 初始化 grpc 对象
	grpcServer := grpc.NewServer()
	// 2 注册服务
	pb.RegisterOrderServiceServer(grpcServer, new(OrderService))
	// 3 创建监听
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("net Listen err: ", err)
		return
	}
	defer listener.Close()
	// 4 绑定服务，启动监听
	grpcServer.Serve(listener)
}
