package main

import (
	_ "GoLang-Academy/TodoApp/Database"
	database "GoLang-Academy/TodoApp/Database"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Create the database
	db := database.CreateDatabase()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nChoose an operation: create, read, update, delete, list, or exit")
		fmt.Print("-> ")
		action, _ := reader.ReadString('\n')
		action = strings.TrimSpace(action)

		switch action {
		case "create":
			fmt.Print("Enter ID: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))

			fmt.Print("Enter item name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Enter status (true/false): ")
			statusStr, _ := reader.ReadString('\n')
			status, _ := strconv.ParseBool(strings.TrimSpace(statusStr))

			item := database.Item{Id: id, ItemName: name, Status: status}
			err := database.InsertItem(db, &item)
			if err != nil {
				fmt.Println("Failed to create item:", err)
			} else {
				fmt.Printf("Created item: %+v\n", item)
			}

		case "read":
			fmt.Print("Enter ID: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))

			item, err := database.GetItemByID(db, id)
			if err != nil {
				fmt.Println("Failed to read item:", err)
			} else {
				fmt.Printf("Read item: %+v\n", item)
			}

		case "update":
			fmt.Print("Enter ID: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))

			fmt.Print("Enter new item name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Enter new status (true/false): ")
			statusStr, _ := reader.ReadString('\n')
			status, _ := strconv.ParseBool(strings.TrimSpace(statusStr))

			newItem := database.Item{Id: id, ItemName: name, Status: status}
			err := database.UpdateItem(db, id, &newItem)
			if err != nil {
				fmt.Println("Failed to update item:", err)
			} else {
				fmt.Printf("Updated item: %+v\n", newItem)
			}

		case "delete":
			fmt.Print("Enter ID: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))

			err := database.DeleteItemById(db, id)
			if err != nil {
				fmt.Println("Failed to delete item:", err)
			} else {
				fmt.Printf("Deleted item with ID: %d\n", id)
			}

		case "list":
			items := database.GetAllItems(db)

			for _, item := range items.Items {
				fmt.Printf("ID: %d, Item: %s, Status: %t\n", item.Id, item.ItemName, item.Status)
			}

		case "exit":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Unknown action. Please choose: create, read, update, delete, list, or exit")
		}
	}
}
