package main

import (
	"fmt"
	"github.com/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"context"
)

func main() {

	fmt.Println("Inside client ...")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	client := greetpb.NewGreetServiceClient(conn)

	doUnary(client)

	doServerStreaming(client)

}

func doUnary(c greetpb.GreetServiceClient)  {
	fmt.Println("Doing unary call ...")

	request := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "jeremy",
			LastName:  "zhang",
		},
	}

	response, _ := c.Greet(context.Background(), request)
	log.Printf("Response -> %v", response.Result)

}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Doing server call ...")

	request := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "jeremy",
			LastName:  "zhang",
		},
	}

	streamResp, err := c.GreetManyTimes(context.Background(), request)
	if err != nil {
		log.Fatalf("Err [%v]", err)
	}

	for {
		msg, err := streamResp.Recv()
		if err == io.EOF {
			// end of file
			break
		}
		log.Println("Streaming response -> ", msg.Result)

	}

}
