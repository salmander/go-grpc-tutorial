//go:generate protoc -I ../userservice --go_out=plugins=grpc:../userservice ../userservice/userservice.proto

package main

import (
	"context"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/salmander/go-grpc-tutorial/userservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const port = ":50052"

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) Healthcheck(ctx context.Context, req *empty.Empty) (*pb.Health, error) {
	return &pb.Health{
		Message:              "all good",
		Errors:               "none",

	}, nil
}

func (s *server) GetUserById(ctx context.Context, req *pb.UserByIdRequest) (*pb.User, error) {
	switch req.GetId() {
	case 9:
		return &pb.User{
			Id:                   9,
			Uuid:                 "batman-911",
			FirstName:            "Bruce",
			LastName:             "Wayne",
			Email:                "imbatman@justiceleague.com",
			NectarCard:           "911",
		}, nil
	default:
		return &pb.User{}, status.Errorf(codes.NotFound, "unknown user id")
	}
}

//func (s *server) GetUserByUuid(ctx context.Context, req *pb.UserByUuidRequest) (*pb.User, error) {
//	return nil, nil
//}
//func (s *server) GetUserByEmail(ctx context.Context, req *pb.UserByEmailRequest) (*pb.User, error) {
//	return nil, nil
//}
//func (s *server) GetUserByNectar(ctx context.Context, req *pb.UserByNectarRequest) (*pb.User, error) {
//	return nil, nil
//}

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

