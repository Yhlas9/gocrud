package repositories

import (
	"gorm.io/gorm"  
	"gocrud/models"
)

type BookRepository interface {         
	GetAllBooks() ([]models.Book, error)  
	GetBookByID(id int) (*models.Book, error)  
	CreateBook(book *models.Book) error    
	UpdateBook(id int, book *models.Book) error 
	DeleteBook(id int) error 
}

type bookRepository struct {   
	db *gorm.DB 
} 

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}   
	
func (r *bookRepository) GetAllBooks() ([]models.Book, error) {
	var books []models.Book  
	
	err := r.db.Find(&books).Error  
	return books, err    
}

func (r *bookRepository) GetBookByID(id int) (*models.Book, error) {   
	var book models.Book  
	 
	err := r.db.First(&book, id).Error  
		return &book, err    
}                            

func (r *bookRepository) CreateBook(book *models.Book) error {
  	return r.db.Create(book).Error   
}                                    
                                    

func (r *bookRepository) UpdateBook(id int, book *models.Book) error {   
	return r.db.Model(&models.Book{}).Where("id = ?", id).Updates(book).Error 
}

func (r *bookRepository) DeleteBook(id int) error {  
	return r.db.Delete(&models.Book{}, id).Error   
}
