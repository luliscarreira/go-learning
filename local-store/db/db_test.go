package db

import (
	"testing"
)

// TODO: create test file or mock the file/behaviour
func TestReadDatabase(t *testing.T) {
	db := DB{}
	err := db.ReadDatabase()
	if err != nil {
		t.Errorf("Error creating database: %v", err)
	}

	products := len(db.products)
	if products == 0 {
		t.Errorf("Expected products to be loaded")
	}
}

func TestGetProducts(t *testing.T) {
	db := DB{}
	err := db.ReadDatabase()
	if err != nil {
		t.Errorf("Error creating database: %v", err)
	}

	products := db.GetProducts()
	if len(products) == 0 {
		t.Errorf("Expected products to be loaded")
	}
}

func TestAddProduct(t *testing.T) {
	db := DB{}
	err := db.ReadDatabase()
	if err != nil {
		t.Errorf("Error creating database: %v", err)
	}

	products := db.GetProducts()
	if len(products) != 2 {
		t.Errorf("Expected products to be loaded")
	}

	product := Product{
		Name:     "salgadinho",
		Price:    3.75,
		Quantity: 15,
	}

	db.AddProduct(product)
	products = db.GetProducts()
	if len(products) != 3 {
		t.Errorf("Expected product to be added")
	}
}
