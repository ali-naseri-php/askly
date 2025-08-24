package main

import (
	"context"
	"log"
	"net"
     authpb "github.com/ali-naseri-php/Askly/proto/auth"


	"google.golang.org/grpc"
)

type authServer struct {
	authpb.UnimplementedAuthServiceServer
}

func (s *authServer) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	// فقط پاسخ تستی
	return &authpb.LoginResponse{Token: "dummy-token"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &authServer{})

	log.Println("Auth service running on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
