package main

import (
	"fmt"
)

const sayHello = "Hello, "
const italianHello = "Ciao, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "Italian" {
		return "Ciao, " + name
	}
	return sayHello + name
}

func main() {
	fmt.Println(Hello("Neil", ""))
}
