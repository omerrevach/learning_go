package main

import (
	"fmt"
)

func fullName(firstName string, familyName string, age int) {
	fmt.Printf("Hello %v %v, you are %d years old\n", firstName, familyName, age)
}

func calculate(x int, y int) int {
	sum := x + y
	return sum
}

func returnValue(x int, y int) (result int) {
	result = x + y
	return
}

func multipleReturnValues(x int, y string) (result int, text string) {
	result = x + x
	text = y + " bomboclat"
	return
}

func main() {
	fullName("omer", "revach", 22)
	fmt.Println(calculate(5, 10))
	total := returnValue(2, 2)
	fmt.Println(total)

	fmt.Println(multipleReturnValues(5, "hello"))
}