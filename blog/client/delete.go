package main

import (
	"context"
	"log"

	pb "github.com/javiergomezve-grpc-course/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("deleteBlog was invoke")

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})

	if err != nil {
		log.Fatalf("Error while deleting: %v\n", err)
	}

	log.Println("Blog was delete")
}
