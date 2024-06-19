package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	r := gin.Default()

	r.GET("/", rootHandler)
	r.GET("/books/:id", booksHandler)
	r.GET("/query", queryHandler)
	r.POST("/books", postBookHandler)

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
	id := c.Query("id")

	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"id":    id,
	})
}

type BookInput struct {
	Title    string      `json:"title" binding:"required"`
	SubTitle string      `json:"sub_title" binding:"required"`
	Price    json.Number `json:"price" binding:"required,number"`
}

func postBookHandler(c *gin.Context) {
	var bookInput BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		// jangan menggunakan log.Fatal karna akan berhenti
		// log.Fatal(err)
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorrMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorrMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title":     bookInput.Title,
		"sub_title": bookInput.SubTitle,
		"price":     bookInput.Price,
	})
}
