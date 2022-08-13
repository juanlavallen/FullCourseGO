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

	students := make(map[string][]int)

	students["Paul"] = []int{7, 9, 10}
	students["Bill"] = []int{10, 9, 10}

	fmt.Println(students)
	fmt.Println(students["Paul"])
	fmt.Println(students["Paul"][2])
}
