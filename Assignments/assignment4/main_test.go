package main

import (
	"reflect"
	"testing"
)

func TestAssignments(t *testing.T) {
	t.Run("can return numbers in array", func(t *testing.T) {
		got := returnArrayAscending()
		want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		assertCorrectMessage(t, got, want)
	})

	t.Run("can return numbers in array", func(t *testing.T) {
		got := returnArrayDescending()
		want := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

		assertCorrectMessage(t, got, want)
	})

	t.Run("can return even numbers in ascending array", func(t *testing.T) {
		got, _, _, _ := returnSeriesOfNumbers(returnArrayAscending(), returnArrayAscending(), returnArrayDescending(), returnArrayDescending())
		want := []int{2, 4, 6, 8, 10}
		assertCorrectMessage(t, got, want)
	})

	t.Run("can return odd numbers in ascending array", func(t *testing.T) {
		_, got, _, _ := returnSeriesOfNumbers(returnArrayAscending(), returnArrayAscending(), returnArrayDescending(), returnArrayDescending())
		want := []int{1, 3, 5, 7, 9}
		assertCorrectMessage(t, got, want)
	})

	t.Run("can return even numbers in descending array", func(t *testing.T) {
		_, _, got, _ := returnSeriesOfNumbers(returnArrayAscending(), returnArrayAscending(), returnArrayDescending(), returnArrayDescending())
		want := []int{10, 8, 6, 4, 2}
		assertCorrectMessage(t, got, want)
	})

	t.Run("can return odd numbers in descending array", func(t *testing.T) {
		_, _, _, got := returnSeriesOfNumbers(returnArrayAscending(), returnArrayAscending(), returnArrayDescending(), returnArrayDescending())
		want := []int{9, 7, 5, 3, 1}
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want []int) {
	// needed to tell the test suite that this method is a helper.
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d", got, want)
	}
}
