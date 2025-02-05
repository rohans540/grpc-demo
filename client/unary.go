package main

import (
	"context"
	"log"
	"time"

	pb "github.com/rohans540/grpc-demo/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}
	log.Printf("%s", res.Message)
}
