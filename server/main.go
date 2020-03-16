package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	pb "proto-sandbox/my_service"
)

type server struct {
}

func (s *server) Add(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	a, b := req.GetA(), req.GetB()

	result := a + b

	return &pb.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	a, b := req.GetA(), req.GetB()
	result := a * b

	return &pb.Response{Result: result}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	pb.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	log.Println("Started server on :4040")
	if err = srv.Serve(listener); err != nil {
		panic(err)
	}
}
