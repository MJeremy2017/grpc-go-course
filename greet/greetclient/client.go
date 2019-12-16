package main

import (
	"fmt"
	"github.com/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"context"
	"time"
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

	doClientStreaming(client)

	DoBiDiStreaming(client)

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

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Doing client side streaming")
	requests := []*greetpb.LongGreetRequest{
		{
			Greeting: &greetpb.Greeting{FirstName: "A"},
		},
		{
			Greeting: &greetpb.Greeting{FirstName: "B"},
		},
		{
			Greeting: &greetpb.Greeting{FirstName: "C"},
		},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Err [%v]", err)
	}

	for _, req := range requests {
		stream.Send(req)
		time.Sleep(time.Second)
	}

	resp, _ := stream.CloseAndRecv()
	log.Printf("Response -> %v", resp)


}

func DoBiDiStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Doing bidirectional streaming ...")

	requests := []*greetpb.GreetEveryoneRequest{
		{
			Greeting: &greetpb.Greeting{FirstName: "A"},
		},
		{
			Greeting: &greetpb.Greeting{FirstName: "B"},
		},
		{
			Greeting: &greetpb.Greeting{FirstName: "C"},
		},
	}

	stream, err := c.GreetEveryOne(context.Background())
	if err != nil {
		log.Fatalf("Err [%v]", err)
	}

	waitc := make(chan struct{})
	// send messages to the client async
	go func() {
		for _, req := range requests {
			fmt.Println("Sending msg -> ", req.Greeting)
			stream.Send(req)
			time.Sleep(time.Second)
		}
		stream.CloseSend()
	}()

	// receive messages
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Err [%v]", err)
				break
			}
			log.Printf("Response -> %v", resp.Result)
		}
		close(waitc)
	}()

	<- waitc
}
