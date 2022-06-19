package main

import (
	"fmt"
)

const sayHello = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return sayHello + name
}

func main() {
	fmt.Println(Hello("Neil"))
}
