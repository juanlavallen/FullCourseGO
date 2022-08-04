package main

import "fmt"

func hello(name string) {
	fmt.Println("Hello", name)
}

func multiplication(a, b int) int {
	return a * b
}

func main() {
	hello("Mick")
	fmt.Println(multiplication(2, 9))
}
