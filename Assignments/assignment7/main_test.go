package main

import (
	"fmt"
	"testing"
)

func TestAssignments(t *testing.T) {
	t.Run("can generate random dice roll number", func(t *testing.T) {
		got1, got2 := roll2Dice()
		want := validateDiceNumbers(got1, got2)

		assertCorrectMessage(t, want, true)
	})
	//t.Run("can concatenate 4 strings", func(t *testing.T) {
	//	got := concatenateStrings("one", "two", "three", "four")
	//	want := "one two three four"
	//
	//	assertCorrectMessage(t, got, want)
	//})
}

const formatMessage = "%d has returned the expected value of %s"

func TestSwitch(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{fmt.Sprintf(formatMessage, 1, "NEUTRAL"), 1, "NEUTRAL"},
		{fmt.Sprintf(formatMessage, 2, "SNAKE-EYES-CRAPS"), 2, "SNAKE-EYES-CRAPS"},
		{fmt.Sprintf(formatMessage, 3, "LOSS-CRAPS"), 3, "LOSS-CRAPS"},
		{fmt.Sprintf(formatMessage, 7, "NATURAL"), 7, "NATURAL"},
		{fmt.Sprintf(formatMessage, 11, "NATURAL"), 11, "NATURAL"},
		{fmt.Sprintf(formatMessage, 10, "NEUTRAL"), 10, "NEUTRAL"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := switchCase(tt.input, 0)
			assertCorrectString(t, got, tt.expected)
		})
	}
}

func assertCorrectString(t testing.TB, got, want string) {
	// needed to tell the test suite that this method is a helper.
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertCorrectMessage(t testing.TB, got, want bool) {
	// needed to tell the test suite that this method is a helper.
	t.Helper()
	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func validateDiceNumbers(first, second int) bool {
	if first > 0 && first < 7 && second > 0 && second < 7 {
		return true
	} else {
		return false
	}
}
