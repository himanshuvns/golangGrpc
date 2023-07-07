package main

import (
	"context"
	pb "golangGrpc/proto"
	"log"
	"time"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Failed to say hello: %v", err)
	}
	log.Printf("%s", res.Message)
}
