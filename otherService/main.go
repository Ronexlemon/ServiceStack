package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
)
var (
	grpcAddr = "localhost:2000"
)

func main(){
	// grpc server
	grpcServer := grpc.NewServer()
	conn,err :=net.Listen("tcp",grpcAddr); 
	if err !=nil{
		log.Fatal(err)
	}
	defer conn.Close()
	NewGRPCHandler(grpcServer)
	
	log.Println("GRPC SERVER START AT",grpcAddr)
	if err := grpcServer.Serve(conn); err !=nil{
		log.Fatal(err)
	}


}