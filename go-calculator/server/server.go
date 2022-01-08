package main

import (
	"context"
	"github.com/sainad2222/go-calculator/api"
	"google.golang.org/grpc"
	"log"
	"math"
	"net"
)

type server struct {
	api.UnimplementedCalculatorServiceServer
}

func (s server) Sum(_ context.Context, req *api.TwoIntInputRequest) (*api.OneIntOutputResponse, error) {
	first := req.GetFirst()
	second := req.GetSecond()
	return &api.OneIntOutputResponse{Result: first + second}, nil
}

func (s server) Difference(_ context.Context, req *api.TwoIntInputRequest) (*api.OneIntOutputResponse, error) {
	first := req.GetFirst()
	second := req.GetSecond()
	return &api.OneIntOutputResponse{Result: int32(int(math.Abs(float64(first - second))))}, nil
}

func (s server) Multiply(_ context.Context, req *api.TwoIntInputRequest) (*api.OneIntOutputResponse, error) {
	first := req.GetFirst()
	second := req.GetSecond()
	return &api.OneIntOutputResponse{Result: first * second}, nil
}

func (s server) Divide(_ context.Context, req *api.TwoIntInputRequest) (*api.OneFloatOutputResponse, error) {
	first := req.GetFirst()
	second := req.GetSecond()
	return &api.OneFloatOutputResponse{Result: float32(first) / float32(second)}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8097")
	if err != nil {
		log.Fatalln("Error listening", err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterCalculatorServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Error Serving", err)
	}
}
