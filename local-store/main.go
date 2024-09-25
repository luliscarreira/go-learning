package main

import (
	"fmt"
	"local-store/db"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome, please login or register to continue.")
	userName := handleUser()
	fmt.Printf("Welcome to your store %s!\n", userName)
	loadedDb := db.DB{}
	err := loadedDb.ReadDatabase()
	if err != nil {
		panic(err)
	}

	fmt.Println("Store loaded.")
	fmt.Println(loadedDb.GetProducts())

	// TODO: Implement make a purchase
	operations := `Choose an operation to perform:
    1. Add item
    2. Delete item
    3. Update item
    4. List items
    5. Exit
    Enter your choice: `

	for {
		operation := promptToUserAndScanIntInput(operations)
		fmt.Println(operation)

		switch operation {
		case 1:
			fmt.Println("Operation add item selected")
			productName := promptToUserAndScanStringInput("Enter product name: ")
			productPrice := promptToUserAndScanFloatInput("Enter product price: ")
			productQuantity := promptToUserAndScanIntInput("Enter product quantity: ")

			newProduct := db.Product{
				Name:     productName,
				Price:    productPrice,
				Quantity: int16(productQuantity),
			}

			loadedDb.AddProduct(newProduct)
			fmt.Println("Product added: ", newProduct)

			continue
		case 2:
			fmt.Println("Operation delete item selected")
			productName := promptToUserAndScanStringInput("Enter product name: ")
			loadedDb.RemoveProduct(productName)
			fmt.Println("Product removed: ", productName)

			continue
		case 3:
			// TODO: Implement proper validation for updating product
			fmt.Println("Operation update item selected")
			productName := promptToUserAndScanStringInput("Enter product to update: ")
			newProductName := promptToUserAndScanStringInput("Enter new product name: ")
			productPrice := promptToUserAndScanFloatInput("Enter product price: ")
			productQuantity := promptToUserAndScanIntInput("Enter product quantity: ")

			updatedProduct := db.Product{
				Name:     newProductName,
				Price:    productPrice,
				Quantity: int16(productQuantity),
			}

			loadedDb.UpdateProduct(productName, updatedProduct)
			fmt.Println("Product updated: ", updatedProduct)

			continue
		case 4:
			fmt.Println("Operation list items selected")
			fmt.Println(loadedDb.GetProducts())

			continue
		default:
			fmt.Println("Goodbye!")

			return
		}
	}
}

// TODO: Create a package for handling user
func handleUser() string {
	userName := promptToUserAndScanStringInput("Enter your username: ")
	userPassword, err := readUserPasswordFromFile(userName)
	if err != nil {
		fmt.Println("Registering with given username...")
		return registerUser(userName)
	}

	return login(userName, userPassword)
}

func promptToUserAndScanStringInput(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scan(&input)

	return input
}

func promptToUserAndScanFloatInput(prompt string) float64 {
	var input float64
	fmt.Print(prompt)
	fmt.Scan(&input)

	return input
}

func promptToUserAndScanIntInput(prompt string) int {
	var input int
	fmt.Print(prompt)
	fmt.Scan(&input)

	return input
}

func readUserPasswordFromFile(userName string) (string, error) {
	userFile := fmt.Sprintf("users/%s", strings.ToLower(userName))
	userPassword, err := os.ReadFile(userFile)
	if err != nil {
		fmt.Println("User not found!")
		return "", err
	}

	return string(userPassword), nil
}

func registerUser(userName string) string {
	userPassword := promptToUserAndScanStringInput("Enter your password: ")
	userFile := fmt.Sprintf("users/%s", strings.ToLower(userName))
	err := os.WriteFile(userFile, []byte(userPassword), 0644)
	if err != nil {
		fmt.Println("Error registering user.")
		panic(err)
	}

	return userName
}

func login(userName, userPassword string) string {
	for {
		password := promptToUserAndScanStringInput("Enter your password: ")
		if password != userPassword {
			fmt.Println("Invalid password, try again...")
			continue
		}

		break
	}

	return userName
}
