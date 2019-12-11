package main

import (
	"fmt"
	"github.com/grpc-go-course/calculator/protobuf"
	"google.golang.org/grpc"
	"context"
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

func main() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()
	protobuf.RegisterSumServiceServer(s, &server{})

	s.Serve(lis)
}


