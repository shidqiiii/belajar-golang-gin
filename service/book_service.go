package service

import (
	"belajar-golang-gin/model"
	"belajar-golang-gin/repository"
)

type BookService interface {
	FindAll() ([]model.Book, error)
	FindById(id int) (model.Book, error)
	CreateNewBook(book *model.Book) error
	DeleteBook(id int) error
}

type bookService struct {
	repository repository.BookRepository
}

func NewBookService(repository repository.BookRepository) *bookService {
	return &bookService{repository}
}

func (s *bookService) FindAll() ([]model.Book, error) {
	books, err := s.repository.FindAll()
	return books, err
	// return s.repository.FindAll()
}

func (s *bookService) FindById(id int) (model.Book, error) {
	book, err := s.repository.FindById(id)
	return book, err
}

func (s *bookService) CreateNewBook(book *model.Book) error {
	err := s.repository.CreateNewBook(book)
	return err
}

func (s *bookService) DeleteBook(id int) error {
	err := s.repository.DeleteBook(id)
	return err
}
