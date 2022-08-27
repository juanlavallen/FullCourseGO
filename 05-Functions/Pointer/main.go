package main

import "fmt"

func change(str *string) {
	fmt.Printf("%p\n", str)
	*str = "Hello World"
}

func main() {
	str := "Hello World!"
	fmt.Printf("%p\n", &str)
	fmt.Println("Before function", str)

	change(&str)
	fmt.Println("After function", str)
}
