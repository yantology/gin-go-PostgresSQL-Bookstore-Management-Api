package configs

import (
	"database/sql"

	"github.com/yantology/gin-go-PostgresSQL-Bookstore-Management-Api/pkg/config/app_config"
	"github.com/yantology/gin-go-PostgresSQL-Bookstore-Management-Api/pkg/config/db_config"
)

func InitConfig() {
	app_config.InitAppConfig()
	db_config.InitDatabaseConfig()
	db_config.ConnectDatabase(sql.Open)
}
