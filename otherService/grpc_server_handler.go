package main

import (
	"context"

	pb "github.com/ronexlemon/otherservice/api"
	"google.golang.org/grpc"
)
type grpcHandler struct{
	pb.UnimplementedRouteGuideServer

}

func NewGRPCHandler(grpcServer *grpc.Server){
	handler := &grpcHandler{}
	pb.RegisterRouteGuideServer(grpcServer,handler)

	
	
}
func (h *grpcHandler) GetMessage(context.Context, *pb.MessageRequest) (*pb.MessageResponse, error) {
	return nil, nil
}
func (h *grpcHandler) GetMessageLists(*pb.MessageKey, grpc.ServerStreamingServer[pb.Keys]) error {
	return nil
}
func (h *grpcHandler) AddMessage(grpc.ClientStreamingServer[pb.MessageRequest, pb.MessageResponse]) error {
	return nil
}
func (h *grpcHandler) AddMessageLists(grpc.BidiStreamingServer[pb.MessageKey, pb.Keys]) error {
	return nil
}