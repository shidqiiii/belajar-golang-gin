package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", rootHandler)
	r.GET("/books/:id", booksHandler)
	r.GET("/query", queryHandler)

	r.Run(":8080")
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Hello World",
	})
}

func booksHandler(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func queryHandler(c *gin.Context) {
	title := c.Query("title")

	c.JSON(http.StatusOK, gin.H{
		"title": title,
	})
}
