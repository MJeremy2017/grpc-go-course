package main

import (
	"fmt"
	"github.com/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
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

	fmt.Println(request)

}
