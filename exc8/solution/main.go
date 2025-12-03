package main

import (
	"exc8/client"
	"exc8/server"
	"fmt"
	"time"
)

func main() {
	go func() {
		// todo start server
		if err := server.StartGrpcServer(); err != nil {
			fmt.Printf("gRPC server error: %v", err)
		}
	}()

	time.Sleep(1 * time.Second)

	// todo start client
	cli, err := client.NewGrpcClient()
	if err != nil {
		fmt.Printf("could not create client: %v", err)
	}

	if err := cli.Run(); err != nil {
		fmt.Printf("client error: %v", err)
	}
	println("Orders complete!")
}
