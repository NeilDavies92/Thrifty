package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	returnCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Say hello to a specific person", func(t *testing.T) {
		got := Hello("Neil")
		want := "Hello, Neil"

		if got != want {
			returnCorrectMessage(t, got, want)
		}
	})

	t.Run("Say hello world if empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			returnCorrectMessage(t, got, want)
		}
	})

	t.Run("In italian", func(t *testing.T) {
		got := Hello("Neil!")
		want := "Ciao mondo Neil!"

		if got != want {
			returnCorrectMessage(t, got, want)
		}
	})
}
