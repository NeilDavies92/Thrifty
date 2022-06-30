package main

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	got := FizzBuzz(4)
	want := "fizz"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
