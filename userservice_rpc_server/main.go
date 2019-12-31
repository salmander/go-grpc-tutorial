//go:generate protoc -I ../userservice --go_out=plugins=grpc:../userservice ../userservice/userservice.proto

package main

import (
	"context"
	"log"
	"net"

	pb "github.com/salmander/go-grpc-tutorial/userservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const port = ":50052"

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) GetCustomerById(ctx context.Context, req *pb.CustomerByIdRequest) (*pb.Customer, error) {
	switch req.GetId() {
	case 9:
		return &pb.Customer{
			Id:                   9,
			Uuid:                 "batman-911",
			FirstName:            "Bruce",
			LastName:             "Wayne",
			Email:                "imbatman@justiceleague.com",
			NectarCard:           "911",
		}, nil
	default:
		return &pb.Customer{}, status.Errorf(codes.NotFound, "unknown customer id")
	}
}

func (s *server) GetCustomerByUuid(ctx context.Context, req *pb.CustomerByUuidRequest) (*pb.Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerByUuid not implemented")
}
func (s *server) GetCustomerByEmail(ctx context.Context, req *pb.CustomerByEmailRequest) (*pb.Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerByEmail not implemented")
}
func (s *server) GetCustomerByNectar(ctx context.Context, req *pb.CustomerByNectarRequest) (*pb.Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerByNectar not implemented")
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen to %s, %v", port, err)
	}
	log.Printf("listening on %s\n", port)

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

