package main

import (
	"fmt"
	"os"
)

func main() {

	if file, error := os.Open("hello.txt"); error != nil {
		fmt.Println("File not found")
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
