package main

import (
	"fmt"
	"gofullstack/grpc-learning/greet/greetpd"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {}

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