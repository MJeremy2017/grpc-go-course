package main

import (
	"context"
	"fmt"
	"github.com/grpc-go-course/calculator/protobuf"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
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

	floatSlice := []float32{1, 2, 3.2}
	doClientStreaming(floatSlice, client)
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
		log.Printf("Response -> %v \n", resp.PrimeNumber)
	}

}

func doClientStreaming(floatSlice []float32, c protobuf.CalculatorServiceClient) {

	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Err [%v]", err)
	}

	for _, f := range floatSlice {
		err := stream.Send(&protobuf.ComputeAverageRequest{
			Number: f,
		})
		time.Sleep(time.Second)

		if err != nil {
			log.Fatalf("Err [%v]", err)
		}

	}

	resp, _ := stream.CloseAndRecv()

	log.Printf("Response avg -> [%v]", resp.Result)

}