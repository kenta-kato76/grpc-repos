package controller

import (
	"context"
	pb "grpc-repos/protos"
	"grpc-repos/usecase"
	"log"
)

// server is used to implement hello.GreeterServer.
type Server struct {
	pb.UnimplementedGreeterServer
	greeterUsecase usecase.GreeterUsecase
}

func NewServer(greeterUsecase usecase.GreeterUsecase) *Server {
	return &Server{
		greeterUsecase: greeterUsecase,
	}
}

// SayHello implements hello.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return s.greeterUsecase.SayHello(ctx, in.GetName())
}
