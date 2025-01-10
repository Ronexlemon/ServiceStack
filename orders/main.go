package main

import (
	"context"
	"log"
	"net"
common "github.com/ronexlemon/commons"
	"google.golang.org/grpc"
	
)
var (
	grpcAdd = common.EnvString("GRPC_ADDR", "localhost:2000")
)

func main(){
	grpcServer := grpc.NewServer()
	l,err := net.Listen("tcp",grpcAdd)
	if err !=nil{
		log.Fatal(err)
	}
	defer l.Close()
	str:=NewStore()
	svc :=NewService(str)
	NewGRPCHandler(grpcServer)
	svc.CreateOrder(context.Background())
	log.Println("GRPC SERVER START AT",grpcAdd)
	if err := grpcServer.Serve(l);err!=nil{
		log.Fatal(err.Error())
	}

}