package main

import (
	"context"
	"fmt"
	"log"

	"github.com/vireakouk/learn-grpc/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("I'm a gRPC client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	//fmt.Printf("Client created: %f", c)

	unaryGreet(c)
}

func unaryGreet(c greetpb.GreetServiceClient) {
	log.Printf("Starting Unary RPC call...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Vireak",
			LastName:  "Ouk",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet RPC: %v", res)
}
