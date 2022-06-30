package main

import "strconv"

func FizzBuzz(i int) string {
	// Divisible by 3
	if i%3 == 0 {
		return "fizz"
	}

	// Divisible by 5
	if i%5 == 0 {
		return "buzz"
	}

	return strconv.Itoa(i)
}

func main() {
	for i := 1; i <= 100; i++ {
		FizzBuzz(i)
	}
}
