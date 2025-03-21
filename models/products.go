package models

import (
	"database/sql"
	"ecommerce/database"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	StatusActive   = "active"
	StatusInactive = "inactive"
)

// Product represents the user entity in the database
type Product struct {
	ID           int64        `json:"id"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	Status       string       `json:"status"`
	Price        float32      `json:"price"`
	Tax          float32      `json:"tax"`
	Discount     float32      `json:"discount"`
	DiscountType string       `json:"discount_type"`
	Stock        int          `json:"stock"`
	Sku          string       `json:"sku"`
	CreatedAt    sql.NullTime `json:"created_at,omitempty"`
	UpdatedAt    sql.NullTime `json:"updated_at,omitempty"`
	DeletedAt    sql.NullTime `json:"deleted_at,omitempty"`
}

func (p *Product) validate() error {
	if p.Name == "" {
		return errors.New("product name is required")
	}
	if p.Description == "" {
		return errors.New("product description is required")
	}

	validStatuses := []string{StatusActive, StatusInactive}
	isValidStatus := false
	for _, status := range validStatuses {
		if strings.ToLower(p.Status) == status {
			isValidStatus = true
		}
	}
	if !isValidStatus {
		return errors.New("invalid product status")
	}

	return nil
}

// SaveProduct inserts a new product into the database
func (p *Product) SaveProduct() (Product, error) {
	if err := p.validate(); err != nil {
		return Product{}, fmt.Errorf("Validation error: %s", err)
	}

	var query string
	if p.ID == 0 {
		query = fmt.Sprintf("INSERT INTO products (name, description, status, price, tax, discount, discount_type, stock, sku, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, '%s', '%s')", time.Now(), time.Now())
	} else {
		query = fmt.Sprintf("UPDATE products SET name = ?, description = ?, status = ?, price = ?, tax = ?, discount = ?, discount_type = ?, stock = ?, sku = ?, updated_at = '%s' WHERE ID = '%d'", time.Now(), p.ID)
	}
	result, err := database.DB.Exec(query, p.Name, p.Description, p.Status, p.Price, p.Tax, p.Discount, p.DiscountType, p.Stock, p.Sku)
	if err != nil {
		return Product{}, fmt.Errorf("error creating product: %w", err)
	}

	// Just updating the product? no need to retrieve last inserted id.
	if p.ID > 0 {
		return GetProductByID(p.ID)
	}

	// Retrieve last inserted ID and retrieve product.
	insertedID, err := result.LastInsertId()
	if err != nil {
		return Product{}, fmt.Errorf("Error fetching inserted ID", http.StatusInternalServerError)
	}

	return GetProductByID(insertedID)
}

// GetProductByID retrieves a product by ID
func GetProductByID(id int64) (Product, error) {
	var product Product
	query := "SELECT id, name, description, status, price, tax, discount, discount_type, stock, sku, created_at, updated_at, deleted_at FROM products WHERE id = ?"
	row := database.DB.QueryRow(query, id)

	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Status, &product.Price, &product.Tax, &product.Discount, &product.DiscountType, &product.Stock, &product.Sku, &product.CreatedAt, &product.UpdatedAt, &product.DeletedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return product, fmt.Errorf("product not found")
		}
		return product, fmt.Errorf("error retrieving product: %w", err)
	}

	return product, nil
}

func GetProducts() ([]Product, error) {
	// Retrieve rows
	var products []Product
	query := "SELECT * FROM products WHERE Status = ? ORDER BY created_at DESC"
	rows, err := database.DB.Query(query, StatusActive)
	if err != nil {
		return products, fmt.Errorf("error retrieving products: %w", err)
	}

	for rows.Next() {
		// Convert row to models.Product
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Status, &product.Price, &product.Tax, &product.Discount, &product.DiscountType, &product.Stock, &product.Sku, &product.CreatedAt, &product.UpdatedAt, &product.DeletedAt)
		if err != nil {
			return products, fmt.Errorf("error retrieving products: %w", err)
		}

		// append model to collection.
		products = append(products, product)
	}
	return products, nil
}
