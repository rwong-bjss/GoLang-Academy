package toDoService

import (
	database "GoLang-Academy/TodoApp/Database"
)

// Service models
type Item struct {
	Number    int    `json:"number"`
	ItemName  string `json:"item_name"`
	Completed bool   `json:"completed"`
}

type List struct {
	Items []Item `json:"items"`
}

func GetItems(db *database.Database) List {
	items := database.GetAllItems(db)

	var list List
	for _, item := range items.Items {
		list.Items = append(list.Items, Item{Number: item.Id, ItemName: item.ItemName, Completed: item.Status})
	}
	return list
}

func GetItem(db *database.Database, id int) (*Item, error) {
	items, err := database.GetItemByID(db, id)
	if err != nil {
		return nil, err
	}
	itemToReturn := &Item{Number: items.Id, ItemName: items.ItemName, Completed: items.Status}

	return itemToReturn, nil
}

func PostItem(db *database.Database, id int, itemName string, status bool) error {
	ItemToSave := database.Item{
		Id:       id,
		ItemName: itemName,
		Status:   status,
	}
	err := database.InsertItem(db, &ItemToSave)
	return errorCheck(err)
}

func UpdateItem(db *database.Database, id int, itemName string, status bool) error {
	ItemToSave := database.Item{
		Id:       id,
		ItemName: itemName,
		Status:   status,
	}
	err := database.UpdateItem(db, id, &ItemToSave)

	return errorCheck(err)
}

func DeleteItem(db *database.Database, id int) error {
	err := database.DeleteItemById(db, id)
	return errorCheck(err)
}

func errorCheck(err error) error {
	if err != nil {
		return err
	}
	return nil
}
