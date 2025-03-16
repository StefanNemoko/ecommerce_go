package main

import (
	"ecommerce/config"
	"ecommerce/database"
	"ecommerce/handlers"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()

	// Initialize DB connection
	database.InitDB()
	defer database.CloseDB()

	// Define routes
	http.HandleFunc("/api/products", handlers.CreateProductHandler) // POST
	//TODO:: GET request handelt de id nog niet
	http.HandleFunc("/api/products/{id}", handlers.GetProductHandler) // GET

	// Start server
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
