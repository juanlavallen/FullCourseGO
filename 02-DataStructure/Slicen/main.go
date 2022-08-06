package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers, len(numbers))

	numbers = append(numbers, 4)
	fmt.Println(numbers, len(numbers))

	calendar := []string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"Octuber",
		"November",
		"December",
	}

	fmt.Printf("Len: %v, Cap: %v, %p \n", len(calendar), cap(calendar), calendar)
}
