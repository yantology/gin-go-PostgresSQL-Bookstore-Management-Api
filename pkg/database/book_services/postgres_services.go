package bookservices

import (
	"database/sql"
	"errors"
	"time"
)

type BookServicesPostgres struct {
	DB *sql.DB
}

func NewBookServicesPostgres(db *sql.DB) *BookServicesPostgres {
	return &BookServicesPostgres{
		DB: db,
	}
}

func (bsp *BookServicesPostgres) CreateBook(book BookRequest) (BookResponse, error) {
	var bookResponse BookResponse
	query := "INSERT INTO books (name, author, publication, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, name, author, publication, created_at, updated_at"
	err := bsp.DB.QueryRow(query, book.Name, book.Author, book.Publication, time.Now(), time.Now()).Scan(&bookResponse.ID, &bookResponse.Name, &bookResponse.Author, &bookResponse.Publication, &bookResponse.CreatedAt, &bookResponse.UpdatedAt)
	if err != nil {
		return BookResponse{}, err
	}
	return bookResponse, nil
}

func (bsp *BookServicesPostgres) GetAllBooks() ([]BookResponse, error) {
	var books []BookResponse
	rows, err := bsp.DB.Query("SELECT id, name, author, publication, created_at, updated_at FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book BookResponse
		if err := rows.Scan(&book.ID, &book.Name, &book.Author, &book.Publication, &book.CreatedAt, &book.UpdatedAt); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (bsp *BookServicesPostgres) GetBookByID(bookID string) (BookResponse, error) {
	var book BookResponse
	query := "SELECT id, name, author, publication, created_at, updated_at FROM books WHERE id = $1"
	err := bsp.DB.QueryRow(query, bookID).Scan(&book.ID, &book.Name, &book.Author, &book.Publication, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return BookResponse{}, errors.New("book not found")
		}
		return BookResponse{}, err
	}
	return book, nil
}

func (bsp *BookServicesPostgres) UpdateBookByID(bookID string, book BookUpdateRequest) (BookResponse, error) {
	var bookResponse BookResponse
	query := "UPDATE books SET name = $1, author = $2, publication = $3, updated_at = $4 WHERE id = $5 RETURNING id, name, author, publication, created_at, updated_at"
	err := bsp.DB.QueryRow(query, book.Name, book.Author, book.Publication, time.Now(), bookID).Scan(&bookResponse.ID, &bookResponse.Name, &bookResponse.Author, &bookResponse.Publication, &bookResponse.CreatedAt, &bookResponse.UpdatedAt)
	if err != nil {
		return BookResponse{}, err
	}
	return bookResponse, nil
}

func (bsp *BookServicesPostgres) DeleteBookByID(bookID string) error {
	query := "DELETE FROM books WHERE id = $1"
	_, err := bsp.DB.Exec(query, bookID)
	if err != nil {
		return err
	}
	return nil
}
