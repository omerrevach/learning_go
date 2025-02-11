package main

import (
	"fmt"
)

func main() {
	var a []int = []int{1, 2, 3 ,4, 5, 6, 7}

	for i := range a {
		fmt.Println(a[i])
	}
}