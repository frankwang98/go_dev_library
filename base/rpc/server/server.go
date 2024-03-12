package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "examplepb/test.pb.go" // 根据实际路径修改
)

type greeterServer struct{}

func (s *greeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	name := req.GetName()
	message := fmt.Sprintf("Hello, %s!", name)
	return &pb.HelloResponse{Message: message}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &greeterServer{})

	log.Println("Server started on port 50051")
	if err := server.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
