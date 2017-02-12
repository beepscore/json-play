package main

// https://eager.io/blog/go-and-json/
// https://blog.golang.org/json-and-go

import (
	"encoding/json"
	"fmt"
)

func main() {

	data := []byte(`
	{
		"id": "k34rAT4",
		"title": "My Awesome App"
	}
	`)

	var app App
	err := json.Unmarshal(data, &app)
	if err != nil {
		fmt.Println("error %d\n", err)
	}
	fmt.Println("app ", app)
	fmt.Println("app.Id ", app.Id)
	fmt.Println("app.Title ", app.Title)
}

type App struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}
