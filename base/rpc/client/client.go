package main

import (
	"context"
	"log"

	pb "./test.pb.go" // 根据实际路径修改

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	req := &pb.HelloRequest{Name: "Alice"}

	resp, err := client.SayHello(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}

	message := resp.GetMessage()
	log.Printf("Server response: %s", message)
}
