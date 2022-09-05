package main

import "fmt"

// Struct
type Person struct {
	name string
	age  int
}

// Method
func (p *Person) imprimer() {
	fmt.Printf("Name: %s\nAge: %d\n", p.name, p.age)
}

func (p *Person) edit(name string) {
	p.name = name
}

func main() {
	// Instance
	// person := Person{"Paul", 34}
	person := Person{name: "Paul", age: 34}
	fmt.Println(person)

	// Call Method
	person.imprimer()

	// Change method name
	person.edit("Bill")
	person.imprimer()
}
