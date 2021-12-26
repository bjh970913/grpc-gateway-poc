package main

import (
	"context"
	pb "github.com/bjh970913/grpc-gateway-poc/gen/go"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedYourServiceServer
}

func (s server) Echo(ctx context.Context, message *pb.StringMessage) (*pb.StringMessage, error) {
	log.Println(message.Value)
	return &pb.StringMessage{
		Value: message.Value,
	}, nil
}

func main() {
	log.Println("TEST")
	lis, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterYourServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
