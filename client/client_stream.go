package main

import (
	"context"
	"log"
	"time"

	pb "github.com/rohans540/grpc-demo/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Client streaming started..\n")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}
	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending: %v\n", err)
		}
		log.Printf("Sent the request with name: %s\n", name)
		time.Sleep(2 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	log.Printf("Client streamingi finished..\n")
	if err != nil {
		log.Fatalf("Error while receiving %v", err)
	}
	log.Println(res.Messages)
}
