package main

import "fmt"

func main() {
	var name string = "Juan"
	var age int = 19

	street := 310
	email := "email@email.com"

	fmt.Println(name, age)
	fmt.Println(street, email)

	var a int     // 0
	var b float64 // 0
	var c string  // ""
	var d bool    // false

	fmt.Println(a, b, c, d)

	const pi = 3.141592
	fmt.Println(pi)
}
