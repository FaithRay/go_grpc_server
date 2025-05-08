package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "go_server/pb/trade"
)

const (
	defaultAddress = "127.0.0.1:50052" // 默认监听端口
)

type server struct {
	pb.UnimplementedTradeServiceServer
}

func (s *server) GetTrade(ctx context.Context, req *pb.TradeRequest) (*pb.TradeListResponse, error) {
	log.Println("Received GetTrade request")
	// 示例响应
	return &pb.TradeListResponse{
		// 填充响应字段
	}, nil
}

func (s *server) GetOrder(ctx context.Context, req *pb.OrderRequest) (*pb.OrderListResponse, error) {
	log.Println("Received GetOrder request")
	return &pb.OrderListResponse{
		// 填充响应字段

	}, nil
}

func StartGrpcServer(address string) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTradeServiceServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	log.Println("Starting gRPC server on", defaultAddress)
	StartGrpcServer(defaultAddress)
	log.Println("gRPC server stopped")
}
