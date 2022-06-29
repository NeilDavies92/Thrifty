package main

func Sum(numbers [5]int) int {

	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}

// In a slice you omit the length
func Slice(numbers []int) int {

	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}
