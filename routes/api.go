package routes

import (
	"ecommerce/handlers"
	"net/http"
)

func InitializeRoutes() {
	// Define routes
	http.HandleFunc("/api/products",
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "GET" {
				handlers.GetProductsHandler(w) // GET
			} else if r.Method == "POST" {
				handlers.CreateProductHandler(w, r) // POST
			}
		})
	http.HandleFunc("/api/products/{id}",
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodGet {
				handlers.GetProductHandler(w, r) // GET
			} else if r.Method == http.MethodPatch {
				handlers.PatchProductHandler(w, r) // PATCH
			}
		})
}
