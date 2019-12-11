package main

import (
	"context"
	"fmt"
	"github.com/grpc-go-course/calculator/protobuf"
	"google.golang.org/grpc"
	"log"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting ", err)
	}
	defer conn.Close()

	client := protobuf.NewSumServiceClient(conn)

	request := &protobuf.SumRequest{
		Num1: 12.3,
		Num2: 2.3,
	}

	response, _ := client.Sum(context.Background(), request)

	fmt.Printf("Successfully get response %v", response.Summation)

}
