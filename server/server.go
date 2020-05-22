package main

import (
	"context"
	"log"
	pb "lozi-training/grpc/customer"
	"net"
	"strings"

	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

// server is used to implement customer.CustomerServer.
type server struct {
	savedCustomers []*pb.CustomerRequest
}

// truyền vào context và customer cần tạo
func (s *server) CreateCustomer(ctx context.Context, inputCustomer *pb.CustomerRequest) (*pb.CustomerResponse, error) {
	s.savedCustomers = append(s.savedCustomers, inputCustomer)
	return &pb.CustomerResponse{Id: inputCustomer.Id, Success: true}, nil
}

// GetCustomers returns all customers by given filter
func (s *server) GetCustomers(filter *pb.CustomerFilter, stream pb.Hello_GetCustomersServer) error {
	for _, customer := range s.savedCustomers {
		if filter.Keyword != "" {
			if !strings.Contains(customer.Name, filter.Keyword) {
				continue
			}
		}
		if err := stream.Send(customer); err != nil {
			return err
		}
	}
	return nil
}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &server{})
	s.Serve(lis)
}
