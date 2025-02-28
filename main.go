package main

import (
	"log"
	"github.com/Yhlas9/gocrud.git/Bookstore/internal/models"
	"github.com/Yhlas9/gocrud.git/Bookstore/internal/handlers"
	"github.com/Yhlas9/gocrud.git/Bookstore/internal/repositories"
	"github.com/Yhlas9/gocrud.git/Bookstore/internal/roters"
	"github.com/Yhlas9/gocrud.git/Bookstore/internal/services"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Bara sqlite bazadanny acya . Catylya
	db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	// ?
	db.AutoMigrate(&models.Book{})

	// ?
	bookRepo := repositories.NewBookRepository(db)

	// ?
	bookService := services.NewBookService(bookRepo)

	// ? 
	bookHandler := handlers.NewBookHandler(bookService)

	// ?
	r := roters.SetupRouter(bookHandler)

	// Serwery isletmek
	r.Run(":8080")
}


// TODO (You can't eat or sleep before finishing this task and you have 24 hours to finish it,  not wast your time and start now !!)

// Task 1: Create Book Store API

// You should have 5 routes:
// 1. /books/ - GET - Return all books
// 2. /books/ - POST - Create a new book
// 3. /books/{id} - GET - Return a book by ID
// 4. /books/{id} - PUT - Update a book by ID
// 5. /books/{id} - DELETE - Delete a book by ID

// Tools and Requirements:
// - Use the Gin framework to create your API
// - Use GORM to interact with the database
// - Use gorm.io/driver/sqlite to connect to a SQLite3 database
// - Use github.com/gin-gonic/gin to create your API
// - Use gorm.io/gorm to interact with the database
// - Implement a clean architecture with separate directories for:
//     - models
//     - repositories
//     - services
//     - handlers
//     - routers

// Book struct should have the following fields:
// - id (auto-incremented, primary key)
// - title (string)
// - author (string)
// - year (int)
// - pages (int)
// - language (string)
// - publisher (string)
// - rating (float64)

// Steps:
// 1. Set up Gin router and define the routes
// 2. Create the Book model
// 3. Implement CRUD functionality in the repository layer
// 4. Implement business logic in the service layer
// 5. Write handlers to handle incoming requests and connect to services
// 6. Connect your application to the SQLite database using GORM
// 7. Test the API using Postman or another HTTP client

// Remember:
// - Don't use AI to write code directly. 
// - Focus on asking about patterns, principles, or how to approach each part of the task.
// - You can always ask questions about design decisions, architecture, or best practices. 
// - I will guide you on how to implement it step by step, not by writing code but by providing clear instructions and answers.