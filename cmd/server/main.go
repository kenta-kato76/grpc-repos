package main

import (
	"grpc-repos/controller"
	pb "grpc-repos/protos"
	"grpc-repos/usecase"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8080"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	greeterUsecase := usecase.NewGreeterUsecase()
	server := controller.NewServer(greeterUsecase)
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, server)
	reflection.Register(s)
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
