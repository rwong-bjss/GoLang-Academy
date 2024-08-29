package main

import (
	database "GoLang-Academy/TodoApp/Database"
	"fmt"
	"time"
)

//19. Remote Commands - Create a server that can concurrently receive a list of pre-defined commands,
//The status of the server and the status of each task should be available via specific commands. [Concurrency][Goroutines] [Channels]

var requests chan operation = make(chan operation)
var done chan struct{} = make(chan struct{})

// this is the actor
type operation struct {
	//what we want to do
	action string
	//data passed in
	item     string
	response chan string
}

func Start() {
	go monitorRequests()
}

func Stop() {
	shutdown := operation{action: "shutdown", item: "", response: nil}
	requests <- shutdown
	<-done
}

// OrderCoffee function
func requestDbOperation(action, data string) {
	op := operation{
		action: action,
		item:   data,
		//response: make(chan string),
	}
	requests <- op
	//fmt.Println("This is the object:" + data)
}

func monitorRequests() {
	lastObjectWritten := "nothing yet"
	status := "available"
	for op := range requests {
		//delete doesn't work in our case because it's delete by ID and the way items are generated it allocates them an ID based on the size of the map.

		switch op.action {
		case "write":
			status = "busy"
			lastObjectWritten = op.item
			post(op.item)
			status = "available"
		case "read":
			get()
		case "status":
			op.response <- status
		case "lastObject":
			op.response <- lastObjectWritten
		case "shutdown":
			fmt.Println("Shutting down")
			close(requests)
		}
	}
	// Signal all requests completed
	fmt.Println("All requests processed")
	close(done)
}

func post(data string) {
	fmt.Println("Writing to the db here: " + data)
	id := len(db.Items) + 1
	initialItem := &database.Item{
		Id:       id,
		ItemName: data,
		Status:   false,
	}
	err := database.InsertItem(db, initialItem)
	if err != nil {
		return
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Object has been written: " + data)
	fmt.Println(database.GetAllItems(db))
}

func get() {
	database.GetAllItems(db)
}

func getLastObject() string {
	answer := make(chan string)
	op := operation{action: "lastObject",
		item:     "",
		response: answer}
	requests <- op
	return <-answer
}

func getStatus() string {
	answer := make(chan string)
	op := operation{action: "status",
		item:     "",
		response: answer}
	requests <- op
	return <-answer
}

var db = database.CreateDatabase()

func main() {
	//db := database.CreateDatabase()
	Start()
	defer Stop()

	fmt.Println(db.Items)

	go requestDbOperation("write", "Laundry")
	go requestDbOperation("write", "Washing")
	go fmt.Println("Item written: " + getLastObject())
	go requestDbOperation("write", "Shower")
	go fmt.Println("Item written: " + getLastObject())
	go requestDbOperation("get", "")
	go requestDbOperation("status", getStatus())
	go fmt.Println("Item written: " + getLastObject())
	//fmt.Println(database.GetAllItems(db))
}
