package bookservices

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockBookServices is a mock implementation of BookServicesInterface
type MockBookServices struct {
	mock.Mock
}

func (m *MockBookServices) CreateBook(book BookRequest) (BookResponse, error) {
	args := m.Called(book)
	return args.Get(0).(BookResponse), args.Error(1)
}

func (m *MockBookServices) GetAllBooks() ([]BookResponse, error) {
	args := m.Called()
	return args.Get(0).([]BookResponse), args.Error(1)
}

func (m *MockBookServices) GetBookByID(bookID string) (BookResponse, error) {
	args := m.Called(bookID)
	return args.Get(0).(BookResponse), args.Error(1)
}

func (m *MockBookServices) UpdateBookByID(bookID string, book BookUpdateRequest) (BookResponse, error) {
	args := m.Called(bookID, book)
	return args.Get(0).(BookResponse), args.Error(1)
}

func (m *MockBookServices) DeleteBookByID(bookID string) error {
	args := m.Called(bookID)
	return args.Error(0)
}

func TestCreateBookRepository(t *testing.T) {
	mockService := new(MockBookServices)
	repo := NewBookServicesRepository(mockService)

	bookRequest := BookRequest{Name: "Test Book", Author: "Test Author", Publication: "Test Publication"}
	bookResponse := BookResponse{ID: 1, Name: "Test Book", Author: "Test Author", Publication: "Test Publication"}

	mockService.On("CreateBook", bookRequest).Return(bookResponse, nil)

	result, err := repo.CreateBook(bookRequest)
	assert.NoError(t, err)
	assert.Equal(t, bookResponse, result)

	mockService.AssertExpectations(t)
}

func TestGetAllBooksRepository(t *testing.T) {
	mockService := new(MockBookServices)
	repo := NewBookServicesRepository(mockService)

	bookResponses := []BookResponse{
		{ID: 1, Name: "Test Book", Author: "Test Author", Publication: "Test Publication"},
	}

	mockService.On("GetAllBooks").Return(bookResponses, nil)

	result, err := repo.GetAllBooks()
	assert.NoError(t, err)
	assert.Equal(t, bookResponses, result)

	mockService.AssertExpectations(t)
}

func TestGetBookByIDRepository(t *testing.T) {
	mockService := new(MockBookServices)
	repo := NewBookServicesRepository(mockService)

	bookID := "1"
	bookResponse := BookResponse{ID: 1, Name: "Test Book", Author: "Test Author", Publication: "Test Publication"}

	mockService.On("GetBookByID", bookID).Return(bookResponse, nil)

	result, err := repo.GetBookByID(bookID)
	assert.NoError(t, err)
	assert.Equal(t, bookResponse, result)

	mockService.AssertExpectations(t)
}

func TestUpdateBookByIDRepository(t *testing.T) {
	mockService := new(MockBookServices)
	repo := NewBookServicesRepository(mockService)

	bookID := "1"
	bookUpdateRequest := BookUpdateRequest{Name: "Updated Book", Author: "Updated Author", Publication: "Updated Publication"}
	bookResponse := BookResponse{ID: 1, Name: "Updated Book", Author: "Updated Author", Publication: "Updated Publication"}

	mockService.On("UpdateBookByID", bookID, bookUpdateRequest).Return(bookResponse, nil)

	result, err := repo.UpdateBookByID(bookID, bookUpdateRequest)
	assert.NoError(t, err)
	assert.Equal(t, bookResponse, result)

	mockService.AssertExpectations(t)
}

func TestDeleteBookByIDRepository(t *testing.T) {
	mockService := new(MockBookServices)
	repo := NewBookServicesRepository(mockService)

	bookID := "1"

	mockService.On("DeleteBookByID", bookID).Return(nil)

	err := repo.DeleteBookByID(bookID)
	assert.NoError(t, err)

	mockService.AssertExpectations(t)
}
