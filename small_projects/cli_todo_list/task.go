package main

import (
    "fmt"
)

type Task struct {
	ID int
	Title string
	Completed bool
}

func (t Task) String() string {
	status := "Not Completed"
	if t.Completed {
		status = "Completed"
	}

	return fmt.Sprintf("%d: %s [%s]", t.ID, t.Title, status) // returns a formatted string
}