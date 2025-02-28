package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/Yhlas9/gocrud.git/Bookstore/internal/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Define the routes
	r.GET("/books", handlers.GetBooks)
	r.GET("/books/:id", handlers.GetBookByID)
	r.POST("/books", handlers.CreateBook)
	r.PUT("/books/:id", handlers.UpdateBook)
	r.DELETE("/books/:id", handlers.DeleteBook)

	return r
}
