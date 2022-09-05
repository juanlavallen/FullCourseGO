package main

import (
	"fmt"
	"os"
)

func main() {

	defer func() {
		if error := recover(); error != nil {
			fmt.Println("An error occurs")
		}
	}()

	if file, error := os.Open("hello.txt"); error != nil {
		panic("File not found")
	} else {

		defer func() {
			fmt.Println("File closed")
			file.Close()
		}()

		content := make([]byte, 254)
		long, _ := file.Read(content)
		fileContent := string(content[:long])

		fmt.Println(fileContent)
	}
}
