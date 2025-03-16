package handlers

import (
	"errors"
	"gocrud/models"
	"gocrud/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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
	bookID,_ := strconv.Atoi(c.Param("id"))
	book, err := h.Service.GetBookByID(bookID)

	if err != nil {
		if errors.Is(err,gorm.ErrRecordNotFound){
			c.JSON(http.StatusNotFound, gin.H{"message": "Книга не найдена"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
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

func (h *BookHandler) UpdateBook(c *gin.Context) {
	bookID,_ := strconv.Atoi(c.Param("id"))

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
	bookID,_ := strconv.Atoi(c.Param("id"))

	if err := h.Service.DeleteBook(bookID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Книга не найдена"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Книга удалена"})
}
