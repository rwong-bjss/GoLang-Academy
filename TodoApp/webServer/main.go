package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	_ "strings"
	"sync"

	database "GoLang-Academy/TodoApp/Database"
)

// Mutex to ensure thread safety
var mu sync.Mutex

// Global database instance
var db = database.CreateDatabase()

// Handler for the homepage
func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	items := database.GetAllItems(db)
	tmpl.Execute(w, items)
}

// Handler to create, update, or delete a to-do item
func manageItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Determine the action based on the `_method` field
	method := r.FormValue("_method")

	id, _ := strconv.Atoi(r.FormValue("id"))
	name := r.FormValue("name")
	status, _ := strconv.ParseBool(r.FormValue("status"))

	item := database.Item{
		Number:   id,
		ItemName: name,
		Status:   status,
	}

	mu.Lock()
	defer mu.Unlock()

	switch method {
	case "PUT":
		err = database.UpdateItem(db, id, &item)
	case "DELETE":
		err = database.DeleteItemById(db, id)
	default:
		err = database.InsertItem(db, &item)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// API to get all items (useful for JavaScript or other client-side rendering)
func getAllItemsAPI(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	items := database.GetAllItems(db)
	mu.Unlock()

	json.NewEncoder(w).Encode(items)
}

// Serve the static files (CSS, JS)
func serveStaticFiles(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/manage", manageItem)
	http.HandleFunc("/api/items", getAllItemsAPI)
	http.HandleFunc("/static/", serveStaticFiles)

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
