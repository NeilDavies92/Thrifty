package perimiter

import (
	"testing"
)

// Calculate the perimiter of a rectangle
func TestPerimiter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimiter(rectangle)
	want := 40.0

	if got != want {
		// %.2f = placeholder for 2 decimal place floating int
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// Calculate the perimiter of a rectangle
func TestArea(t *testing.T) {

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		got := Area(rectangle)
		want := 72.0

		if got != want {
			// %.2f = placeholder for 2 decimal place floating int
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{12.0, 6.0}
		got := Area(circle)
		want := 72.0

		if got != want {
			// %.2f = placeholder for 2 decimal place floating int
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

}
