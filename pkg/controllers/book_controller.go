package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	bookservices "github.com/yantology/gin-go-PostgresSQL-Bookstore-Management-Api/pkg/database/book_services"
)

type BookController struct {
	BookService bookservices.BookServicesInterface
}

func NewBookController(bookService bookservices.BookServicesInterface) *BookController {
	return &BookController{
		BookService: bookService,
	}
}

func (bc *BookController) GetAllBooks(c *gin.Context) {
	books, err := bc.BookService.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (bc *BookController) GetBookByID(c *gin.Context) {
	bookID := c.Param("bookID")
	book, err := bc.BookService.GetBookByID(bookID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (bc *BookController) CreateBook(c *gin.Context) {
	var bookRequest bookservices.BookRequest
	if err := c.ShouldBindJSON(&bookRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book, err := bc.BookService.CreateBook(bookRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (bc *BookController) UpdateBookByID(c *gin.Context) {
	bookID := c.Param("bookID")
	var bookUpdateRequest bookservices.BookUpdateRequest
	if err := c.ShouldBindJSON(&bookUpdateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book, err := bc.BookService.UpdateBookByID(bookID, bookUpdateRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (bc *BookController) DeleteBookByID(c *gin.Context) {
	bookID := c.Param("bookID")
	err := bc.BookService.DeleteBookByID(bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
