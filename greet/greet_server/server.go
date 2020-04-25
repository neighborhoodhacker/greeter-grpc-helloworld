package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/neighborhoodhacker/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreetServiceServer
}

func (s *server) Hello(ctx context.Context, in *pb.HelloName) (*pb.HelloReply, error) {
	return &pb.HelloReply{Reply: ("Hello " + in.GetName())}, nil
}

func main() {
	fmt.Println("Hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
}
