package main

import (
	"fmt"
	"os"
	"strconv"
)

var pl = fmt.Println

func main() {

	tasks, err := loadTasks()
	if err != nil {
		pl("Error loading tasks", err)
		return
	}

	if len(os.Args) < 2 {
		pl("Usage: go run main.go [add|list|done|remove] [task]")
		return
	}

	// Get the title form the cli
	title := os.Args[2]

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			pl("Error, you need to enter a task")
			return
		}

		// CReate a new Task with a unique ID
		newTask := Task {
			ID: len(tasks) + 1,
			Title: title,
			Completed: false,
		}

		// Append the task to the new list
		tasks = append(tasks, newTask)

		err = saveTasks(tasks)
		if err != nil {
			pl("Error saving tasks")
			return
		}

		pl("Task Added: ", title)
//--------------------------------------------------------
	case "list":
		if len(tasks) == 0 {
			pl("No Tasks Found")
			return
		}

		pl("Tasks: ")
		for _, task := range tasks {
			pl(task) // Uses Task's String() method automatically
		}

//--------------------------------------------------------
	case "done":
		if len(os.Args) < 3 {
			pl("Error, you need to enter a Task ID")
			return
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			pl("Task Id is Invalid")
			return
		}

		found := false
		for i, task := range tasks {
			if task .ID == id {
				tasks[i].Completed = true
				found = true
				break
			}
		}

		if !found {
			pl("Task not found.")
			return
		}
	
		err = saveTasks(tasks)
		if err != nil {
			pl("Error saving tasks:", err)
			return
		}
	
		pl("Task marked as completed.")
//--------------------------------------------------------
	case "remove":
		if len(os.Args) < 3 {
			pl("Error, you need to Provide task ID")
			return
		}

		newTasks := []Task{}
		found:= false

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			pl("Task ID Is Invalid")
		}

		for _, task := range tasks {
			if task.ID == id {
				found = true
				continue
			}

			newTasks = append(newTasks, task)
		}

		if !found {
			pl("Task not found")
			return
		}

		err = saveTasks(newTasks)
		if err != nil {
			pl("Error saving tasks: ", err)
			return
		}

		pl("Task Removed")

//--------------------------------------------------------
	default:
        fmt.Println("Invalid command. Use: add, list, done, remove")
	}
}