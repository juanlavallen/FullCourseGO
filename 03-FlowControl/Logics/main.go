package main

import "fmt"

func main() {
	// NOT
	fmt.Println(!true)  // => false
	fmt.Println(!false) // => true

	// AND
	fmt.Println(true && true)  // => true
	fmt.Println(false && true) // => false

	// OR
	fmt.Println(true || true)  // => true
	fmt.Println(true || false) // => true
}
