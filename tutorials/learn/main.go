package main

import (
	"fmt"
	"os"
	"strconv"
)

var pl = fmt.Println
var pf = fmt.Printf



func main() {
	pl(os.Args)
	args := os.Args[1:]
	var iArgs = []int{}
	for _, i := range args {
		val, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		iArgs = append(iArgs, val)
	}
	max := 0
	for _, val := range iArgs {
		if val > max {
			max = val
		}
	}
	pl("Max Value : ", max)

	// to run this
	// go build main.go
	// ./main 1 2 55 6 88
}