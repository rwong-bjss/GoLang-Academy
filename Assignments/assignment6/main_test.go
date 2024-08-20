package main

import "testing"

func TestAssignments(t *testing.T) {
	t.Run("Can calculate 1997 date of birth", func(t *testing.T) {
		got := calculateAgeFromToday(1997, 03, 24)
		want := 27

		assertCorrectMessage(t, got, want)
	})

	t.Run("Can calculate 2000 date of birth", func(t *testing.T) {
		got := calculateAgeFromToday(2000, 02, 24)
		want := 24

		assertCorrectMessage(t, got, want)
	})

	t.Run("Can calculate leap year date of birth", func(t *testing.T) {
		got := calculateAgeFromToday(2020, 02, 29)
		want := 4

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want int) {
	// needed to tell the test suite that this method is a helper.
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
