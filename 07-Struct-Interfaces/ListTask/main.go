package main

import "fmt"

type TaskList struct {
	tasks []*Task
}

func (tl *TaskList) AppendTask(task *Task) {
	tl.tasks = append(tl.tasks, task)
}

func (tl *TaskList) RemoveTask(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

type Task struct {
	name        string
	description string
	completed   bool
}

func (t *Task) ToPrint() {
	fmt.Printf("Nombre: %s \nDescripcion: %s \nCompletado: %t\n", t.name, t.description, t.completed)
}

func (t *Task) MarkCompleted() {
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

	firstTask.ToPrint()
	secondTask.ToPrint()
}
