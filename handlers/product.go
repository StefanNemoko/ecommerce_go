package handlers

import (
	"ecommerce/helpers"
	"ecommerce/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Call the method in the Product model to create the product
	product, err = product.SaveProduct()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(product)
	if err != nil {
		return
	}
}

func GetProductsHandler(w http.ResponseWriter) {
	products, err := models.GetProducts()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieveing products: \"%s\"", err), http.StatusBadRequest)
		return
	}

	type Products struct {
		TotalCount int              `json:"total_count"`
		Items      []models.Product `json:"items"`
	}

	response := Products{
		TotalCount: len(products),
		Items:      products,
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.RetrieveIdFromUri(r.URL.Path)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid product ID %s", err), http.StatusBadRequest)
		return
	}

	product, err := models.GetProductByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(product)
	if err != nil {
		return
	}
}

func PatchProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.RetrieveIdFromUri(r.URL.Path)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid product ID %s", err), http.StatusBadRequest)
		return
	}

	product, err := models.GetProductByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}
	product.ID = id

	_, err = product.SaveProduct()
	if err != nil {
		http.Error(w, "Error saving product: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
