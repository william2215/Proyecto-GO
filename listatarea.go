package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) addList(nuevaTarea *task) {
	t.tasks = append(t.tasks, nuevaTarea)
}

type task struct {
	nombre      string
	description string
	compleado   bool
}

func (t *task) marcarCompleta() {
	t.compleado = true
}

func (t *task) actualizarDescription(description string) {
	t.description = description
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func main() {
	tarea1 := &task{
		nombre:      "task1",
		description: "Esto es para completar tarea",
	}
	tarea2 := &task{
		nombre:      "task2",
		description: "Esto es para completar tarea",
	}
	tarea3 := &task{
		nombre:      "Completar Tarea GO",
		description: "Esto es para completar tarea",
	}
	lista := &taskList{
		tasks: []*task{
			tarea1, tarea2, tarea3,
		},
	}
	lista.addList(tarea3)
	fmt.Printf("%+v\n", lista.tasks[0], lista.tasks[1], lista.tasks[2])
}
