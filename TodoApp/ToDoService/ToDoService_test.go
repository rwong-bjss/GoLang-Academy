package toDoService

import (
	database "GoLang-Academy/TodoApp/Database"
	assert "GoLang-Academy/TodoApp/TestHelpers"
	"reflect"
	"testing"
)

func TestDeleteItem(t *testing.T) {
	type args struct {
		db *database.Database
		id int
	}

	// Initialize a mock or in-memory database
	db := database.CreateDatabase()

	// Add some mock data to the database for testing
	mockItem1 := &database.Item{Id: 1, ItemName: "Test Item 1", Status: true}
	mockItem2 := &database.Item{Id: 2, ItemName: "Test Item 2", Status: false}
	database.InsertItem(db, mockItem1)
	database.InsertItem(db, mockItem2)

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{{
		name:    "Delete item 1",
		args:    args{db: db, id: 1},
		wantErr: false,
	},
		{
			name:    "Delete item 1 again",
			args:    args{db: db, id: 1},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteItem(tt.args.db, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetItem(t *testing.T) {
	type args struct {
		db *database.Database
		id int
	}
	// Initialize a mock or in-memory database
	db := database.CreateDatabase()

	// Add some mock data to the database for testing
	mockItem1 := &database.Item{Id: 1, ItemName: "Test Item 1", Status: true}
	mockItem2 := &database.Item{Id: 2, ItemName: "Test Item 2", Status: false}
	database.InsertItem(db, mockItem1)
	database.InsertItem(db, mockItem2)

	expectItem1 := &Item{Number: 1, ItemName: "Test Item 1", Completed: true}
	expectItem2 := &Item{Number: 2, ItemName: "Test Item 2", Completed: false}
	tests := []struct {
		name    string
		args    args
		want    *Item
		wantErr bool
	}{
		{
			name:    "Item Exists",
			args:    args{db: db, id: 1},
			want:    expectItem1,
			wantErr: false,
		},
		{
			name:    "Item Does Not Exist",
			args:    args{db: db, id: 999},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Invalid ID",
			args:    args{db: db, id: -1},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Another Item Exists",
			args:    args{db: db, id: 2},
			want:    expectItem2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetItem(tt.args.db, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.EqualInterface(t, got, tt.want)
		})
	}
}

func TestGetItems(t *testing.T) {
	db := database.CreateDatabase()
	mockItem1 := &database.Item{Id: 1, ItemName: "Test Item 1", Status: true}
	mockItem2 := &database.Item{Id: 2, ItemName: "Test Item 2", Status: false}
	database.InsertItem(db, mockItem1)
	database.InsertItem(db, mockItem2)

	expectItem1 := &Item{Number: 1, ItemName: "Test Item 1", Completed: true}
	expectItem2 := &Item{Number: 2, ItemName: "Test Item 2", Completed: false}

	list := List{
		Items: []Item{*expectItem1, *expectItem2}}
	type args struct {
		db *database.Database
	}
	tests := []struct {
		name string
		args args
		want List
	}{
		{name: "Get Items",
			args: args{db: db},
			want: list},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetItems(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostItem(t *testing.T) {
	type args struct {
		db       *database.Database
		id       int
		itemName string
		status   bool
	}

	db := database.CreateDatabase()

	// Add some mock data to the database for testing
	expectedMockItem1 := Item{Number: 1, ItemName: "Test Item 1", Completed: true}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Post item 1",
			args:    args{db: db, id: 1, itemName: "Test Item 1", status: true},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PostItem(tt.args.db, tt.args.id, tt.args.itemName, tt.args.status); (err != nil) != tt.wantErr {
				t.Errorf("PostItem() error = %v, wantErr %v", err, tt.wantErr)
			}
			get, _ := GetItem(db, 1)
			assert.EqualInterface(t, *get, expectedMockItem1)
		})
	}
}

func TestUpdateItem(t *testing.T) {
	db := database.CreateDatabase()

	// Add some mock data to the database for testing
	mockItem1 := &database.Item{Id: 1, ItemName: "Test Item 1", Status: true}
	mockItem2 := &database.Item{Id: 2, ItemName: "Test Item 2", Status: false}
	database.InsertItem(db, mockItem1)
	database.InsertItem(db, mockItem2)

	expectItem1 := &Item{Number: 1, ItemName: "Test Item 1 Updated", Completed: false}

	type args struct {
		db       *database.Database
		id       int
		itemName string
		status   bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Update item 1",
			args:    args{db: db, id: 1, itemName: "Test Item 1 Updated", status: false},
			wantErr: false,
		},
		{
			name:    "Update item doesn't exist",
			args:    args{db: db, id: 3, itemName: "Test Item 1 Updated", status: false},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateItem(tt.args.db, tt.args.id, tt.args.itemName, tt.args.status); (err != nil) != tt.wantErr {
				t.Errorf("UpdateItem() error = %v, wantErr %v", err, tt.wantErr)
			}

			get, _ := GetItem(db, 1)
			assert.EqualInterface(t, *get, *expectItem1)
		})
	}
}

func Test_errorCheck(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := errorCheck(tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("errorCheck() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
