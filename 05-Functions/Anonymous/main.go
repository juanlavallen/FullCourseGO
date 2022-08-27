package main

import (
	"fmt"
	"strings"
)

// Closures
func repeat(n int) func(text string) string {
	return func(text string) string {
		return strings.Repeat(text, n)
	}
}

func main() {
	func() {
		fmt.Println("Hello World")
	}()

	myFunc := func() {
		fmt.Println("Hello World")
	}
	myFunc()

	result := repeat(2)
	fmt.Println(result("Golang"))
}
