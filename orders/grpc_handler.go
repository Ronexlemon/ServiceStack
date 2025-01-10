package main

import (
	"context"
	"log"

	pb "github.com/ronexlemon/commons/api"
	"google.golang.org/grpc"
)

type grpcHandler struct{
	pb.UnimplementedOrderServiceServer

}
func NewGRPCHandler(grpcServer *grpc.Server){
	handler := &grpcHandler{}
	pb.RegisterOrderServiceServer(grpcServer,handler)
	
}

func (h *grpcHandler) CreateOrder(context.Context, *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Println("New Order Received")
	O :=&pb.Order{
ID: "20",
	}
	return O,nil
}