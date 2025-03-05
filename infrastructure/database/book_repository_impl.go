package database

import (
	"grpc-repos/domain/entity"
	"grpc-repos/domain/repository"

	"gorm.io/gorm"
)

type bookRepositoryImpl struct {
	db *gorm.DB
}

func NewBookRepositoryImpl(db *gorm.DB) repository.BookRepository {
	return &bookRepositoryImpl{db: db}
}

func (r *bookRepositoryImpl) Create(book *entity.Book) error {
	return r.db.Create(book).Error
}

func (r *bookRepositoryImpl) FindByISBN(isbn string) (*entity.Book, error) {
	var b entity.Book
	// Authorを同時に読み込み
	if err := r.db.Preload("Author").First(&b, "isbn = ?", isbn).Error; err != nil {
		return nil, err
	}
	return &b, nil
}
