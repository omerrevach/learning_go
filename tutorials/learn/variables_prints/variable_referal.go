package main

import "fmt"

func main() {
	name := "Omer"
	i := 3.141
	fmt.Printf("%.2f\n", i) // prints only the last 2 digits after decimal
	fmt.Printf("Hello %[1]v. %[1]v has a value of %[1]T", name) // [1] refers to the first arguments and makes it that i dont need to type name 3 times
}