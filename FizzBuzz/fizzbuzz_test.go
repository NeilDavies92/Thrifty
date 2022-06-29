package main

import (
	"testing"
)

func TestFizz(t *testing.T) {
	got := "fizz"
	want := "fizz"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
