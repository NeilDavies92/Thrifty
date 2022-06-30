package main

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {

	t.Run("Test Fizz", func(t *testing.T) {
		got := FizzBuzz(3)
		want := "fizz"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Test Buzz", func(t *testing.T) {
		got := FizzBuzz(5)
		want := "buzz"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
