package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello World")
	}()

	myFunc := func() {
		fmt.Println("Hello World")
	}
	myFunc()
}
