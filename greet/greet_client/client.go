package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/neighborhoodhacker/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	reply, err := c.Hello(ctx, &greetpb.HelloName{Name: "Braydon"})
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	fmt.Println(reply.GetReply())
}
