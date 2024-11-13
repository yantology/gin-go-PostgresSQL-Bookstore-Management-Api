package bookservices

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateBook(t *testing.T) {
	tests := []struct {
		name    string
		book    BookRequest
		wantErr bool
	}{
		{
			name: "CreateBook_Success",
			book: BookRequest{
				Name:        "Test Book",
				Author:      "Test Author",
				Publication: "Test Publication",
			},
			wantErr: false,
		},
		{
			name: "CreateBook_Failure",
			book: BookRequest{
				Name:        "",
				Author:      "Test Author",
				Publication: "Test Publication",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			bsp := NewBookServicesPostgres(db)

			if !tt.wantErr {
				mock.ExpectQuery("INSERT INTO books").
					WithArgs(tt.book.Name, tt.book.Author, tt.book.Publication, sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "author", "publication", "created_at", "updated_at"}).
						AddRow(1, tt.book.Name, tt.book.Author, tt.book.Publication, time.Now(), time.Now()))
			} else {
				mock.ExpectQuery("INSERT INTO books").
					WithArgs(tt.book.Name, tt.book.Author, tt.book.Publication, sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(errors.New("insert error"))
			}

			bookResponse, err := bsp.CreateBook(tt.book)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.book.Name, bookResponse.Name)
				assert.Equal(t, tt.book.Author, bookResponse.Author)
				assert.Equal(t, tt.book.Publication, bookResponse.Publication)
			}
		})
	}
}

func TestGetAllBooks(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "GetAllBooks_Success",
			wantErr: false,
		},
		{
			name:    "GetAllBooks_Failure",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			bsp := NewBookServicesPostgres(db)

			if !tt.wantErr {
				mock.ExpectQuery("SELECT id, name, author, publication, created_at, updated_at FROM books").
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "author", "publication", "created_at", "updated_at"}).
						AddRow(1, "Test Book", "Test Author", "Test Publication", time.Now(), time.Now()))
			} else {
				mock.ExpectQuery("SELECT id, name, author, publication, created_at, updated_at FROM books").
					WillReturnError(errors.New("select error"))
			}

			books, err := bsp.GetAllBooks()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Len(t, books, 1)
				assert.Equal(t, "Test Book", books[0].Name)
			}
		})
	}
}

func TestGetBookByID(t *testing.T) {
	tests := []struct {
		name    string
		bookID  string
		wantErr bool
	}{
		{
			name:    "GetBookByID_Success",
			bookID:  "1",
			wantErr: false,
		},
		{
			name:    "GetBookByID_Failure",
			bookID:  "2",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			bsp := NewBookServicesPostgres(db)

			if !tt.wantErr {
				mock.ExpectQuery("SELECT id, name, author, publication, created_at, updated_at FROM books WHERE id = \\$1").
					WithArgs(tt.bookID).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "author", "publication", "created_at", "updated_at"}).
						AddRow(1, "Test Book", "Test Author", "Test Publication", time.Now(), time.Now()))
			} else {
				mock.ExpectQuery("SELECT id, name, author, publication, created_at, updated_at FROM books WHERE id = \\$1").
					WithArgs(tt.bookID).
					WillReturnError(errors.New("select error"))
			}

			book, err := bsp.GetBookByID(tt.bookID)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, "Test Book", book.Name)
			}
		})
	}
}

func TestUpdateBookByID(t *testing.T) {
	tests := []struct {
		name    string
		bookID  string
		book    BookUpdateRequest
		wantErr bool
	}{
		{
			name:   "UpdateBookByID_Success",
			bookID: "1",
			book: BookUpdateRequest{
				Name:        "Updated Book",
				Author:      "Updated Author",
				Publication: "Updated Publication",
			},
			wantErr: false,
		},
		{
			name:   "UpdateBookByID_Failure",
			bookID: "2",
			book: BookUpdateRequest{
				Name:        "Updated Book",
				Author:      "Updated Author",
				Publication: "Updated Publication",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			bsp := NewBookServicesPostgres(db)

			if !tt.wantErr {
				mock.ExpectQuery("UPDATE books SET").
					WithArgs(tt.book.Name, tt.book.Author, tt.book.Publication, sqlmock.AnyArg(), tt.bookID).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "author", "publication", "created_at", "updated_at"}).
						AddRow(1, tt.book.Name, tt.book.Author, tt.book.Publication, time.Now(), time.Now()))
			} else {
				mock.ExpectQuery("UPDATE books SET").
					WithArgs(tt.book.Name, tt.book.Author, tt.book.Publication, sqlmock.AnyArg(), tt.bookID).
					WillReturnError(errors.New("update error"))
			}

			bookResponse, err := bsp.UpdateBookByID(tt.bookID, tt.book)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.book.Name, bookResponse.Name)
				assert.Equal(t, tt.book.Author, bookResponse.Author)
				assert.Equal(t, tt.book.Publication, bookResponse.Publication)
			}
		})
	}
}

func TestDeleteBookByID(t *testing.T) {
	tests := []struct {
		name    string
		bookID  string
		wantErr bool
	}{
		{
			name:    "DeleteBookByID_Success",
			bookID:  "1",
			wantErr: false,
		},
		{
			name:    "DeleteBookByID_Failure",
			bookID:  "2",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			bsp := NewBookServicesPostgres(db)

			if !tt.wantErr {
				mock.ExpectExec("DELETE FROM books WHERE id = \\$1").
					WithArgs(tt.bookID).
					WillReturnResult(sqlmock.NewResult(1, 1))
			} else {
				mock.ExpectExec("DELETE FROM books WHERE id = \\$1").
					WithArgs(tt.bookID).
					WillReturnError(errors.New("delete error"))
			}

			err = bsp.DeleteBookByID(tt.bookID)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
