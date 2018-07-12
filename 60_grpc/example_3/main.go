package main

import (
	"fmt"
	"go-learn/60_grpc/example_3/grpc"
	controllers "go-learn/60_grpc/example_3/proto"
)

func main() {
	// sample gRPC API
	p2 := grpc.NewGRPC("localhost:8081")

	// a sample implementation of a gRPC controller
	sample := controllers.NewSampleRPC(p2.Server)

	fmt.Println("START")

	go p2.Start()

	response, err := sample.Ping(nil, nil)
	fmt.Println("MSG:", response, "ERROR:", err)
}
