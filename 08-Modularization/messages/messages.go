package messages

import "fmt"

func Send() { // Func public
	fmt.Println("Hello World - Public")
}

func send() { // Func private
	fmt.Println("Hello World - Private")
}
