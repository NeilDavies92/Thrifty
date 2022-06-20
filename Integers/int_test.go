package integers

import (
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	result = 4

	if sum != result {
		t.Errorf("expected '%d' but got '%d'", expected, result)
	}
}
