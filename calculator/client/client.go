package main

import (
	"context"
	"fmt"
	"github.com/grpc-go-course/calculator/protobuf"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting ", err)
	}
	defer conn.Close()

	client := protobuf.NewCalculatorServiceClient(conn)

	doUnary(client)
	doStreamingServer(42104500, client)
}

func doUnary(c protobuf.CalculatorServiceClient) {

	request := &protobuf.SumRequest{
		Num1: 12.3,
		Num2: 2.3,
	}

	response, _ := c.Sum(context.Background(), request)

	fmt.Printf("Successfully get response %v", response.Summation)

}

func doStreamingServer(number int64, c protobuf.CalculatorServiceClient) {

	request := &protobuf.PrimeNumberRequest{
		Number: number,
	}

	streamResp, _ := c.PrimeNumberDecomposition(context.Background(), request)

	for {
		resp, err := streamResp.Recv()
		if err == io.EOF {
			break
		}
		log.Println("Response -> ", resp.PrimeNumber)
	}

}