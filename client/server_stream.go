package main

import (
	"context"
	"io"
	"log"

	pb "github.com/rohans540/grpc-demo/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Streaming has started..")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("Could not send names: %v\n", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while streaming: %v\n", err)
		}
		log.Println(message)
	}
	log.Printf("Streaming finished...\n")

}
