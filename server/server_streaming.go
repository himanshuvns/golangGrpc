package main

import (
	pb "golangGrpc/proto"
	"log"
	"time"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("listed names are %v", req.Message)
	for _, m := range req.Message {
		res := &pb.HelloResponse{
			Message: "Hello " + m,
		}
		if err := stream.Send(res); err != nil {
			log.Fatalf("send error: %v", err)
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
