package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {

		// Divisible by 3
		if i%3 == 0 {
			{
				fmt.Printf("fizz")
			}
		}

		// Divisible by 5
		if i%5 == 0 {
			fmt.Printf("buzz")
		}

		// Neither
		if i%3 != 0 && i%5 != 0 {
			fmt.Printf("%d", i)
		}

		fmt.Printf("\n")
	}
}
