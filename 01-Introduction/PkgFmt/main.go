package main

import "fmt"

func main() {
	name := "Mick Jagger"
	age := 79

	fmt.Printf("Hi, I'm %s. I'm %d years old \n", name, age)

	message := fmt.Sprintf("Hi, I'm %s. I'm %d years old", name, age)
	fmt.Println(message)

	fmt.Printf("name: %T \n", name)
	fmt.Printf("age: %T \n", age)

	fmt.Print("Enter another name: ")
	fmt.Scanln(&name)

	fmt.Println(name)
}
