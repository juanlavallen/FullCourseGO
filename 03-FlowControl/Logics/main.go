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

	number := 2
	result := number == 2 && 4 > number

	fmt.Println(result)
}
