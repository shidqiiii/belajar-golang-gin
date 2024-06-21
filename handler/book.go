package handler

import (
	"belajar-golang-gin/book"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Hello World",
	})
}

func BooksHandler(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func QueryHandler(c *gin.Context) {
	title := c.Query("title")
	id := c.Query("id")

	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"id":    id,
	})
}

var books = []book.BookInput{
	{Id: 1, Title: "ini buku", SubTitle: "ini sub title buku", Price: 2000},
	{Id: 2, Title: "ini buku 2", SubTitle: "ini sub title buku 2", Price: 2500},
}

func GetBookHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func PostBookHandler(c *gin.Context) {
	var bookInput book.BookInput

	// var validate *validator.Validate
	validate := validator.New()

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	bookInput.Id = len(books) + 1
	if err := validate.Struct(bookInput); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorsMap := make(map[string]string)
		for _, validationErr := range validationErrors {
			errorsMap[validationErr.Field()] = validationErr.Field() + " is " + validationErr.ActualTag()
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": errorsMap})
		return
	}

	books = append(books, bookInput)

	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func GetDetailBookHandler(c *gin.Context) {
	id := c.Param("book_id")
	parseId, _ := strconv.Atoi(id)

	for _, book := range books {
		if book.Id == parseId {
			c.JSON(http.StatusOK, gin.H{
				"data": book,
			})
			return
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{"errors": "book not found"})

}

func DeleteBookHandler(c *gin.Context) {
	id := c.Param("book_id")
	parseId, _ := strconv.Atoi(id)

	filteredBook := []book.BookInput{}

	for _, book := range books {
		if book.Id != parseId {
			filteredBook = append(filteredBook, book)
		}
	}

	books = filteredBook

	c.JSON(http.StatusOK, gin.H{"message": "success delete book"})

}

func UpdateBookHandler(c *gin.Context) {
	id := c.Param("book_id")
	parseId, _ := strconv.Atoi(id)

	var bookInput book.BookInput
	validate := validator.New()

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := validate.Struct(bookInput); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorsMap := make(map[string]string)
		for _, validationErr := range validationErrors {
			errorsMap[validationErr.Field()] = validationErr.Field() + " is " + validationErr.ActualTag()
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": errorsMap})
		return
	}

	for idx := range books {
		if books[idx].Id == parseId {
			books[idx].Title = bookInput.Title
			books[idx].SubTitle = bookInput.SubTitle
			books[idx].Price = bookInput.Price
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "success update book"})

}
