package iteration

import "testing"

func TestRepeat(t *testing.T) {
	returnCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()

		repeated := Repeat("a", 0)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	t.Run("Loop x amount of times", func(t *testing.T) {
		got := Repeat("n", 6)
		want := "nnnnnn"

		if got != want {
			returnCorrectMessage(t, got, want)
		}
	})
}
