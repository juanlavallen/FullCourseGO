package main

import "fmt"

func factory(n int) int {
	if n == 0 {
		return 1
	}
	f := n * factory(n-1)
	return f
}

func main() {
	result := factory(5)
	fmt.Println(result)
}
