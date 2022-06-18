package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("Say hello to a specific person", func(t *testing.T) {
		got := Hello("Neil")
		want := "Hello, Neil"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Say hello world if empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}
