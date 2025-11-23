package main

import "fmt"

func main() {
	var toDos tasks

	toDos.add("Learn GoLang")

	taskList := toDos.list()

	fmt.Println(taskList)
}
