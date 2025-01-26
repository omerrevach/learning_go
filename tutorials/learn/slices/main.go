// Slices are similar to arrays, but are more powerful and flexible
// Like arrays, slices are also used to store multiple values of the same type in a single variable.
// However, unlike arrays, the length of a slice can grow and shrink as you see fit

package main

import (
	"fmt"
)

func main() {
	myslice1 := [...]int{1}
	fmt.Println(myslice1)
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))

	myslice2 := []string{"I", "like", "pie"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)

	// create slice from array
	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice3 := arr1[0:3]
	fmt.Printf("myslice = %v\n", myslice3)
	fmt.Printf("length = %d\n", len(myslice3))
	fmt.Printf("cap = %d", cap(myslice3))
}