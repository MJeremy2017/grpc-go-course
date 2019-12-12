package main

import (
	"context"
	"fmt"
	"github.com/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

type server struct {

}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	result := "Hello " + req.Greeting.FirstName + req.Greeting.LastName

	response := &greetpb.GreetResponse{
		Result: result,
	}
	return response, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	firstName := req.Greeting.FirstName
	for i := 0; i <= 10; i++ {
		resp := "Hello " + firstName + " number " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: resp,
		}
		stream.Send(res)

		//time.Sleep(time.Second)
	}

	return nil

}

func main() {

	fmt.Println("Inside server ...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")  // open a port
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
