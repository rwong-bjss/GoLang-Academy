package main

import (
	database "GoLang-Academy/TodoApp/Database"
	assert "GoLang-Academy/TodoApp/TestHelpers"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func setupTestDatabase() *database.Database {
	testDB := database.CreateDatabase()

	// Pre-populate with one item
	initialItem := &database.Item{
		Id:       1,
		ItemName: "Test item",
		Status:   true,
	}

	testDB.Items[initialItem.Id] = initialItem
	return testDB
}

// Similiar tests / examples: TodoApp/httpServer/httpServer_test.go
func TestIntegrationPost(t *testing.T) {
	db = setupTestDatabase()              // Ensure db is set to the test database
	ts := httptest.NewServer(newAppMux()) // Set up a new test server
	defer ts.Close()

	client := ts.Client()

	formData := url.Values{}
	formData.Set("id", "2")
	formData.Set("name", "Test item 2")
	formData.Set("status", "true")
	formData.Set("_method", "POST")

	res, err := client.PostForm(fmt.Sprintf("%v/items", ts.URL), formData)
	assert.AssertNoError(t, err)

	defer res.Body.Close()

	assert.Equal(t, res.StatusCode, http.StatusOK)

	// Optionally, follow the redirect to ensure it lands on the correct page
	location, err := res.Location()
	assert.AssertNoError(t, err)
	assert.Equal(t, location.String(), "/") // The location should be the root URL

	items := database.GetAllItems(db)
	assert.AssertNoError(t, err)

	want := database.Item{
		Id:       2,
		ItemName: "Test item 2",
		Status:   true,
	}

	found := false
	for _, got := range items.Items {
		if got.Id == want.Id && got.ItemName == want.ItemName && got.Status == want.Status {
			found = true
			break
		}
	}

	assert.Equal(t, found, true)
}
