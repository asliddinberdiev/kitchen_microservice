package main

import (
	"log"
	"net"

	"github.com/asliddinberdiev/kitchen_microservice/services/orders/handler"
	"github.com/asliddinberdiev/kitchen_microservice/services/orders/service"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (svr *gRPCServer) Run() error {
	listen, err := net.Listen("tcp", svr.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	orderService := service.NewOrderService()
	handler.NewGrpcOrdersService(grpcServer, orderService)

	log.Println("starting gRPC server on: ", svr.addr)

	return grpcServer.Serve(listen)
}
