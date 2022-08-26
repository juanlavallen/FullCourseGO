package main

import "fmt"

func sum(numbers ...int) int {
	var total int
	for _, number := range numbers {
		total += number
	}
	return total
}

func sum2(name string, numbers ...int) (string, int) {
	message := fmt.Sprintf("The result of %s sum is", name)
	var total int
	for _, number := range numbers {
		total += number
	}
	return message, total
}

func main() {
	result := sum(150, 350)
	fmt.Println(result)

	message, result := sum2("Lionel", 2, 10)
	fmt.Println(message, result)
}
