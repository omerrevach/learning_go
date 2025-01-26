package main

import (
	"fmt"
)

func main() {
	array1 := [5]int{1,2,3,4,5} // here i define the size of the array
	array2 := [...]int{1,2,3,4,5} // go decides the size of the array and makes this dynamic
	array3 := [5]int{} // not initialized so it will print [0 0 0 0 0]
	array4 := [...]int{} // this will initialize an empty array
	fmt.Println(array1)
	fmt.Println(array2)

	array1[2] = 10 // to change the value of an item in the array
	fmt.Println(array2[0]) // fetches the first item in the array
	fmt.Println(array1[2])
	fmt.Println(array3)
	fmt.Print(array4)
}