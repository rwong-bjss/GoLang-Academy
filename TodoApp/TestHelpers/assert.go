package assert

import (
	"reflect"
	"testing"
)

func Equal[T comparable](t *testing.T, actual, expected T) {
	t.Helper()

	if actual != expected {
		t.Errorf("got: %v; want: %v", actual, expected)
	}
}

func EqualSlice(t *testing.T, actual, expected interface{}) {
	t.Helper()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got: %v; want: %v", actual, expected)
	}
}

func AssertSize(t *testing.T, got interface{}, want int) {
	t.Helper()

	gotValue := reflect.ValueOf(got)
	if gotValue.Kind() != reflect.Slice {
		t.Errorf("expected a slice type, but got %T", got)
		return
	}
	if gotValue.Len() != want {
		t.Errorf("expected size %d, but got %d", want, gotValue.Len())
	}
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// needed to tell the test suite that this method is a helper.
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
