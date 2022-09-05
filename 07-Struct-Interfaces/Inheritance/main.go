package main

import "fmt"

type Person struct {
	firstname string
	lastname  string
	age       int
}

type Employee struct {
	Person
	salary     float64
	profession string
}

func (e *Employee) format() {
	fmt.Printf("Name: %s\nAge: %s\nProfession: %s\n", e.firstname, e.lastname, e.profession)
}

func main() {
	employee := Employee{salary: 100, profession: "Programmer"}
	employee.firstname = "Juan"
	employee.lastname = "Lavallen"
	employee.age = 20
	fmt.Println(employee)

	employee.format()
}
