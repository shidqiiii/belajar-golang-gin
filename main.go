package main

import (
	"github.com/gin-gonic/gin"

	"belajar-golang-gin/handler"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	v1.GET("/", handler.RootHandler)
	v1.GET("/:id", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.GET("/books", handler.GetBookHandler)
	v1.POST("/book", handler.PostBookHandler)
	v1.GET("/book/:book_id", handler.GetDetailBookHandler)
	v1.DELETE("/:book_id", handler.DeleteBookHandler)
	v1.PUT("/:book_id", handler.UpdateBookHandler)

	r.Run(":8080")
}
