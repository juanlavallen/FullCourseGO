package main

import "fmt"

func main() {
	characters := make(map[int]string)

	// Add Data
	characters[1] = "Homer"
	characters[2] = "Marge"
	characters[3] = "Bart"
	characters[4] = "Lisa"
	characters[5] = "Maggie"

	fmt.Println(characters)

	characters[5] = "Moe"
	fmt.Println(characters)

	delete(characters, 5)
	fmt.Println(characters)
}
