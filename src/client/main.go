package main

import (
	"context"
	"log"

	pb "github.com/apalyukha/go-gRPC/src/client/pb/books"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewInventoryClient(conn)
	bookList, err := client.GetBookList(context.Background(), &pb.GetBookListRequest{})
	if err != nil {
		log.Fatalf("failed to get book list: %v", err)
	}
	log.Printf("book list: %v", bookList)
}
