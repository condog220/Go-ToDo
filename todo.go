package main

import (
	"fmt"
	"time"
)

type Task struct {
	task         string
	completed    bool
	created_at   time.Time
	completed_at *time.Time
}

type tasks []Task

func (toDos *tasks) add(task string) {
	todo := Task{
		task:         task,
		completed:    false,
		created_at:   time.Now(),
		completed_at: nil,
	}
	*toDos = append(*toDos, todo)
}

func (toDos *tasks) remove(index int) {
	*toDos = append((*toDos)[:index], (*toDos)[index+1:]...)
}

func (toDos *tasks) update(index int, task string) {
	(*toDos)[index].task = task
}

func (toDos *tasks) markCompleted(index int) {
	(*toDos)[index].completed = true
	completedTime := time.Now()
	(*toDos)[index].completed_at = &completedTime
}

func (toDos *tasks) list() []Task {
	for index, task := range *toDos {
		fmt.Printf("Task %d:\n", index+1)
		fmt.Println(task.task)
		fmt.Println(task.created_at.Format(time.RFC1123))
		fmt.Println(task.completed)
		fmt.Println(task.completed_at)
	}
	return *toDos
}
