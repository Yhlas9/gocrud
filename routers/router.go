package routers

import (
	"gocrud/handlers"
	"gocrud/repositories"
	"gocrud/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()	

	handlers := handlers.NewBookHandler(services.NewBookService(repositories.NewBookRepository(db)))

	// Define the routes
	r.GET("/books", handlers.GetBooks)
	r.GET("/books/:id", handlers.GetBookByID)
	r.POST("/books", handlers.CreateBook)
	r.PUT("/books/:id", handlers.UpdateBook)
	r.DELETE("/books/:id", handlers.DeleteBook)

	return r
}
