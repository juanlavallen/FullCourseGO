package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers, len(numbers))

	numbers = append(numbers, 4)
	fmt.Println(numbers, len(numbers))
}
