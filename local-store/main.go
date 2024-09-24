package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome, please login or register to continue.")
	userName := handleUser()

	fmt.Printf("Welcome to your store %s!\n", userName)

	// read database from local file
	// log to user that the database is read
	// if user logs in successfully, print to user the operations that can be done
	// ask user for the operation to be done
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
	userName := promptToUserAndScanInput("Enter your username: ")
	userPassword, err := readUserPasswordFromFile(userName)
	if err != nil {
		fmt.Println("Registering with given username...")
		return registerUser(userName)
	}

	return login(userName, userPassword)
}

func promptToUserAndScanInput(prompt string) string {
	var input string
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
	userPassword := promptToUserAndScanInput("Enter your password: ")
	userFile := fmt.Sprintf("users/%s", strings.ToLower(userName))
	err := os.WriteFile(userFile, []byte(userPassword), 0644)
	if err != nil {
		fmt.Println("Error registering user.")
		panic(err)
	}

	return userName
}

func login(userName, userPassword string) string {
	for true {
		password := promptToUserAndScanInput("Enter your password: ")
		if password != userPassword {
			fmt.Println("Invalid password, try again...")
			continue
		}

		break
	}

	return userName
}
