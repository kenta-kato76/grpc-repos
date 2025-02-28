package controller

import (
	"context"
	"fmt"
	pb "grpc-repos/protos"
	"grpc-repos/usecase"
)

type UserController struct {
	pb.UnimplementedUserServer
	userUsecase usecase.UserUsecase
}

func NewUserController(u usecase.UserUsecase) *UserController {
	return &UserController{
		userUsecase: u,
	}
}

// POST /users
func (c *UserController) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user, err := c.userUsecase.CreateUser(in.Name, in.Email)
	if err != nil {
		return nil, err
	}

	fmt.Println("User: ", user)

	return &pb.CreateUserResponse{}, nil
}

// GET /users/:id
func (c *UserController) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := c.userUsecase.GetUserByID(in.Name)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserResponse{
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

// その他: UpdateUser, DeleteUser etc.
