package main

import (
	"bufio"
	"fmt"
	"note-app/note"
	"note-app/todo"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Print()
}

func main() {
	title, content := getNoteDate()
	createdNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error creating note: ", err)

		return
	}

	todoText := getUserInput("Enter todo text: ")
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println("Error creating todo: ", err)

		return
	}

	displayData(createdNote)
	displayData(todo)
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return ""
	}

	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")

	return input
}

func getNoteDate() (string, string) {
	title := getUserInput("Enter note title: ")
	content := getUserInput("Enter note content: ")

	return title, content
}

func displayData(data outputtable) {
	data.Print()
	saveData(data)
}

func saveData(data saver) {
	err := data.Save()
	if err != nil {
		fmt.Println("Error saving data: ", err)

		panic(err)
	}

	fmt.Println("Data saved successfully!")
}
