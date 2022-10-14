package main

import (
	"context"
	"log"
	"time"

	pb "github.com/javiergomezve-grpc-course/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoke")

	reqS := []*pb.GreetRequest{
		{FirstName: "Javier"},
		{FirstName: "Eli"},
		{FirstName: "Ramon"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v\n", err)
	}

	for _, req := range reqS {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(time.Second * 1)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
