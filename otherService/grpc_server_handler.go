package main

import (
	"context"
	"log"

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
func (h *grpcHandler) GetMessage(c context.Context, p *pb.MessageRequest) (*pb.MessageResponse, error) {
	log.Printf("Get message %v",p)
	
	Response:= &pb.MessageResponse{
		Response: "Get The response",
		Request: &pb.MessageRequest{
			Name: "Ronex Ondimu",
		Message: "Get the job done alright",

		},
		
	}
	return Response, nil
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