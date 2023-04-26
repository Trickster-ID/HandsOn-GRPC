package main

import (
	"go-grpc/cmd/services"
	productPB "go-grpc/pb/product"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	netListen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen %v", err.Error())
	}

	grpcServer := grpc.NewServer()
	productService := services.ProductService{}
	productPB.RegisterProductServiceServer(grpcServer, &productService)

	log.Printf("Server Started at %v", netListen.Addr())
	if err := grpcServer.Serve(netListen); err != nil {
		log.Fatalf("Failed to serve %v", err.Error())
	}
}
