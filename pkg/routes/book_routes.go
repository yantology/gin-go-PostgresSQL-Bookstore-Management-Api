package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yantology/gin-go-PostgresSQL-Bookstore-Management-Api/pkg/controllers"
)

func RegisterBookRoutes(router *gin.Engine, bookController *controllers.BookController) {

	bookRoutes := router.Group("/books")
	{
		bookRoutes.GET("/", bookController.GetAllBooks)
		bookRoutes.GET("/:bookID", bookController.GetBookByID)
		bookRoutes.POST("/", bookController.CreateBook)
		bookRoutes.PUT("/:bookID", bookController.UpdateBookByID)
		bookRoutes.DELETE("/:bookID", bookController.DeleteBookByID)
	}

}
