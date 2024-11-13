package bookservices

type BookServicesInterface interface {
	CreateBook(book BookRequest) (BookResponse, error)
	GetAllBooks() ([]BookResponse, error)
	GetBookByID(bookID string) (BookResponse, error)
	UpdateBookByID(bookID string, book BookUpdateRequest) (BookResponse, error)
	DeleteBookByID(bookID string) error
}
