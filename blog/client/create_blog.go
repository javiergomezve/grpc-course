package main

import (
	"context"
	"log"

	pb "github.com/javiergomezve-grpc-course/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("createBlog was invoke")

	blog := &pb.Blog{
		AuthorId: "Javier",
		Title:    "My first blog",
		Content:  "This is the content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("blog has been created: %s\n", res.Id)
	return res.Id
}
