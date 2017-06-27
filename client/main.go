package main

import (
	"log"
	"os"

	pb "grpcserver/generated/servicedef"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:1234"
	defaultName = "Go Tester"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCounterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.Increment(context.Background(), &pb.IncrementRequest{Requester: name})
	if err != nil {
		log.Fatalf("could not increment: %v", err)
	}
	log.Printf("It's the %d request the server processed", r.Count)
}
