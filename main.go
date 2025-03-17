package main

import (
	"ecommerce/config"
	"ecommerce/database"
	"ecommerce/routes"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()

	// Initialize DB connection
	database.InitDB()
	defer database.CloseDB()

	// Define routes
	routes.InitializeRoutes()

	// Start server
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
