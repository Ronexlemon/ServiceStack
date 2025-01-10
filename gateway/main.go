package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	common "github.com/ronexlemon/commons"
	pb "github.com/ronexlemon/commons/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	httpAddr            = common.EnvString("HTTP_ADDR", ":9090")
	orderServiceAddress = "localhost:2000"
)

func main() {
	conn, err := grpc.Dial(orderServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("gRPC connected to order service", orderServiceAddress)
	c := pb.NewOrderServiceClient(conn)
	mux := http.NewServeMux()
	handler := NewHandler(c)
	handler.registerRoute(mux)
	log.Printf("Startig HTTP server at %s", httpAddr)
	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("failed to start server")
	}

}
