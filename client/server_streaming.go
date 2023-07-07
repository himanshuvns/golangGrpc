package main

import (
	"context"
	"fmt"
	pb "golangGrpc/proto"
	"io"
	"log"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	fmt.Println("Streaming Started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("unable to send the names %v", err)
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Unable to receive the message %v", err)
		}
		log.Println(message)
	}
	fmt.Println("Streaming finished")
}
