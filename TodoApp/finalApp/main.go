package main

import (
	database "GoLang-Academy/TodoApp/Database"
	toDoService "GoLang-Academy/TodoApp/ToDoService"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var db = database.CreateDatabase()

func homePage(w http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("TodoApp/finalApp/index.html"))
	items := toDoService.GetItems(db)
	tmpl.Execute(w, items)
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	method, id, itemName, status := extractFormData(w, r)
	switch method {
	case http.MethodGet:
		items := toDoService.GetItems(db)
		if err := json.NewEncoder(w).Encode(items); err != nil {
			//todo note the errors aren't formatted in JSON
			http.Error(w, "Failed to encode Items", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		err := toDoService.PostItem(db, id, itemName, status)
		if err != nil {
			http.Error(w, "Failed to create item: "+err.Error(), http.StatusBadRequest)
			return
		}
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func itemHandler(w http.ResponseWriter, r *http.Request) {
	method, id, itemName, status := extractFormDataWithPathID(w, r)

	switch method {
	case http.MethodGet:
		item, err := toDoService.GetItem(db, id)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to retrieve item: %s", err.Error()), http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(item); err != nil {
			http.Error(w, "Failed to encode item as JSON", http.StatusInternalServerError)
			return
		}
	case http.MethodPut:
		err := toDoService.UpdateItem(db, id, itemName, status)
		if err != nil {
			return
		}
	case http.MethodDelete:
		err := toDoService.DeleteItem(db, id)
		if err != nil {
			return
		}
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func extractFormData(w http.ResponseWriter, req *http.Request) (method string, id int, itemName string, status bool) {
	if req.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	err := req.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	itemId, _ := strconv.Atoi(req.FormValue("id"))
	method = req.FormValue("_method")
	name := req.FormValue("name")
	status = req.FormValue("status") == "true"
	return method, itemId, name, status
}

func extractFormDataWithPathID(w http.ResponseWriter, req *http.Request) (method string, id int, itemName string, status bool) {
	pathId, conversionErr := strconv.Atoi(req.PathValue("id"))
	if conversionErr != nil {
		http.Error(w, "Failed to convert parameter: "+conversionErr.Error(), http.StatusInternalServerError)
	}
	m, _, n, s := extractFormData(w, req)
	return m, pathId, n, s
}

func newAppMux() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", homePage)
	router.HandleFunc("/itemsHandler", itemsHandler)
	router.HandleFunc("/itemsHandler/{id}", itemHandler)

	return router
}

func main() {
	mux := newAppMux()

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
