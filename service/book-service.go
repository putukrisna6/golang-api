package service

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"github.com/putukrisna6/golang-api/dto"
	"github.com/putukrisna6/golang-api/entity"
	"github.com/putukrisna6/golang-api/repository"
)

type BookService interface {
	Insert(b dto.BookCreateDTO) entity.Book
	Update(b dto.BookUpdateDTO) entity.Book
	Get(bookID uint64) entity.Book
	Delete(b entity.Book)
	All() []entity.Book
	IsAllowedToEdit(userID string, bookID uint64) bool
}

type bookService struct {
	bookRepository repository.BookRepository
}

func NewBookService(bookRepository repository.BookRepository) BookService {
	return &bookService{
		bookRepository: bookRepository,
	}
}

func (service *bookService) Insert(b dto.BookCreateDTO) entity.Book {
	newBook := entity.Book{}
	err := smapping.FillStruct(&newBook, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("failed to map: %v", err.Error())
	}

	res := service.bookRepository.InsertBook(newBook)
	return res
}

func (service *bookService) Update(b dto.BookUpdateDTO) entity.Book {
	book := entity.Book{}
	err := smapping.FillStruct(&book, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("failed to map: %v", err.Error())
	}

	res := service.bookRepository.UpdateBook(book)
	return res
}

func (service *bookService) Get(bookID uint64) entity.Book {
	return service.bookRepository.GetBook(bookID)
}

func (service *bookService) Delete(b entity.Book) {
	service.bookRepository.DeleteBook(b)
}

func (service *bookService) All() []entity.Book {
	return service.bookRepository.AllBooks()
}

func (service *bookService) IsAllowedToEdit(userID string, bookID uint64) bool {
	book := service.bookRepository.GetBook(bookID)
	return userID == fmt.Sprintf("%v", book.UserID)
}
