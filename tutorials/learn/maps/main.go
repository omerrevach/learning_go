// Maps are used to store data values in key:value pairs.

// Each element in a map is a key:value pair.

// A map is an unordered and changeable collection that does not allow duplicates.

// The length of a map is the number of its elements. You can find it using the len() function.

// The default value of a map is nil.

// Maps hold references to an underlying hash table.

// Go has multiple ways for creating maps.

package main

import (
	"fmt"
)

// var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
// b := map[KeyType]ValueType{key1:value1, key2:value2,...}

func main() {
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	var myMap1 = map[string]uint8{"Adam":23, "Shlomi":45}
	fmt.Println(myMap1["Shlomi"])
	fmt.Print(a)

	intArr := [...]int{1, 2, 3, 4, 5}

	// delete(myMap1, "Adam") // deletes from the map

	var age, ok = myMap1["Shlomi"] // ok acts as a boolean
	if ok {
		fmt.Printf("\nThe age is %v", age)
	} else {
		fmt.Printf("\nInvalid name")
	}

	for name := range(myMap1) {
		fmt.Printf("\nName: %v\n", name)
	}

	// iterate over every index and value in an array
	for index, value := range intArr {
		fmt.Printf("Index: %v, Value: %v \n", index, value)
	}
}