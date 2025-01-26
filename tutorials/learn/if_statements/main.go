package main

import (
    "fmt"
)

func main() {
	if 2 > 1 {
		fmt.Println("2 is greater than 1")
	}

	x := 2
	y := 1
	if x > y {
		fmt.Println("x is greater than y")
	}

	time := 9
	if (time < 10) {
		fmt.Println("Good Morning")
	} else if (time < 20) {
		fmt.Println("Good Day")
	} else {
		fmt.Println("Good Night")
	}
}