package main

import (
	"encoding/json"
	"os"
)

func SavetoJson(filename string, toDos tasks) error {
	data, err := json.MarshalIndent(toDos, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func LoadfromJson(filename string) (tasks, error) {
	var toDos tasks
	data, err := os.ReadFile(filename)
	if err != nil {
		return toDos, err
	}

	err = json.Unmarshal(data, &toDos)
	return toDos, err
}
