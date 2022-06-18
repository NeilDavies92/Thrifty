package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Neil")
	want := "Hello, Neil"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
