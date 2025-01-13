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

func (h *grpcHandler) CreateOrder(c context.Context,p *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Printf("New Order Received! Order %v",p)
	O :=&pb.Order{
ID: "20",

	}
	return O,nil
}