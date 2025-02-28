package services

import (
	"gocrud/models"
	"gocrud/repositories"
)

//Obshyy dusunmedim name zatt. Name ucin error beryar.   Name edip beryar?

type BookService struct {
	Repository repositories.BookRepository
}

func NewBookService(repository repositories.BookRepository) *BookService {
	return &BookService{Repository: repository}
}

func (s *BookService) GetBooks() ([]models.Book, error) {
	return s.Repository.GetAllBooks()
}

func (s *BookService) GetBookByID(id int) (*models.Book, error) {
	return s.Repository.GetBookByID(id)
}

func (s *BookService) CreateBook(book *models.Book) error {
	return s.Repository.CreateBook(book)
}

func (s *BookService) UpdateBook(id int, book *models.Book) error {
	return s.Repository.UpdateBook(id, book)
}

func (s *BookService) DeleteBook(id int) error {
	return s.Repository.DeleteBook(id)
}