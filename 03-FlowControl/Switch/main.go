package main

import "fmt"

func main() {
	var answer string

	fmt.Print("Insert Answer: ")
	fmt.Scanln(&answer)

	switch answer {
	case "Yes":
		fmt.Println("Correct")
	case "No":
		fmt.Println("Incorrect")
	default:
		fmt.Println("Invalid Answer")
	}
}
