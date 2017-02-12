package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	unmarshalApp()
	// unmarshalAppDictData()

	unmarshalJsonMapStringString()
	unmarshalJsonMapStringInterface()
	unmarshalJsonArray()
	unmarshalJsonDict()
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

func unmarshalJsonMapStringString() {

	// backticks for raw string, preserves whitespace
	jsonData := []byte(`
	{ "name": "Homer Simpson"
}
`)

	// variable name type
	var obj map[string]string

	err := json.Unmarshal(jsonData, &obj)
	if err != nil {
		fmt.Println("error %d\n", err)
		panic(err)
	}
	fmt.Println("\n")
	fmt.Println("obj ", obj)
	//fmt.Println("obj.name ", obj.name)
}

func unmarshalJsonMapStringInterface() {

	// backticks for raw string, preserves whitespace
	jsonData := []byte(`
	{ "name": "Homer Simpson",
	"age": 42
}
`)

	// interface{} is an empty interface,
	// can take any type e.g. string or int
	var obj map[string]interface{}

	err := json.Unmarshal(jsonData, &obj)
	if err != nil {
		fmt.Println("error %d\n", err)
		panic(err)
	}
	fmt.Println("obj ", obj)
}

// example from godoc
func unmarshalJsonArray() {

	// backticks for raw string, preserves whitespace
	jsonBlob := []byte(`[
	{ "Name": "Platypus", "Order": "Monotremata" },
	{ "Name": "Quoll", "Order": "Dasyurus" }
	]`)

	type Animal struct {
		Name  string
		Order string
	}

	var animals []Animal

	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("\n")
	fmt.Printf("%+v", animals)
}

func unmarshalJsonDict() {

	// backticks for raw string, preserves whitespace
	jsonBytes := []byte(`{
		"id": 1,
		"name": "reds",
		"colors": ["Crimson", "Red", "Ruby", "Maroon"]
	}`)

	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}

	var colorGroup ColorGroup

	err := json.Unmarshal(jsonBytes, &colorGroup)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("\n")
	fmt.Printf("%+v", colorGroup)
}
