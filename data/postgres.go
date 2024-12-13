package data

import (
	"database/sql"
	"fmt"

	"github.com/AmirmahdiShahbazi/clean-web-api/config"
	"github.com/AmirmahdiShahbazi/clean-web-api/pkg/logger"
	_ "github.com/lib/pq"
)

var PostgresClient *sql.DB

func ConnectToPostgres() {  
	potgresConfigs := config.GetPostgresConfigs()  
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+  
		"password=%s dbname=%s sslmode=disable",  
		potgresConfigs.Host, potgresConfigs.Port, potgresConfigs.User, potgresConfigs.Password, potgresConfigs.DBName)  
	var err error  
	PostgresClient, err = sql.Open("postgres", psqlInfo)  
	if err != nil {  
		logger.NewLogger().Fatal(logger.Postgres, logger.Startup, err.Error(), nil)
	}  
	err = PostgresClient.Ping()  
	if err != nil {  
		logger.NewLogger().Fatal(logger.Postgres, logger.Startup, err.Error(), nil)
	}  
}  