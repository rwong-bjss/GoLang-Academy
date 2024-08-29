package main

import (
	database "GoLang-Academy/TodoApp/Database"
	assert "GoLang-Academy/TodoApp/TestHelpers"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestGETItems(t *testing.T) {
	db = setupTestDatabase()

	t.Run("returns pre populated Items", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/itemsHandler", nil)
		response := httptest.NewRecorder()

		Items(response, request)
		var got database.List
		err := json.Unmarshal(response.Body.Bytes(), &got)
		assert.AssertNoError(t, err)

		// Create a list containing the item
		want := database.List{
			Items: []database.Item{{
				Id:       1,
				ItemName: "Test item",
				Status:   true,
			}},
		}
		assert.Equal(t, response.Code, http.StatusOK)
		assert.EqualInterface(t, got, want)
	})
}

func TestPOSTItems(t *testing.T) {
	t.Run("returns created item", func(t *testing.T) {
		item := database.Item{
			Id:       1,
			ItemName: "Test item",
			Status:   true,
		}
		reader, _ := jsonReaderFactory(item)
		request, _ := http.NewRequest(http.MethodPost, "/itemsHandler", reader)
		response := httptest.NewRecorder()
		Items(response, request)
		var got database.Item
		err := json.Unmarshal(response.Body.Bytes(), &got)
		if err != nil {
			return
		}
		want := item
		assert.Equal(t, response.Code, http.StatusCreated)
		assert.EqualInterface(t, got, want)
	})
}

func jsonReaderFactory(in interface{}) (io.Reader, error) {
	buf := bytes.NewBuffer(nil)
	enc := json.NewEncoder(buf)
	err := enc.Encode(in)
	if err != nil {
		return nil, fmt.Errorf("creating reader: error encoding data: %s", err)
	}
	return buf, nil
}

func TestPUTItems(t *testing.T) {
	t.Run("returns updated item", func(t *testing.T) {
		item := database.Item{
			Id:       1,
			ItemName: "changed item",
			Status:   false,
		}

		reader, _ := jsonReaderFactory(item)
		request, _ := http.NewRequest(http.MethodPut, "/itemsHandler/1", reader)
		request.SetPathValue("id", "1")
		response := httptest.NewRecorder()
		Item(response, request)
		var got database.Item
		err := json.Unmarshal(response.Body.Bytes(), &got)
		if err != nil {
			return
		}
		want := item
		assert.Equal(t, response.Code, http.StatusAccepted)
		assert.EqualInterface(t, got, want)
	})
}

func TestDELETEItems(t *testing.T) {
	t.Run("returns deleted item", func(t *testing.T) {
		//todo this is a bit hacky to test the path params
		request, _ := http.NewRequest(http.MethodDelete, "/itemsHandler/1", nil)
		request.SetPathValue("id", "1")
		response := httptest.NewRecorder()
		Item(response, request)
		assert.Equal(t, response.Code, http.StatusOK)
	})
}

func TestGETItem(t *testing.T) {
	t.Run("returns item", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/itemsHandler/1", nil)
		response := httptest.NewRecorder()
		Item(response, request)
		var got database.Item
		err := json.Unmarshal(response.Body.Bytes(), &got)
		if err != nil {
			return
		}
		want := database.Item{
			Id:       1,
			ItemName: "Test item",
			Status:   true,
		}
		assert.Equal(t, response.Code, http.StatusOK)
		assert.EqualInterface(t, got, want)
	})
}

func TestIntegrationGet(t *testing.T) {
	db = setupTestDatabase() // Ensure db is set to the test database
	ts := httptest.NewServer(newAppMux())
	defer ts.Close()

	client := ts.Client()

	//res, err := client.Get(fmt.Sprintf("%v/items", ts.URL), "application/json", strings.NewReader(`{
	//    "number": 0,
	//    "item_name": "",
	//    "completed": false
	//}`))
	res, err := client.Get(fmt.Sprintf("%v/items", ts.URL))
	if err != nil {
		t.Errorf("Wasn't expecting error. Got: %v", err)
	}

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Errorf("Wasn't expecting error. Got: %v", err)
	}

	var got database.List
	err = json.Unmarshal(resBody, &got)
	if err != nil {
		return
	}

	// Create a list containing the item
	want := database.List{
		Items: []database.Item{{
			Id:       1,
			ItemName: "Test item",
			Status:   true,
		}},
	}
	assert.Equal(t, res.StatusCode, http.StatusOK)
	assert.EqualInterface(t, got, want)
}

func TestIntegrationPost(t *testing.T) {
	db = setupTestDatabase() // Ensure db is set to the test database
	ts := httptest.NewServer(newAppMux())
	defer ts.Close()

	client := ts.Client()

	res, err := client.Post(fmt.Sprintf("%v/items", ts.URL), "application/json", strings.NewReader(`{
	   "number": 2,
	   "item_name": "Test item 2",
	   "completed": true
	}`))
	assert.AssertNoError(t, err)

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	assert.AssertNoError(t, err)

	var got database.Item
	err = json.Unmarshal(resBody, &got)
	assert.AssertNoError(t, err)

	want := database.Item{
		Id:       2,
		ItemName: "Test item 2",
		Status:   true,
	}
	assert.Equal(t, res.StatusCode, http.StatusCreated)
	assert.EqualInterface(t, got, want)
}
