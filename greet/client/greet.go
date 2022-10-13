package main

import (
	"context"
	"log"

	pb "github.com/javiergomezve-grpc-course/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoke")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Javier",
	})
	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
