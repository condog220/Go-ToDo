package main

import (
	"fmt"
	"strings"
	"time"
)

type Task struct {
	Task        string     `json:"task"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

type tasks []Task

func (toDos *tasks) add(task string) {
	todo := Task{
		Task:        strings.TrimSpace(task),
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	*toDos = append(*toDos, todo)
}

func (toDos *tasks) remove(index int) {
	*toDos = append((*toDos)[:index], (*toDos)[index+1:]...)
}

func (toDos *tasks) update(index int, task string) {
	(*toDos)[index].Task = task
}

func (toDos *tasks) markCompleted(index int) {
	(*toDos)[index].Completed = true
	completedTime := time.Now()
	(*toDos)[index].CompletedAt = &completedTime
}

func (toDos *tasks) list() []Task {
	for index, task := range *toDos {
		fmt.Println("Task", index+1)
		fmt.Println(task.Task)
		fmt.Println(task.Completed)
		fmt.Println(task.CreatedAt.Format(time.RFC1123))
		if task.CompletedAt != nil {
			fmt.Println(task.CompletedAt.Format(time.RFC1123))
		} else {
			fmt.Println("Not completed yet")
		}
		fmt.Println("-------------------")
	}
	return *toDos
}
