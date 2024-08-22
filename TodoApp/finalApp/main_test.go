package main

import (
	database "GoLang-Academy/TodoApp/Database"
)

func setupTestDatabase() *database.Database {
	testDB := database.CreateDatabase()

	// Pre-populate with one item
	initialItem := &database.Item{
		Number:   1,
		ItemName: "Test item",
		Status:   true,
	}

	testDB.Items[initialItem.Number] = initialItem
	return testDB
}

// Similiar tests / examples: TodoApp/httpServer/httpServer_test.go
