package usecase

import (
	"grpc-repos/domain/entity"
	"grpc-repos/domain/repository"
)

type UserUsecase interface {
	CreateUser(name, email string) (*entity.User, error)
	GetUserByID(name string) (*entity.User, error)
	UpdateUser(id, name, email string) (*entity.User, error)
	DeleteUser(id uint) error
}

type userUsecaseImpl struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecaseImpl{userRepo: repo}
}

func (u *userUsecaseImpl) CreateUser(name, email string) (*entity.User, error) {
	// バリデーションなど
	user := &entity.User{
		Name:  name,
		Email: email,
	}
	err := u.userRepo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecaseImpl) GetUserByID(name string) (*entity.User, error) {
	return u.userRepo.FindByID(name)
}

func (u *userUsecaseImpl) UpdateUser(id, name, email string) (*entity.User, error) {
	user, err := u.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	user.Name = name
	user.Email = email
	err = u.userRepo.Update(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecaseImpl) DeleteUser(id uint) error {
	return u.userRepo.Delete(id)
}
