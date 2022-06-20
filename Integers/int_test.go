package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	result := 4

	if sum != result {
		t.Errorf("expected '%d' but got '%d'", result, sum)
	}
}

// Go examples are executed just like tests so you
// can be confident xamples reflect what the
// code actually does.

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
