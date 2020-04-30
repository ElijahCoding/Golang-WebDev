package main

import (
	"fmt"
	"gofullstack/grpc-learning/greet/greetpd"
	"google.golang.org/grpc"
	"log"
)

func main()  {
	fmt.Println("client side")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer cc.Close()

	c := greetpd.NewGreetServiceClient(cc)
	fmt.Println("Created client")
}
