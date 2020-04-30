package main

import (
	"context"
	"fmt"
	"gofullstack/grpc-learning/greet/greetpd"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {}

func (*server) Greet(ctx context.Context, req *greetpd.GreetRequest) (*greetpd.GreetResponse, error) {
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &greetpd.GreetResponse{
		Result: result,
	}
	return res, nil
}

func main()  {
	fmt.Println("working")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatal("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpd.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}