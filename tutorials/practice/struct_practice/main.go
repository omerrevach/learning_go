package main

import (
	"time"
)

type Book struct {
	title string
	author string
	numPages int

	isSaved bool
	savedAt time.Time
}

func (book *Book) saveBook() {
	book.isSaved = true
	book.savedAt= time.Now()
}

func saveBook(book *Book) {
	book.isSaved = true
	book.savedAt= time.Now()
}

// 1. read data
// 2. write data

func main() {

}