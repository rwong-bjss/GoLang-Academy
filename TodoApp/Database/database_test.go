package database

import (
	assert "GoLang-Academy/TodoApp/TestHelpers"
	"fmt"
	"testing"
)

// Global variable to hold the database

func setupDatabase() *Database {
	testDB := CreateDatabase()

	// Pre-populate with one item
	initialItem := &Item{
		Number:   1,
		ItemName: "Pre-populated Item",
		Status:   false,
	}

	testDB.Items[initialItem.Number] = initialItem
	return testDB
}

func TestDatabase(t *testing.T) {
	t.Run("Can create a db and add an item", func(t *testing.T) {
		database := setupDatabase()
		err := InsertItem(database, &Item{2, "Laundry", true})
		assert.AssertNoError(t, err)
	})

	t.Run("Can create a db and error on item exists", func(t *testing.T) {
		database := setupDatabase()
		got := InsertItem(database, &Item{1, "Laundry", true})
		want := fmt.Errorf("item with ID %d already exists", 1)
		assert.EqualInterface(t, got, want)
	})

	t.Run("Can create a db and get an item", func(t *testing.T) {
		database := setupDatabase()
		item := &Item{2, "Laundry", true}
		err := InsertItem(database, item)
		assert.AssertNoError(t, err)

		actualItem, err := GetItemByID(database, 2)
		assert.Equal(t, actualItem, item)
	})

	t.Run("Can throw error if item doesn't exist", func(t *testing.T) {
		database := setupDatabase()
		_, got := GetItemByID(database, 2)
		want := fmt.Errorf("item with ID %d not found", 2)
		assert.EqualInterface(t, got, want)
	})

	t.Run("Can get list of items", func(t *testing.T) {
		database := setupDatabase()
		err := InsertItem(database, &Item{2, "Laundry", true})
		assert.AssertNoError(t, err)

		items, _ := GetAllItems(database)
		assert.AssertSize(t, items.Items, 2)

		expectedList := []Item{
			{Number: 1, ItemName: "Pre-populated Item", Status: false},
			{Number: 2, ItemName: "Laundry", Status: true},
		}

		assert.EqualInterface(t, items.Items, expectedList)
	})

	t.Run("Can delete item by id", func(t *testing.T) {
		database := setupDatabase()
		err := InsertItem(database, &Item{2, "Laundry", true})
		assert.AssertNoError(t, err)

		items, _ := GetAllItems(database)
		assert.AssertSize(t, items.Items, 2)

		errDel := DeleteItemById(database, 2)
		assert.AssertNoError(t, errDel)

		itemsActual, _ := GetAllItems(database)
		assert.AssertSize(t, itemsActual.Items, 1)
	})

	t.Run("Can fail to delete item by id if it doesn't exist", func(t *testing.T) {
		database := setupDatabase()

		got := DeleteItemById(database, 2)
		want := fmt.Errorf("item with ID %d not found", 2)
		assert.EqualInterface(t, got, want)

	})

	t.Run("Can update item", func(t *testing.T) {
		database := setupDatabase()
		updateItem := &Item{
			Number:   1,
			ItemName: "newer item",
			Status:   true,
		}
		err := UpdateItem(database, 1, updateItem)
		assert.AssertNoError(t, err)

		itemActual, _ := GetItemByID(database, 1)
		assert.EqualInterface(t, itemActual, updateItem)
	})

	t.Run("Can fail update item that does not exist", func(t *testing.T) {
		database := setupDatabase()
		updateItem := &Item{
			Number:   1,
			ItemName: "newer item",
			Status:   true,
		}
		got := UpdateItem(database, 5, updateItem)
		want := fmt.Errorf("item with ID %d not found", 5)
		assert.EqualInterface(t, got, want)
	})
}
