package services

import (
	"fmt"
	"io"
	"log"

	"github.com/joaosczip/grpc_chat/pkg/pb"
)

type ChatService struct {
	pb.UnimplementedChatServiceServer
}

func NewChatService() *ChatService {
	return &ChatService{}
}

func (*ChatService) StartChat(stream pb.ChatService_StartChatServer) error {
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error receiving the message: %v", err)
			return err
		}

		fmt.Printf("Receive message %s from %s\n", req.GetText(), req.GetFrom())

		err = stream.Send(&pb.Message{
			From: req.GetFrom(),
			Text: req.GetText(),
		})

		if err != nil {
			log.Fatalf("Error sending the message to the client: %v", err)
			return err
		}
	}
}
