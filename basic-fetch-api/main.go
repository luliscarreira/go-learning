package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	var posts []Post
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		fmt.Println("Error fetching data")
		return
	}

	defer response.Body.Close()
	resp, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response")
		return
	}

	json.Unmarshal(resp, &posts)
	for _, post := range posts {
		prettyPost, _ := json.MarshalIndent(post, "", "    ")
		fmt.Println(string(prettyPost))
	}
}
