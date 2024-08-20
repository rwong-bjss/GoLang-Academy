package main

import "testing"

func TestAssignments(t *testing.T) {
	t.Run("can concatenate a full name", func(t *testing.T) {
		got := concatenateStrings("rachel", "sophie", "wong")
		want := "rachel sophie wong"

		assertCorrectMessage(t, got, want)
	})

	t.Run("can concatenate 4 strings", func(t *testing.T) {
		got := concatenateStrings("one", "two", "three", "four")
		want := "one two three four"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// needed to tell the test suite that this method is a helper.
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
