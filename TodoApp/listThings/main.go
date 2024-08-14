package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Item struct {
	Number    int    `json:"number"`
	ItemName  string `json:"item_name"`
	Completed bool   `json:"completed"`
}

func writeToFile(content, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("File written successfully.")
}

func createList(things ...Item) []Item {
	var listOfThings []Item
	for _, thing := range things {
		fmt.Printf("Adding: %v\n", thing)
		listOfThings = append(listOfThings, thing)
	}
	return listOfThings
}

func convertToJsonString(items []Item) string {
	b, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return ""
	}
	return string(b)
}

func jsonToGo(input string) {
	var obj any
	err := json.Unmarshal([]byte(input), &obj)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	m := obj.(map[string]any)
	fmt.Println("Name:", m["name"])
	fmt.Println("Age:", m["age"])
}

func main() {

	todo := []Item{
		{Number: 1, ItemName: "laundry", Completed: false},
		{Number: 2, ItemName: "dishes", Completed: true},
		{Number: 3, ItemName: "vacuuming", Completed: false},
		{Number: 4, ItemName: "homework", Completed: false},
		{Number: 5, ItemName: "brush teeth", Completed: true},
		{Number: 6, ItemName: "shower", Completed: true},
		{Number: 7, ItemName: "eat breakfast", Completed: false},
		{Number: 8, ItemName: "get the mail", Completed: false},
		{Number: 9, ItemName: "make bed", Completed: true},
		{Number: 10, ItemName: "clean bathroom", Completed: false},
	}

	list := createList(todo...)
	jsonString := convertToJsonString(list)
	fmt.Println(jsonString)
	writeToFile(jsonString, "items.json")
}
