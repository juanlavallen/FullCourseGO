package main

import "fmt"

func main() {
	a := 5
	b := 6

	var result bool

	result = a == b
	fmt.Println(result)

	result = a != b
	fmt.Println(result)

	result = a > b
	fmt.Println(result)

	result = a < b
	fmt.Println(result)

	result = a >= b
	fmt.Println(result)

	result = a <= b
	fmt.Println(result)
}
