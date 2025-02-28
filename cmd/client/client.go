package main

import (
	"context"
	"log"
	"time"

	pb "grpc-repos/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address     = "localhost:8080"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.CreateUser(ctx, &pb.CreateUserRequest{Name: "name", Email: "mail"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
}
