package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	bookservices "github.com/yantology/gin-go-PostgresSQL-Bookstore-Management-Api/pkg/database/book_services"
)

type MockBookService struct {
	mock.Mock
}

func (m *MockBookService) CreateBook(book bookservices.BookRequest) (bookservices.BookResponse, error) {
	args := m.Called(book)
	return args.Get(0).(bookservices.BookResponse), args.Error(1)
}

func (m *MockBookService) GetAllBooks() ([]bookservices.BookResponse, error) {
	args := m.Called()
	return args.Get(0).([]bookservices.BookResponse), args.Error(1)
}

func (m *MockBookService) GetBookByID(bookID string) (bookservices.BookResponse, error) {
	args := m.Called(bookID)
	return args.Get(0).(bookservices.BookResponse), args.Error(1)
}

func (m *MockBookService) UpdateBookByID(bookID string, book bookservices.BookUpdateRequest) (bookservices.BookResponse, error) {
	args := m.Called(bookID, book)
	return args.Get(0).(bookservices.BookResponse), args.Error(1)
}

func (m *MockBookService) DeleteBookByID(bookID string) error {
	args := m.Called(bookID)
	return args.Error(0)
}

func TestGetAllBooks(t *testing.T) {
	tests := []struct {
		name           string
		mockReturn     []bookservices.BookResponse
		mockError      error
		expectedStatus int
		expectedBooks  []bookservices.BookResponse
	}{
		{
			name: "Success",
			mockReturn: []bookservices.BookResponse{{
				ID:          1,
				Name:        "Test Book",
				Author:      "Test Author",
				Publication: "Test Publication",
				CreatedAt:   time.Date(2024, time.November, 14, 5, 30, 9, 142178900, time.Local),
				UpdatedAt:   time.Date(2024, time.November, 14, 5, 30, 9, 142178900, time.Local),
			}},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBooks: []bookservices.BookResponse{{
				ID:          1,
				Name:        "Test Book",
				Author:      "Test Author",
				Publication: "Test Publication",
				CreatedAt:   time.Date(2024, time.November, 14, 5, 30, 9, 142178900, time.Local),
				UpdatedAt:   time.Date(2024, time.November, 14, 5, 30, 9, 142178900, time.Local),
			}},
		},
		{
			name:           "Error",
			mockReturn:     []bookservices.BookResponse{},
			mockError:      errors.New("database error"),
			expectedStatus: http.StatusInternalServerError,
			expectedBooks:  []bookservices.BookResponse{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(MockBookService)
			controller := NewBookController(mockService)

			mockService.On("GetAllBooks").Return(tt.mockReturn, tt.mockError)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			controller.GetAllBooks(c)

			assert.Equal(t, tt.expectedStatus, w.Code)

			var actualBooks []bookservices.BookResponse
			if tt.expectedStatus == http.StatusOK {
				err := json.Unmarshal(w.Body.Bytes(), &actualBooks)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBooks, actualBooks)
			}

			mockService.AssertExpectations(t)
		})
	}
}

func TestGetBookByID(t *testing.T) {
	tests := []struct {
		name           string
		bookID         string
		mockReturn     bookservices.BookResponse
		mockError      error
		expectedStatus int
	}{
		{
			name:   "Success",
			bookID: "1",
			mockReturn: bookservices.BookResponse{
				ID:          1,
				Name:        "Test Book",
				Author:      "Test Author",
				Publication: "Test Publication",
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Not Found",
			bookID:         "999",
			mockReturn:     bookservices.BookResponse{},
			mockError:      errors.New("book not found"),
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(MockBookService)
			controller := NewBookController(mockService)

			mockService.On("GetBookByID", tt.bookID).Return(tt.mockReturn, tt.mockError)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Params = []gin.Param{{Key: "bookID", Value: tt.bookID}}

			controller.GetBookByID(c)

			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusOK {
				var actualBook bookservices.BookResponse
				err := json.Unmarshal(w.Body.Bytes(), &actualBook)
				assert.NoError(t, err)
				assert.Equal(t, tt.mockReturn, actualBook)
			}

			mockService.AssertExpectations(t)
		})
	}
}

func TestCreateBook(t *testing.T) {
	tests := []struct {
		name           string
		bookRequest    interface{}
		mockReturn     bookservices.BookResponse
		mockError      error
		expectedStatus int
	}{
		{
			name: "Success",
			bookRequest: bookservices.BookRequest{
				Name:        "New Book",
				Author:      "New Author",
				Publication: "New Publication",
			},
			mockReturn: bookservices.BookResponse{
				ID:          1,
				Name:        "New Book",
				Author:      "New Author",
				Publication: "New Publication",
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Bad Request",
			bookRequest:    "invalid json",
			mockReturn:     bookservices.BookResponse{},
			mockError:      nil,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Server Error",
			bookRequest: bookservices.BookRequest{
				Name:        "New Book",
				Author:      "New Author",
				Publication: "New Publication",
			},
			mockReturn:     bookservices.BookResponse{},
			mockError:      errors.New("database error"),
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(MockBookService)
			controller := NewBookController(mockService)

			var jsonData []byte
			var err error

			if tt.name == "Bad Request" {
				jsonData = []byte(tt.bookRequest.(string))
			} else {
				jsonData, err = json.Marshal(tt.bookRequest)
				assert.NoError(t, err)
			}

			mockService.On("CreateBook", tt.bookRequest).Return(tt.mockReturn, tt.mockError)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest("POST", "/", bytes.NewBuffer(jsonData))
			c.Request.Header.Set("Content-Type", "application/json")

			controller.CreateBook(c)

			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusOK {
				var actualBook bookservices.BookResponse
				err := json.Unmarshal(w.Body.Bytes(), &actualBook)
				assert.NoError(t, err)
				assert.Equal(t, tt.mockReturn, actualBook)
			}
			if tt.expectedStatus != http.StatusBadRequest {
				mockService.AssertExpectations(t)
			}
		})
	}
}

func TestUpdateBookByID(t *testing.T) {
	tests := []struct {
		name           string
		bookID         string
		updateRequest  bookservices.BookUpdateRequest
		mockReturn     bookservices.BookResponse
		mockError      error
		expectedStatus int
	}{
		{
			name:   "Success",
			bookID: "1",
			updateRequest: bookservices.BookUpdateRequest{
				Name:        "Updated Book",
				Author:      "Updated Author",
				Publication: "Updated Publication",
			},
			mockReturn: bookservices.BookResponse{
				ID:          1,
				Name:        "Updated Book",
				Author:      "Updated Author",
				Publication: "Updated Publication",
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Error",
			bookID:         "1",
			updateRequest:  bookservices.BookUpdateRequest{},
			mockReturn:     bookservices.BookResponse{},
			mockError:      errors.New("update error"),
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(MockBookService)
			controller := NewBookController(mockService)

			mockService.On("UpdateBookByID", tt.bookID, tt.updateRequest).Return(tt.mockReturn, tt.mockError)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Params = []gin.Param{{Key: "bookID", Value: tt.bookID}}

			jsonData, _ := json.Marshal(tt.updateRequest)
			c.Request = httptest.NewRequest("PUT", "/", bytes.NewBuffer(jsonData))
			c.Request.Header.Set("Content-Type", "application/json")

			controller.UpdateBookByID(c)

			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusOK {
				var actualBook bookservices.BookResponse
				err := json.Unmarshal(w.Body.Bytes(), &actualBook)
				assert.NoError(t, err)
				assert.Equal(t, tt.mockReturn, actualBook)
			}

			mockService.AssertExpectations(t)
		})
	}
}

func TestDeleteBookByID(t *testing.T) {
	tests := []struct {
		name           string
		bookID         string
		mockError      error
		expectedStatus int
	}{
		{
			name:           "Success",
			bookID:         "1",
			mockError:      nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Error",
			bookID:         "999",
			mockError:      errors.New("delete error"),
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(MockBookService)
			controller := NewBookController(mockService)

			mockService.On("DeleteBookByID", tt.bookID).Return(tt.mockError)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Params = []gin.Param{{Key: "bookID", Value: tt.bookID}}

			controller.DeleteBookByID(c)

			assert.Equal(t, tt.expectedStatus, w.Code)

			mockService.AssertExpectations(t)
		})
	}
}
