package service

import (
	"github.com/putukrisna6/golang-api/dto"
	"github.com/putukrisna6/golang-api/entity"
	"github.com/putukrisna6/golang-api/repository"
)

type BookService interface {
	Insert(b dto.BookCreateDTO) entity.Book
	Update(b dto.BookUpdateDTO) entity.Book
	Get(bookID string) entity.Book
	Delete(b entity.Book)
	All() []entity.Book
	IsAllowedToEdit(userID string, bookID uint64) bool
}

type bookService struct {
	bookReposity repository.BookRepository
}

func NewBookService(bookRepository repository.BookRepository) BookService {
	return &bookService{
		bookReposity: bookRepository,
	}
}

func (service *bookService) Insert(b dto.BookCreateDTO) entity.Book {

}

func (service *bookService) Update(b dto.BookUpdateDTO) entity.Book {

}

func (service *bookService) Get(bookID string) entity.Book {

}

func (service *bookService) Delete(b entity.Book) {

}

func (service *bookService) All() []entity.Book {

}

func (service *bookService) IsAllowedToEdit(userID string, bookID uint64) bool {

}
