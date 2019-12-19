package main

import (
	"context"
	"fmt"
	"github.com/grpc-go-course/blog/blogpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Blog client ...")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := blogpb.NewBlogServiceClient(cc)

	request := &blogpb.CreateBlogRequest{
		Blog: &blogpb.Blog{
			AuthorId: "SS",
			Title: "first blog",
			Content: "Nothing fun",
		},
	}

	resp, err := client.CreateBlog(context.Background(), request)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("response => ", resp)
}
