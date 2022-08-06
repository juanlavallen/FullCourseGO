package main

import "fmt"

func main() {
	characters := make([]string, 0, 5)

	characters = append(characters, "Homer")

	fmt.Println(characters, len(characters), cap(characters))

	numbers := make([]int, 5, 5)

	numbers[0] = 100
	numbers[1] = 200
	numbers[2] = 300
	numbers[3] = 400
	numbers[4] = 500

	numbers = append(numbers, 600)

	fmt.Println(numbers, len(numbers), cap(numbers))
}
