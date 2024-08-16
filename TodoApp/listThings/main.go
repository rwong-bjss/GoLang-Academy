package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
)

type Item struct {
	Number    int    `json:"number"`
	ItemName  string `json:"item_name"`
	Completed bool   `json:"completed"`
}

type Items struct {
	Items []Item `json:"items"`
}

func readFromFile(filename string) {
	fileContent, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	fmt.Println("the file is opened successfully...")
	defer fileContent.Close()

	byteResult, err := io.ReadAll(fileContent)
	if err != nil {
		fmt.Printf("failed to read the file. error: %s\n", err)
		return
	}
	var itemList Items

	err = json.Unmarshal(byteResult, &itemList)
	if err != nil {
		fmt.Printf("there was an error decoding the json. err = %s", err)
		return
	}

	fmt.Println(itemList)

}

func writeToFile(content []byte, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("error creating file:", err)
		return
	}
	defer file.Close()
	err = os.WriteFile(filename, content, 0644)
	if err != nil {
		fmt.Println("error writing to file:", err)
		return
	}
	fmt.Println("file written successfully.")
}

func createList(things ...Item) Items {
	var listOfThings []Item
	for _, thing := range things {
		fmt.Printf("Adding: %v\n", thing)
		listOfThings = append(listOfThings, thing)
	}
	return Items{Items: listOfThings}
}

func convertToJsonString(items Items) []byte {
	b, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return []byte("")
	}
	return b
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
	bytes := convertToJsonString(list)
	fmt.Println(string(bytes))
	concurrentList(list)
	//writeToFile(bytes, "items.json")
	//readFromFile("items.json")
}

func concurrentList(list Items) {
	var wg sync.WaitGroup
	//Add(int) adds delta, which may be negative, to the WaitGroup counter.
	wg.Add(2)

	itemChannel := make(chan Item)
	statusChannel := make(chan bool)

	go func() {
		//Done() decrements the WaitGroup counter by one.
		defer wg.Done()
		for _, item := range list.Items {
			itemChannel <- item
			fmt.Printf("item: %s\n", item.ItemName)
			<-statusChannel
		}
		close(itemChannel)
	}()

	go func() {
		defer wg.Done()
		for item := range itemChannel {
			fmt.Printf("status: %t\n", item.Completed)
			statusChannel <- item.Completed
		}
		close(statusChannel)
	}()
	//Wait() blocks until the WaitGroup counter is zero.
	wg.Wait()
}
