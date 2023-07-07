package main

import (
	"context"
	pb "golangGrpc/proto"
	"log"
	"time"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Streaming has started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names %v", err)
	}
	for _, name := range names.Message {
		req := &pb.HelloRequest{
			Message: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("could not send HelloRequest %v", err)
		}
		log.Printf("Sent the request with name %v", name)
		time.Sleep(2 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	log.Printf("Client streaming finished")
	if err != nil {
		log.Fatalf("Error while receiving the Response %v", err)
	}
	log.Println(res.Message)
}
