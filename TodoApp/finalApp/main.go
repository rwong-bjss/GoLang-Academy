package main

import (
	database "GoLang-Academy/TodoApp/Database"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Global database instance
var db = database.CreateDatabase()

// Handler for the homepage
func homePage(w http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	items := database.GetAllItems(db)
	tmpl.Execute(w, items)
}

func Items(w http.ResponseWriter, req *http.Request) {

	if req.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	err := req.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	method := req.FormValue("_method")
	id, _ := strconv.Atoi(req.FormValue("id"))
	name := req.FormValue("name")
	status := req.FormValue("status") == "true"

	item := database.Item{
		Number:   id,
		ItemName: name,
		Status:   status,
	}

	switch method {
	case "POST":
		if err := database.InsertItem(db, &item); err != nil {
			http.Error(w, fmt.Sprintf("Failed to insert item: %s", err.Error()), http.StatusInternalServerError)
			return
		}
	case "GET":
		items := database.GetAllItems(db)
		if err := json.NewEncoder(w).Encode(items); err != nil {
			//todo note the errors aren't formatted in JSON
			http.Error(w, "Failed to encode Items", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func Item(w http.ResponseWriter, req *http.Request) {
	id, conversionErr := strconv.Atoi(req.PathValue("id"))
	if conversionErr != nil {
		http.Error(w, "Failed convert parameter", http.StatusInternalServerError)
	}

	if req.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	err := req.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	method := req.FormValue("_method")
	name := req.FormValue("name")
	status := req.FormValue("status") == "true"

	item := database.Item{
		Number:   id,
		ItemName: name,
		Status:   status,
	}

	switch method {
	case "PUT":
		if err := database.UpdateItem(db, id, &item); err != nil {
			http.Error(w, fmt.Sprintf("Failed to update item: %s", err.Error()), http.StatusInternalServerError)
			return
		}
	case "DELETE":
		if err := database.DeleteItemById(db, id); err != nil {
			http.Error(w, fmt.Sprintf("Failed to delete item: %s", err.Error()), http.StatusInternalServerError)
			return
		}
	case "GET":
		item, err := database.GetItemByID(db, id)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to retrieve item: %s", err.Error()), http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(item); err != nil {
			http.Error(w, "Failed to encode item as JSON", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func newAppMux() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", homePage)
	router.HandleFunc("/items", Items)
	router.HandleFunc("/items/{id}", Item)
	return router
}

func main() {
	mux := newAppMux()

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
