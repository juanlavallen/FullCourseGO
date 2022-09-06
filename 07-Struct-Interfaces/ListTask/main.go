package main

import "fmt"

type TaskList struct {
	tasks []*Task
}

func (tl *TaskList) appendTask(task *Task) {
	tl.tasks = append(tl.tasks, task)
}

func (tl *Task) remove(task *Task) {
}

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
	firstTask := Task{
		name:        "Curso de GO",
		description: "Backend con GO",
		completed:   true,
	}

	secondTask := Task{
		name:        "Curso de Rust",
		description: "Backend con Rust",
		completed:   false,
	}

	firstTask.toPrint()
	secondTask.toPrint()
}
