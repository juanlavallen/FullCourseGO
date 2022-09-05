package main

import "fmt"

// Struct
type Person struct {
	name string
	age  int
}

func main() {
	// Instance
	// person := Person{"Paul", 34}
	person := Person{name: "Paul", age: 34}
	fmt.Println(person)
}
