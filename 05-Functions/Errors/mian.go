package main

import (
	"errors"
	"fmt"
)

func div(num1 int, num2 int) (int, error) {
	if num2 == 0 {
		return 0, errors.New("Error")
	} else {
		return num1 / num2, nil
	}
}

func main() {
	result, error := div(10, 0)
	if error == nil {
		fmt.Println(result)
	} else {
		fmt.Println(error)
	}
}
