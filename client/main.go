package main

import (
	"log"

	pb "github.com/rohans540/grpc-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v\n", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NameList{
		Names: []string{"Alice", "Bob", "Chris", "Daniel", "Emma", "Fiona", "Gregory"},
	}

	//Unary api call
	// callSayHello(client)

	//Server Streaming call
	// callSayHelloServerStream(client, names)

	//Client streaming call
	// callSayHelloClientStream(client, names)

	//Bi-directional streaming call
	callSayHelloBidirectionalStream(client, names)
}
