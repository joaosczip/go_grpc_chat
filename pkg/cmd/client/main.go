package main

import (
	"bufio"
	"context"
	"log"
	"os"

	"github.com/joaosczip/grpc_chat/pkg/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}

	defer conn.Close()

	client := pb.NewChatServiceClient(conn)
	reader := bufio.NewReader(os.Stdin)

	StartChat(client, reader)
}

func StartChat(client pb.ChatServiceClient, reader *bufio.Reader) {
	stream, err := client.StartChat(context.Background())
	if err != nil {
		log.Fatalf("Error starting chat: %v", err)
	}

	for {
		input, _ := reader.ReadString('\n')

		err = stream.Send(&pb.Message{
			From: "someone",
			Text: input,
		})
		if err != nil {
			log.Fatalf("Failed to send the message to the server: %v", err)
		}
	}
}
