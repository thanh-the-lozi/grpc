package main

import (
	"context"
	pd "lozi-training/grpc/customer"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pd.UnimplementedGreetServer
}

func (s *server) GetStudent(ctx context.Context, r *pd.SayHello) (*pd.HelloReply, error) {
	data := pd.HelloReply{
		Message: r.Name,
	}
	return &data, nil
}

func (s *server) GetStudentAgain(ctx context.Context, r *pd.SayHello) (*pd.HelloReply, error) {
	data := pd.HelloReply{
		Message: "hello again " + r.Name,
	}
	return &data, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":8080")
	s := grpc.NewServer()
	pd.RegisterGreetServer(s, &server{})
	s.Serve(lis)
}
