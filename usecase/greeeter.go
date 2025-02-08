package usecase

import (
	"context"
	pb "grpc-repos/protos"
)

// GreeterUsecase defines the interface for the greeter use case
type GreeterUsecase interface {
	SayHello(ctx context.Context, name string) (*pb.HelloResponse, error)
}

// greeterUsecase is the implementation of GreeterUsecase
type greeterUsecase struct{}

// NewGreeterUsecase creates a new instance of greeterUsecase
func NewGreeterUsecase() GreeterUsecase {
	return &greeterUsecase{}
}

// SayHello implements the business logic for the greeter use case
func (u *greeterUsecase) SayHello(ctx context.Context, name string) (*pb.HelloResponse, error) {
	message := "Hello " + name
	return &pb.HelloResponse{Message: message}, nil
}
