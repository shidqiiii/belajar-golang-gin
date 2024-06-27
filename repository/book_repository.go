package repository

import (
	"belajar-golang-gin/model"
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type BookRepository interface {
	FindAll() ([]model.Book, error)
	FindById(id int) (model.Book, error)
	CreateNewBook(book *model.Book) error
	DeleteBook(id int) error
}

type bookRepository struct {
	conn *pgx.Conn // akses ke db
}

func NewBookRepository(conn *pgx.Conn) *bookRepository {
	return &bookRepository{conn}
}

func (r *bookRepository) FindAll() ([]model.Book, error) {
	rows, err := r.conn.Query(context.Background(), "SELECT * FROM books")
	if err != nil {
		return nil, err
	}

	var books []model.Book
	for rows.Next() {
		var book model.Book
		if err := rows.Scan(&book.Id, &book.Title, &book.Price, &book.CreatedAt, &book.UpdatedAt); err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (r *bookRepository) FindById(id int) (model.Book, error) {
	var book model.Book
	err := r.conn.QueryRow(context.Background(), "SELECT * FROM books WHERE book_id = $1", id).Scan(&book.Id, &book.Title, &book.Price, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		return model.Book{}, err
	}

	return book, nil
}

func (r *bookRepository) CreateNewBook(book *model.Book) error {
	book.CreatedAt = time.Now()
	book.UpdatedAt = book.CreatedAt

	_, err := r.conn.Exec(context.Background(), "INSERT INTO books (title, price, created_at, updated_at) VALUES ($1, $2, $3, $4)", book.Title, book.Price, book.CreatedAt, book.UpdatedAt)
	if err != nil {
		return err
	}
	return nil

}

func (r *bookRepository) DeleteBook(id int) error {
	_, err := r.conn.Exec(context.Background(), "DELETE FROM books where book_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
