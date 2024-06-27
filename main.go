package main

import (
	"belajar-golang-gin/handler"
	"belajar-golang-gin/repository"
	"belajar-golang-gin/service"
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func main() {
	// Koneksi ke db
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:12345@localhost:5432/belajar-golang-gin")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("connect to database")

	defer conn.Close(context.Background())

	// Handle request dan response book
	bookRepository := repository.NewBookRepository(conn)
	bookService := service.NewBookService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	r := gin.Default()

	v1 := r.Group("/v1")
	v1.GET("/books", bookHandler.FindAllBookHandler)
	v1.POST("/book", bookHandler.CreateNewBookHandler)
	v1.GET("/book/:book_id", bookHandler.FindByIdHandler)
	v1.DELETE("/book/:book_id", bookHandler.DeleteBookHandler)
	// v1.PUT("/:book_id", handler.UpdateBookHandler)

	r.Run(":8080")

	// main
	// handler
	// service
	// repository
	// db
}
