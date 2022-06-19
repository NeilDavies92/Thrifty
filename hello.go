package main

import (
	"fmt"
)

const italian = "Italian"
const englishHello = "Hello, "
const italianHello = "Ciao, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == italian {
		return italianHello + name
	}
	return englishHello + name
}

func main() {
	fmt.Println(Hello("Neil", ""))
}
