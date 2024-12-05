package data

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
	"github.com/AmirmahdiShahbazi/clean-web-api/config"
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
		log.Fatal("Error opening Postgres connection: ", err)  
	}  
	err = PostgresClient.Ping()  
	if err != nil {  
		log.Fatal("Error connecting to Postgres: ", err)  
	}  
}  