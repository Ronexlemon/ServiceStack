package main

import (
	"log"
"net/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/ronexlemon/otherservice/api"
)


var (
	grpcServerAddr = "localhost:2000"
	httpAddr = ":8181"
)

func main(){
	conn,err:= grpc.Dial(grpcServerAddr,grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil{
		log.Fatal(err)}
		defer conn.Close()
		log.Println("gRPC connected to order service", grpcServerAddr)
		c:= pb.NewRouteGuideClient(conn)
		mux:= http.NewServeMux()
		handler:= NewHandler(c)
		handler.registerRoute(mux)
		log.Println("http server listening on port",httpAddr)
		if err:= http.ListenAndServe(httpAddr,mux); err !=nil{
			log.Fatal(err)
		}


}