package main

import (
	"context"
	"fmt"
	"github.com/grpc-go-course/calculator/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"time"
)

func main() {
	floatSlice := []float32{1, 5, 3, 6, 2, 20}

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting ", err)
	}
	defer conn.Close()

	client := protobuf.NewCalculatorServiceClient(conn)

	doUnary(client)

	doStreamingServer(42104500, client)

	doClientStreaming(floatSlice, client)

	doBiDiStreaming(floatSlice, client)

	doErrUnary(client)
}

func doErrUnary(c protobuf.CalculatorServiceClient) {
	fmt.Println("Doing err unary ...")
	request := &protobuf.SquareRootRequest{
		Number: -10,
	}
	resp, err := c.SquareRoot(context.Background(), request)
	if err != nil {
		stat, ok := status.FromError(err)
		if ok {
			fmt.Printf("status code [%v] | status msg [%v]", stat.Code(), stat.Message())
		} else {
			log.Fatal(err)
		}
		return
	}
	log.Printf("Response => %v", resp.SquareRoot)
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

func doBiDiStreaming(floatSlice []float32, c protobuf.CalculatorServiceClient) {

	stream, err := c.ComputeMaximum(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	waitChann := make(chan struct{})
	go func() {
		for _, v := range floatSlice {
			fmt.Println("sending number ", v)
			req := &protobuf.ComputeMaximumRequest{
				Number: v,
			}
			err := stream.Send(req)
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
		err := stream.CloseSend()
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}
			result := resp.MaxNumber
			log.Printf("Current maximum => [%v]", result)
		}
		close(waitChann)
	}()

	<- waitChann

}