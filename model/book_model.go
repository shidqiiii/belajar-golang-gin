package model

import "time"

type Book struct {
	Id        int       `json:"book_id"`
	Title     string    `json:"title" binding:"required"`
	Price     int       `json:"price" binding:"required"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
