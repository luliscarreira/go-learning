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
	db := db.DB{}
	err := db.ReadDatabase()
	if err != nil {
		panic(err)
	}

	fmt.Println("Store loaded.")
	fmt.Println(db.GetProducts())

	// Missing make a purchase
	operations := `Choose an operation to perform:
    1. Add item
    2. Delete item
    3. Update item
    4. List items
    5. Exit
    Enter your choice: `

	operation := promptToUserAndScanIntInput(operations)

	fmt.Println(operation)

	// User can add, delete, update, or list items and exit
	// User can make a purchase that is persisted to the database
	// if user chooses to add, ask for item name and price and quantity
	// write to database
	// if user chooses to delete, ask for item name
	// delete item from database
	// if user chooses to update, ask for item name and new price and quantity
	// update item in database
	// if user chooses to list, list all items
}

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
