package handlers

import (
	"net/http"
	"github.com/Yhlas9/gocrud.git/Bookstore/internal/models"
	"github.com/Yhlas9/gocrud.git/Bookstore/internal/services"
	"github.com/gin-gonic/gin"
)

//Korocem repositories baza dannydan alyp beryan yeri. Bu handleram duzetyan yerimi kodung icinde. Baza danna girman
// Eyle dal bolsa dusundirip berai azajyk
type BookHandler struct {
	Service *services.BookService
}

func NewBookHandler(service *services.BookService) *BookHandler {
	return &BookHandler{Service: service}
}

func (h *BookHandler) GetBooks(c *gin.Context) {
	books, err := h.Service.GetBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Не удалось получить книги"})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (h *BookHandler) GetBookByID(c *gin.Context) {
	id := c.Param("id")
	bookID := uint(id)
	book, err := h.Service.GetBookByID(bookID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Книга не найдена"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Неверный формат данных"})
		return
	}

	if err := h.Service.CreateBook(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Не удалось создать книгу"})
		return
	}
	c.JSON(http.StatusCreated, book)
}

// UpdateBook обрабатывает PUT /books/:id и обновляет книгу по ID
func (h *BookHandler) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	bookID := uint(id)

	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Неверный формат данных"})
		return
	}

	if err := h.Service.UpdateBook(bookID, &book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Не удалось обновить книгу"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	bookID := uint(id)

	if err := h.Service.DeleteBook(bookID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Книга не найдена"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Книга удалена"})
}
