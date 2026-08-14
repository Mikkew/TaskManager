package main

import (
	"fmt"

	"github.com/mikkew/taskmanager/internal/tasks"
)

func main() {
	// task1 := tasks.NewTask(1, "Aprende Go Modules")
	// task2 := tasks.NewTask(2, "Construir un API REST")

	// fmt.Printf("Tarea %d: %s (Completada: %v)\n", task1.ID, task1.Title, task1.Done)
	// fmt.Printf("Tarea %d: %s (Completada: %v)\n", task2.ID, task2.Title, task2.Done)

	store := tasks.NewStore()
	store.Add("Aprende Go Modules")
	store.Add("Hacer proyecto del curso de Python")
	store.Add("FastAPI desde cero")
	store.Add("DUPLICADO")

	fmt.Println("===MIS TAREAS===")
	for _, task := range store.List() {
		status := "[ ]"

		if task.Done {
			status = "[x]"
		}

		fmt.Printf("%s %d. %s\n", status, task.ID, task.Title)
	}

	fmt.Println("Completar tarea:")
	store.Complete(1)
	store.Complete(0)

	store.Delete(4)

	for _, task := range store.List() {
		status := "[ ]"

		if task.Done {
			status = "[x]"
		}

		fmt.Printf("%s %d. %s\n", status, task.ID, task.Title)
	}
}
