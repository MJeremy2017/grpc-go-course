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

	fmt.Println("Creating blogs")
	resp, err := client.CreateBlog(context.Background(), request)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("response => ", resp)

	fmt.Println("Reading blogs")
	readRequest := &blogpb.ReadBlogRequest{
		BlogId: resp.Blog.Id,
	}

	readResp, _ := client.ReadBlog(context.Background(), readRequest)
	fmt.Println("response => ", readResp)

	fmt.Println("Updating blogs")
	updateRequest := &blogpb.UpdateBlogRequest{
		Blog: &blogpb.Blog{
			Id: resp.Blog.Id,
			Title: "new title",
			AuthorId: "new author",
			Content: "new content",
		},
	}

	updateResp, _ := client.UpdateBlog(context.Background(), updateRequest)
	fmt.Printf("response => [%v]", updateResp)

}
