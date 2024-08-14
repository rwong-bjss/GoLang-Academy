package main

import (
	assert "GoLang-Academy/TodoApp/TestHelpers"
	"testing"
)

func TestAssignments(t *testing.T) {
	t.Run("can return correct size list of todo with one item", func(t *testing.T) {
		got := createList(Item{1, "todo", true})
		want := 1

		assert.AssertSize(t, got, want)
	})
	t.Run("can return correct size list of todo with multiple items", func(t *testing.T) {
		got := createList(Item{1, "todo", true}, Item{2, "todo", false})
		want := 2

		assert.AssertSize(t, got, want)
	})

	t.Run("can return list of todo with multiple items", func(t *testing.T) {
		got := createList(Item{1, "todo", true}, Item{2, "todo", false})
		want := []Item{{1, "todo", true}, {2, "todo", false}}

		assert.EqualSlice(t, got, want)
	})
}
