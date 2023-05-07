package main

import (
	"context"
	"encoding/json"
	"fmt"
	productPB "go-grpc/pb/product"
	"log"

	"google.golang.org/grpc"
)

func main() {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(":50051", opts...)
	if err != nil {
		log.Fatal("error dial: ", err)
	}
	defer conn.Close()
	client := productPB.NewProductServiceClient(conn)
	GetAllProduct(client)
}

func GetAllProduct(client productPB.ProductServiceClient) {
	empty := &productPB.Empty{}
	products, err := client.GetProducts(context.Background(), empty)
	if err != nil {
		log.Fatal("error get products: ", err)
	}
	res, _ := json.MarshalIndent(products, "", "  ")
	fmt.Println(string(res))
}
