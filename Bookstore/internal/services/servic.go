package services

import (
	"github.com/Yhlas9/gocrud.git/Bookstore/internal/models"
	"github.com/Yhlas9/gocrud.git/Bookstore/internal/repositories"
)

//Obshyy dusunmedim name zatt. Name ucin error beryar.   Name edip beryar?

type BookService struct {
	Repository *repositories.BookRepository
}

func NewBookService(repository *repositories.BookRepository) *BookService {
	return &BookService{Repository: repository}
}

func (s *BookService) GetBooks() ([]models.Book, error) {
	return s.Repository.GetAllBooks()
}

func (s *BookService) GetBookByID(id uint) (models.Book, error) {
	return s.Repository.GetBookByID(id)
}

func (s *BookService) CreateBook(book *models.Book) error {
	return s.Repository.CreateBook(book)
}

func (s *BookService) UpdateBook(id uint, book *models.Book) error {
	return s.Repository.UpdateBook(id, book)
}

func (s *BookService) DeleteBook(id uint) error {
	return s.Repository.DeleteBook(id)
}
