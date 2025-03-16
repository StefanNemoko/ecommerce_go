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
	http.HandleFunc("/api/products/{id}",
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodGet {
				handlers.GetProductHandler(w, r) // GET
			} else if r.Method == http.MethodPatch {
				handlers.PatchProductHandler(w, r) // PATCH
			}
		})

	// Start server
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
