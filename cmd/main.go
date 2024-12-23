package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	configs "github.com/yantology/gin-go-PostgresSQL-Bookstore-Management-Api/pkg/config"
	"github.com/yantology/gin-go-PostgresSQL-Bookstore-Management-Api/pkg/config/app_config"
	"github.com/yantology/gin-go-PostgresSQL-Bookstore-Management-Api/pkg/config/cors_config"
	"github.com/yantology/gin-go-PostgresSQL-Bookstore-Management-Api/pkg/config/db_config"
	"github.com/yantology/gin-go-PostgresSQL-Bookstore-Management-Api/pkg/controllers"
	bookservices "github.com/yantology/gin-go-PostgresSQL-Bookstore-Management-Api/pkg/database/book_services"
	"github.com/yantology/gin-go-PostgresSQL-Bookstore-Management-Api/pkg/routes"
)

func main() {
	// Initialize all configurations
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	configs.InitConfig()

	// Create Gin router
	router := gin.Default()

	// Apply CORS configuration
	router.Use(cors_config.CorsConfig())

	// Initialize services and controllers
	bookService := bookservices.NewBookServicesPostgres(db_config.GetDB())
	bookController := controllers.NewBookController(bookService)

	// Register routes
	routes.RegisterBookRoutes(router, bookController)

	// Serve static files
	router.Static(app_config.PUBLIC_ROUTE, app_config.PUBLIC_ASSETS_DIR)
	router.StaticFile("/", "./public/index.html")

	// Start server
	router.Run(app_config.PORT)
}
