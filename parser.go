package main

// https://eager.io/blog/go-and-json/
// https://blog.golang.org/json-and-go

import (
	"encoding/json"
	"fmt"
)

func main() {

	unmarshalApp()
	// unmarshalAppDictData()
}

// structs

type App struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type Mymap struct {
	app App `json:"app"`
}

// unmarshal methods

func unmarshalApp() {

	// backticks for raw string, preserves whitespace
	appData := []byte(`
	{ "id": "k34rAT4",
	"title": "My Awesome App"
}
`)

	// variable name type
	var app App

	err := json.Unmarshal(appData, &app)
	if err != nil {
		fmt.Println("error %d\n", err)
	}
	fmt.Println("app ", app)
	fmt.Println("app.Id ", app.Id)
	fmt.Println("app.Title ", app.Title)
}

// TODO: FIXME
func unmarshalAppDictData() {

	appDictData := []byte(`
	{ "app": {
		"id": "k34rAT4",
		"title": "My Awesome App"
	}
}
`)

	// variable name type
	var mymap Mymap

	err := json.Unmarshal(appDictData, &mymap)
	if err != nil {
		fmt.Println("error %d\n", err)
	}
	fmt.Println("mymap ", mymap)
	fmt.Println("mymap.app ", mymap.app)
}
