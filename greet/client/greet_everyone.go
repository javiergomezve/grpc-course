package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/javiergomezve-grpc-course/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("goGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqS := []*pb.GreetRequest{
		{FirstName: "Eli"},
		{FirstName: "Javier"},
		{FirstName: "Alexander"},
	}

	waitC := make(chan struct{})

	go func() {
		for _, req := range reqS {
			log.Printf("Send request: %v\n", req)

			stream.Send(req)
			time.Sleep(time.Second * 1)
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}

			log.Printf("Received: %v\n", res.Result)
		}

		close(waitC)
	}()

	<-waitC
}
