package main

import (
	"fmt"
)

// for statement1; statement2; statement3 {
//
// }

// statement1 Initializes the loop counter value.
// statement2 Evaluated for each loop iteration. If it evaluates to TRUE, the loop continues. If it evaluates to FALSE, the loop ends.
// statement3 Increases the loop counter value.

func main() {
	for i:=0; i < 5; i++ {
		fmt.Println(i)
	}

	for i:=0; i <= 100; i+=10 {
		fmt.Println(i)
	}

	// Continue
	for i:=0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	// Break
	for i:=0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}

	// Nested Loops
	adj := [...]string{"big", "tasty"}
	fruits := [...]string{"apple", "orange", "banana"}
	for i:=0; i <len(adj); i++ {
		for j:=0; j < len(fruits); j++ {
			fmt.Println(adj[i], fruits[j])
		}
	}

	// Range
	//for index, value := range array|slice|map {
   	// 
   	// }

	vegetables := [...]string{"cucumber", "lettuce", "carrot"}
	for index, value := range vegetables {
		fmt.Printf("%v\t%v\n", index, value)
	}

	// use _ to ignore the output of value or the index
	cars := [...]string{"bmw", "nissan", "kia"}
	for _, value := range cars {
		fmt.Printf("%v\n", value)
	}

	colors := [...]string{"red", "blue", "yellow"}
	for index, _ := range colors {
		fmt.Printf("%v\n", index)
	}
}