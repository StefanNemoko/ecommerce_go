package database

import (
	"database/sql"
	"ecommerce/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"log"
)

var DB *sql.DB

// InitDB initializes the database connection
func InitDB() {
	var err error

	// Use the config.Config to access the configuration values
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.Config.DBUser, config.Config.DBPassword, config.Config.DBHost, config.Config.DBPort, config.Config.DBName)

	// Establish the connection
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	// Ping the database to check if the connection is successful
	if err := DB.Ping(); err != nil {
		log.Fatal("Error pinging the database: ", err)
	}

	// Successful connection message
	fmt.Println("Database connected!")
}

// CloseDB closes the database connection
func CloseDB() {
	err := DB.Close()
	if err != nil {
		log.Fatal("Error closing the database: ", err)
	}
}
