package main

import "fmt"

type Animal interface {
	move() string
}

type Dog struct{}

func (*Dog) move() string {
	return "Moving"
}

type Cat struct{}

func (*Cat) move() string {
	return "Moving"
}

type Bird struct{}

func (*Bird) move() string {
	return "Moving"
}

func MoveAnimal(animal Animal) {
	fmt.Println(animal.move())
}

func main() {
	dog := Dog{}
	cat := Cat{}
	bird := Bird{}

	MoveAnimal(&dog)
	MoveAnimal(&cat)
	MoveAnimal(&bird)
}
