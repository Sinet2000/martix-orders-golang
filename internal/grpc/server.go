package grpc

import (
	"log"
	"net"

	"github.com/Sinet2000/Martix-Orders-Go/internal/app/order"
	"github.com/Sinet2000/Martix-Orders-Go/proto"
	"google.golang.org/grpc"
)

func StartGRPCServer(orderUseCase order.OrderUseCase, address string) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterOrderServiceServer(grpcServer, order.NewOrderServiceServer(orderUseCase))

	log.Printf("gRPC server running on %s", address)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
