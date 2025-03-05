package repository

import "grpc-repos/domain/entity"

type BookRepository interface {
	Create(book *entity.Book) error
	FindByISBN(isbn string) (*entity.Book, error)
}
