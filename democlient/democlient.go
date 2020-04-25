package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dkijkuit/grpc-demo/demoproto"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello client ...")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := demoproto.NewDemoServiceClient(cc)
	request := &demoproto.DemoRequest{Name: "David"}

	resp, _ := client.Demo(context.Background(), request)
	fmt.Printf("Receive response => [%v]", resp.Greeting)
}
