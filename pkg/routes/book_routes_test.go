package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yantology/gin-go-PostgresSQL-Bookstore-Management-Api/pkg/controllers"
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

func TestBookRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockBookService := new(MockBookService)
	bookController := controllers.NewBookController(mockBookService)

	router := gin.Default()
	RegisterBookRoutes(router, bookController)

	tests := []struct {
		method       string
		url          string
		body         string
		mockFunc     func()
		expectedCode int
	}{
		{
			method: "GET",
			url:    "/books/",
			mockFunc: func() {
				mockBookService.On("GetAllBooks").Return([]bookservices.BookResponse{}, nil).Once()
			},
			expectedCode: http.StatusOK,
		},
		{
			method: "GET",
			url:    "/books/1",
			mockFunc: func() {
				mockBookService.On("GetBookByID", "1").Return(bookservices.BookResponse{}, nil).Once()
			},
			expectedCode: http.StatusOK,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		tt.mockFunc()
		req, _ := http.NewRequest(tt.method, tt.url, nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, tt.expectedCode, resp.Code)
		mockBookService.AssertExpectations(t)
	}
}
