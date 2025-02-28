package controller

import (
	"context"
	"grpc-repos/infrastructure/database"
	pb "grpc-repos/protos"
	"grpc-repos/usecase"
	"log"
	"net"

	"github.com/k0kubun/pp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func ServerStart(lis net.Listener) {
	pp.Println("ServerStart")
	d, err := database.NewMySQLDB()
	if err != nil {
		log.Fatalf("failed to start db: %v", err)
	}

	userRepository := database.NewUserRepositoryImpl(d.DB)
	userUsecase := usecase.NewUserUsecase(userRepository)
	user := NewUserController(userUsecase)

	u := grpc.NewServer()
	pb.RegisterUserServer(u, user)
	reflection.Register(u)
	log.Printf("Server listening at %v", lis.Addr())
	if err := u.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

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
