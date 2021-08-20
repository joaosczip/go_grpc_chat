package main

import (
	"log"
	"net"

	"github.com/joaosczip/grpc_chat/pkg/pb"
	"github.com/joaosczip/grpc_chat/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:5051")
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer, services.NewChatService())

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("Could not serve gRPC server: %v", err)
	}
}
