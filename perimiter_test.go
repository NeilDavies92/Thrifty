package perimiter

import (
	"testing"
)

// Calculate the perimiter of a rectangle
func TestPerimiter(t *testing.T) {
	got := Perimiter(10.0, 10.0)
	want := 40.0

	if got != want {
		// %.2f = placeholder for 2 decimal place floating int
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// Calculate the perimiter of a rectangle
func TestArea(t *testing.T) {
	got := Area(12.0, 6.0)
	want := 72.0

	if got != want {
		// %.2f = placeholder for 2 decimal place floating int
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
