package main

import (
	database "GoLang-Academy/TodoApp/Database"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Handler to list all Items
func Items(w http.ResponseWriter, req *http.Request) {
	// Set content type
	w.Header().Set("Content-Type", "application/json")

	switch req.Method {
	case "GET":
		// Retrieve all Items
		items := database.GetAllItems(db)

		// Encode Items to JSON
		if err := json.NewEncoder(w).Encode(items); err != nil {
			//todo note the errors aren't formatted in JSON
			http.Error(w, "Failed to encode Items", http.StatusInternalServerError)
			return
		}
	case "POST":
		var newItem database.Item
		if err := json.NewDecoder(req.Body).Decode(&newItem); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
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

func Item(w http.ResponseWriter, req *http.Request) {
	// Set content type
	w.Header().Set("Content-Type", "application/json")
	id, conversionErr := strconv.Atoi(req.PathValue("id"))
	if conversionErr != nil {
		http.Error(w, "Failed convert parameter", http.StatusInternalServerError)
	}

	switch req.Method {
	case "GET":
		// Retrieve all Items
		items, err := database.GetItemByID(db, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Encode Items to JSON
		if err := json.NewEncoder(w).Encode(items); err != nil {
			http.Error(w, "Failed to encode Items", http.StatusInternalServerError)
			return
		}
	case "DELETE":
		err := database.DeleteItemById(db, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
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
//use interface to mock in test
//or use real db for the unit test

// in this case it may have been easier to use the main db in the unit tests and not bother with overriding a global variable and have main set the db.
// however I wanted an example of how we can override the global variable in the test
var db = setupDatabase()

func newAppMux() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/items", Items)
	router.HandleFunc("/items/{id}", Item)
	router.HandleFunc("/headers", headers)
	return router
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	mux := newAppMux()
	// Start the server
	err := http.ListenAndServe(":8090", mux)
	if err != nil {
		return
	}
}
