package bookservices

func NewBookServicesRepository(bs BookServicesInterface) *BookServicesRepository {
	return &BookServicesRepository{
		BookServices: bs,
	}
}

type BookServicesRepository struct {
	BookServices BookServicesInterface
}

func (bsr *BookServicesRepository) CreateBook(book BookRequest) (BookResponse, error) {
	return bsr.BookServices.CreateBook(book)
}

func (bsr *BookServicesRepository) GetAllBooks() ([]BookResponse, error) {
	return bsr.BookServices.GetAllBooks()
}

func (bsr *BookServicesRepository) GetBookByID(bookID string) (BookResponse, error) {
	return bsr.BookServices.GetBookByID(bookID)
}

func (bsr *BookServicesRepository) UpdateBookByID(bookID string, book BookUpdateRequest) (BookResponse, error) {
	return bsr.BookServices.UpdateBookByID(bookID, book)
}

func (bsr *BookServicesRepository) DeleteBookByID(bookID string) error {
	return bsr.BookServices.DeleteBookByID(bookID)
}
