package main

import "fmt"

func main() {
	numbers := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	for i := 0; i < len(numbers); i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
