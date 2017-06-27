package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "grpcserver/generated/servicedef"
)

const (
	port = ":1234"
)

type server struct {
	count int32
}

func (s *server) Increment(ctx context.Context, in *pb.IncrementRequest) (*pb.IncrementResponse, error) {
	log.Print("Received increment request from " + in.Requester)
	s.count++
	return &pb.IncrementResponse{Count: s.count}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCounterServer(s, &server{})

	log.Print("Listening...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
