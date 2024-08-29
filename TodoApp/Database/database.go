package database

import (
	"fmt"
)

type Item struct {
	Id       int    `json:"id"`
	ItemName string `json:"item_name"`
	Status   bool   `json:"status"`
}

type List struct {
	Items []Item `json:"items"`
}

type Database struct {
	Items map[int]*Item
}

// The function CreateDatabase is defined to return a pointer to a Database struct (*Database).
// This means that when you call CreateDatabase,
// it returns the memory address of a Database instance rather than a copy of the instance itself.
// Returning a pointer is efficient, especially for large data structures,
// because it avoids copying the entire struct.
func CreateDatabase() *Database {
	//In the return statement, &Database{} is used to create a new instance of the Database struct and return its memory address.
	//The & operator here takes the newly created Database struct and returns a pointer to it.
	return &Database{
		Items: make(map[int]*Item),
	}
}

// InsertItem Insert an item into the database
func InsertItem(db *Database, item *Item) error {
	if _, exists := db.Items[item.Id]; exists {
		return fmt.Errorf("item with ID %d already exists", item.Id)
	}
	db.Items[item.Id] = item
	return nil
}

// GetItemByID Function to get a single item by its ID
func GetItemByID(db *Database, id int) (*Item, error) {
	item, exists := db.Items[id]
	if !exists {
		return nil, fmt.Errorf("item with ID %d not found", id)
	}
	return item, nil
}

func GetAllItems(db *Database) List {
	var list List
	for _, item := range db.Items {
		list.Items = append(list.Items, *item) // Dereference the pointer to get the actual Item value
	}
	return list
}

// DeleteItemById Function to delete an item by its ID
func DeleteItemById(db *Database, id int) error {
	if _, exists := db.Items[id]; !exists {
		return fmt.Errorf("item with ID %d not found", id)
	}
	delete(db.Items, id)
	return nil
}

// UpdateItem Update an item in the database
func UpdateItem(db *Database, id int, item *Item) error {
	if _, exists := db.Items[id]; !exists {
		return fmt.Errorf("item with ID %d not found", id)
	}
	db.Items[id] = item
	return nil
}
