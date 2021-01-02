package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/vireakouk/learn-grpc/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Printf("Greet RPC is being invoke with: %v", req)
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()
	result := "Hello " + firstName + " " + lastName

	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	fmt.Println("gRPC server is running")
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
