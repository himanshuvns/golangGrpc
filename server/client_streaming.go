package main

import (
	pb "golangGrpc/proto"
	"io"
	"log"
)

func (s *helloServer) sayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Message: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("Got Request with name %v", req.Message)
		messages = append(messages, "Hello ", req.Message)
	}
}
