//  responsible for saving tasks to a file and loading tasks from a file. This ensures that the tasks persist even after the program is closed

package main

import (
	"encoding/json"
	"os"
)

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile("tasks.json", data, 0666)
}

func loadTasks() ([]Task, error) {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		// If the file doesnt exist then return an empty slice
		if os.IsNotExist(err) {
			return []Task{}, nil // tasks are empty
		}
		return nil, err // Return err if its another issue
	}

	// Conver Json into go struct
	var tasks []Task
	if len(data) == 0 {
		return []Task{}, nil
	}
	err = json.Unmarshal(data, &tasks)
	if nil != err {
		return nil, err // if json corrupted return error
	}

	return tasks, nil
}