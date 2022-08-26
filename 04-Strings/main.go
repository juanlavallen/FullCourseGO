package main

import (
	"fmt"
	"strings"
)

func convert(word string) {
	fmt.Println(word)
	word = strings.ToLower(word)
	fmt.Println(word)
	word = strings.ToUpper(word)
	fmt.Println(word)
}

func main() {
	convert("Lionel Messi")
}
