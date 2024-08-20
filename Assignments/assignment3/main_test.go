package main

import (
	"fmt"
	"testing"
)

const formatMessage = "%d is between 1 and 10"

func TestAssignments(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{fmt.Sprintf(formatMessage, 5), 5, true},
		{fmt.Sprintf(formatMessage, 0), 0, false},
		{fmt.Sprintf(formatMessage, 1), 1, true},
		{fmt.Sprintf(formatMessage, 10), 10, true},
		{fmt.Sprintf(formatMessage, 11), 11, false},
		{fmt.Sprintf(formatMessage, -1), -1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := isNumberBetween1And10(tt.input)
			assertCorrectBoolean(t, got, tt.want)
		})
	}
}

func assertCorrectBoolean(t testing.TB, got, want bool) {
	// needed to tell the test suite that this method is a helper.
	t.Helper()
	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}
