package handlers

import (
	"ecommerce/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Status       string  `json:"status"`
	Price        float32 `json:"price"`
	Tax          float32 `json:"tax"`
	Discount     float32 `json:"discount"`
	DiscountType string  `json:"discount_type"`
	Stock        int     `json:"stock"`
	Sku          string  `json:"sku"`
}

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Call the method in the Product model to create the product
	product, err = product.CreateProduct()
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

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
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
