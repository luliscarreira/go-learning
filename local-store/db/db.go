package db

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type DB struct {
	products []Product
}

func (db *DB) CreateDatabase() error {
	// TODO: Cannot clear the database if it already exists
	_, err := os.Create("db/database")
	if err != nil {
		return err
	}

	db.products = []Product{}
	return nil
}

func (db *DB) ReadDatabase() error {
	products := []Product{}
	test, err := os.ReadFile("db/database")
	if err != nil {
		fmt.Println("Error reading database file.")

		return err
	}

	for _, line := range strings.Split(string(test), "\n") {
		productFromDb := strings.Split(line, " ")
		productName := productFromDb[0]
		productPrice, _ := strconv.ParseFloat(productFromDb[1], 64)
		productQuantity, _ := strconv.ParseInt(productFromDb[2], 10, 16)

		product := Product{
			Name:     productName,
			Price:    productPrice,
			Quantity: int16(productQuantity),
		}

		products = append(products, product)
	}

	db.products = products

	return nil
}

func (db *DB) SaveDatabase() {
	file, err := os.OpenFile("db/database", os.O_RDWR, 0755)
	if err != nil {
		fmt.Println("Error opening database file.")
		panic(err)
	}

	defer file.Close()

	productsToWrite := ""
	for _, product := range db.products {
		productLine := fmt.Sprintf("%s %.2f %d\n", product.Name, product.Price, product.Quantity)
		productsToWrite += productLine
	}

	err = file.Truncate(0)
	if err != nil {
		fmt.Println("Error erasing file content.")
		panic(err)
	}

	_, err = file.WriteString(productsToWrite)
	if err != nil {
		fmt.Println("Error writing to file.")
		panic(err)
	}

	fmt.Println("Database saved.")
}

func (db *DB) GetProducts() []Product {
	return db.products
}

func (db *DB) AddProduct(product Product) {
	db.products = append(db.products, product)
	db.SaveDatabase()
}

func (db *DB) RemoveProduct(productName string) {
	for i, product := range db.products {
		if product.Name == productName {
			db.products = append(db.products[:i], db.products[i+1:]...)
			db.SaveDatabase()

			break
		}
	}
}

func (db *DB) UpdateProduct(productName string, product Product) {
	for i, p := range db.products {
		if p.Name == productName {
			db.products[i] = product
			db.SaveDatabase()

			break
		}
	}
}
