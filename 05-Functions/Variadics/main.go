package main

import "fmt"

func sum(numbers ...int) int {
	var total int
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	result := sum(150, 350)
	fmt.Println(result)
}
