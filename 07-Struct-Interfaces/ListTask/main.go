package main

import "fmt"

type Task struct {
	name        string
	description string
	completed   bool
}

func (t *Task) toPrint() {
	fmt.Printf("Nombre: %s \nDescripcion: %s \nCompletado: %t\n", t.name, t.description, t.completed)
}

func (t *Task) markCompleted() {
	t.completed = true
}

func main() {

}
