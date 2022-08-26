package main

import "fmt"

func main() {
	users := [...]string{"Strider", "Snake07", "Mike32"}

	for i := 0; i < len(users); i++ {
		fmt.Println(users[i])
	}

	for index, element := range users {
		fmt.Println(index)
		fmt.Println(element)
	}

	for _, element := range users {
		fmt.Println(element)
	}
}
