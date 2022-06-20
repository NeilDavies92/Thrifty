package main

import (
	"fmt"
)

const italian = "Italian"
const french = "French"
const bonjour = "Bonjour, "
const hello = "Hello, "
const ciao = "Ciao, "

// Public functions start with a capital letter
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return Greeting(language) + name
}

// Private functions start with a lower case letter
func setGreeting(language string) (prefix string) {
	// switch statements handle multiple if
	switch language {
	case french:
		prefix = bonjour
	case italian:
		prefix = ciao
	default:
		prefix = hello
	}
	return
}

func main() {
	fmt.Println(Hello("Neil", ""))
}
