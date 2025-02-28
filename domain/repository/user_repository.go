package repository

import "grpc-repos/domain/entity"

type UserRepository interface {
	Create(user *entity.User) error
	FindByID(id string) (*entity.User, error)
	Update(user *entity.User) error
	Delete(id uint) error
	// 必要に応じてその他メソッド(FindByEmail, ListAllなど)
}
