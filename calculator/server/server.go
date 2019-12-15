package main

import (
	"fmt"
	"github.com/grpc-go-course/calculator/protobuf"
	"google.golang.org/grpc"
	"context"
	"io"
	"log"
	"net"
)

type server struct {

}

func (*server) Sum(ctx context.Context, req *protobuf.SumRequest) (*protobuf.SumResponse, error) {

	fmt.Println("received info -> ", req)

	response := &protobuf.SumResponse{
		Summation: req.Num1 + req.Num2,
	}
	return response, nil
}

func (*server) PrimeNumberDecomposition(req *protobuf.PrimeNumberRequest, stream protobuf.CalculatorService_PrimeNumberDecompositionServer) error {
	inputNumber := req.Number
	fmt.Println("received number", inputNumber)
	n := int64(2)

	for inputNumber > 1 {
		if inputNumber % n == 0 {
			resp := &protobuf.PrimeNumberResponse{
				PrimeNumber: n,
			}
			stream.Send(resp)

			inputNumber = inputNumber / n
		} else {
			n += 1
		}
	}

	return nil
}

func (*server) ComputeAverage(stream protobuf.CalculatorService_ComputeAverageServer) error {
	sum := float32(0)
	count := float32(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&protobuf.ComputeAverageResponse{
				Result: sum/count,
			})
		}
		if err != nil {
			log.Fatalf("Err [%v]", err)
		}

		log.Println("received float ", req.Number)
		sum += req.Number
		count += 1
	}

}

func main() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()
	protobuf.RegisterCalculatorServiceServer(s, &server{})

	s.Serve(lis)
}


