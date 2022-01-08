package main

import (
	"context"
	"fmt"
	"github.com/sainad2222/go-calculator/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8097", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Error Dialing", err)
	}

	c := api.NewCalculatorServiceClient(conn)
	sumRes, err := c.Sum(context.Background(), &api.TwoIntInputRequest{First: 25, Second: 5})
	diffRes, err := c.Difference(context.Background(), &api.TwoIntInputRequest{First: 25, Second: 5})
	mulRes, err := c.Multiply(context.Background(), &api.TwoIntInputRequest{First: 25, Second: 5})
	divRes, err := c.Divide(context.Background(), &api.TwoIntInputRequest{First: 25, Second: 5})
	fmt.Println("Result of sum:", sumRes)
	fmt.Println("Result of sum:", diffRes)
	fmt.Println("Result of sum:", mulRes)
	fmt.Println("Result of sum:", divRes)
}
