package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/dkijkuit/grpc-demo/demoproto"
	"google.golang.org/grpc"
)

type server struct {
}

func (*server) Demo(ctx context.Context, request *demoproto.DemoRequest) (*demoproto.DemoResponse, error) {
	name := request.Name
	response := &demoproto.DemoResponse{
		Greeting: "Hello " + name,
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
	demoproto.RegisterDemoServiceServer(s, &server{})

	s.Serve(lis)
}
