package main

import (
	"fmt"
	"net/http"
)

func handlePing(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("pong\n"))
	fmt.Println("Ping request received, answered with pong")
}

func main() {
	http.HandleFunc("/ping", handlePing)
	fmt.Println("Server started at port 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error starting server: ", err)
		return
	}
}
