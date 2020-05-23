package main

import (
	"context"
	"net"
	pd "source/grpc/proto"

	"google.golang.org/grpc"
)

type server struct {
	pd.UnimplementedGreetServer
}

func (s *server) Hello(ctx context.Context, r *pd.SayHello) (*pd.HelloReply, error) {
	data := pd.HelloReply{
		Message: "Hello: " + r.Name,
	}
	return &data, nil
}

func (s *server) HelloAgain(ctx context.Context, r *pd.SayHello) (*pd.HelloReply, error) {
	data := pd.HelloReply{
		Message: "Hello again: " + r.Name,
	}
	return &data, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":8080")
	s := grpc.NewServer()
	pd.RegisterGreetServer(s, &server{})
	s.Serve(lis)
}
