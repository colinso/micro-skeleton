package main

import (
	"Personal/micro-skeleton/internal/api/grpc"
	"context"
	"fmt"

	ggrpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	address := fmt.Sprintf("%s:%s", "localhost", "8080")
	conn, err := ggrpc.Dial(address, ggrpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(fmt.Sprintf("fail to dial: %v", err))
	}
	defer conn.Close()
	client := grpc.NewMicroSkeletonClient(conn)

	item := grpc.Item{
		Name:  "Name1",
		Price: 12.01,
	}
	i, err := client.CreateItem(context.Background(), &item)
	if err != nil {
		panic(err)
	}
	fmt.Println(i)
}
