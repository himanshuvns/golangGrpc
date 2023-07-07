package main

import (
	"log"

	pb "golangGrpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8006"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Message: []string{"Himanshu", "Ankit", "Gillu"},
	}
	//callSayHello(client)
	//callSayHelloServerStream(client, names)

	callSayHelloClientStream(client, names)
	defer conn.Close()
}
