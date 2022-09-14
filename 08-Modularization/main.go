package main

import (
	"fmt"
	"fullcourse/messages"
	"fullcourse/models"
)

func main() {
	messages.Send()

	person := models.Person{}
	person.Constructor("Juan", 20)
	fmt.Println(person)

	person.SetName("Lionel")
	fmt.Println(person.GetName())
}
