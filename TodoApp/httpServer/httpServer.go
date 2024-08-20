package main

import (
	database "GoLang-Academy/TodoApp/Database"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Handler to list all items
func items(w http.ResponseWriter, req *http.Request) {
	// Set content type
	w.Header().Set("Content-Type", "application/json")

	switch req.Method {
	case "GET":
		// Retrieve all items
		items := database.GetAllItems(db)

		// Encode items to JSON
		if err := json.NewEncoder(w).Encode(items); err != nil {
			//todo note the errors aren't formatted in JSON
			http.Error(w, "Failed to encode items", http.StatusInternalServerError)
			return
		}
	case "POST":
		var newItem database.Item
		if err := json.NewDecoder(req.Body).Decode(&newItem); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		defer req.Body.Close()
		// Insert the item into the database
		if err := database.InsertItem(db, &newItem); err != nil {
			http.Error(w, fmt.Sprintf("Failed to insert item: %s", err.Error()), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(newItem); err != nil {
			http.Error(w, "Failed to encode response as JSON", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func item(w http.ResponseWriter, req *http.Request) {
	// Set content type
	w.Header().Set("Content-Type", "application/json")
	id, conversionErr := strconv.Atoi(req.PathValue("id"))
	if conversionErr != nil {
		http.Error(w, "Failed convert parameter", http.StatusInternalServerError)
	}

	switch req.Method {
	case "GET":
		// Retrieve all items
		items, err := database.GetItemByID(db, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Encode items to JSON
		if err := json.NewEncoder(w).Encode(items); err != nil {
			http.Error(w, "Failed to encode items", http.StatusInternalServerError)
			return
		}
	case "DELETE":
		err := database.DeleteItemById(db, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case "PUT":
		//todo think about case where you can update a ID but pass a different ID through the body. E.g validate ID in param matches ID in body?
		var newItem database.Item
		if err := json.NewDecoder(req.Body).Decode(&newItem); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		defer req.Body.Close()
		// Insert the item into the database
		if err := database.UpdateItem(db, id, &newItem); err != nil {
			http.Error(w, fmt.Sprintf("Failed to update item: %s", err.Error()), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusAccepted)
		if err := json.NewEncoder(w).Encode(newItem); err != nil {
			http.Error(w, "Failed to encode response as JSON", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Handler to echo request headers
func headers(w http.ResponseWriter, req *http.Request) {
	// Set content type
	w.Header().Set("Content-Type", "application/json")

	// Create a map to hold header information
	headerMap := make(map[string][]string)
	for name, headers := range req.Header {
		headerMap[name] = headers
	}

	// Encode headers to JSON
	if err := json.NewEncoder(w).Encode(headerMap); err != nil {
		http.Error(w, "Failed to encode headers", http.StatusInternalServerError)
	}
}

// Setup initial database with pre-populated data
func setupDatabase() *database.Database {
	testDB := database.CreateDatabase()

	// Pre-populate with one item
	initialItem := &database.Item{
		Number:   1,
		ItemName: "Pre-populated Item",
		Status:   false,
	}

	testDB.Items[initialItem.Number] = initialItem
	return testDB
}

// Initialize database
var db = setupDatabase()

func main() {

	mux := http.NewServeMux()

	// Register handlers
	mux.HandleFunc("/items", items)
	// Register the handler for /item/{id}
	mux.HandleFunc("/items/{id}", item)
	mux.HandleFunc("/headers", headers)

	// Start the server
	err := http.ListenAndServe(":8090", mux)
	if err != nil {
		return
	}
}
