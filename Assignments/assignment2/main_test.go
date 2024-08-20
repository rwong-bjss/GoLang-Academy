package main

import "testing"

func TestAssignments(t *testing.T) {
	t.Run("can concatonate a full name", func(t *testing.T) {
		got := concatenateFullName("rachel", "sophie", "wong")
		want := "rachel sophie wong"

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
