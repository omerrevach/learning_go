package main

import (
	"fmt"
	"os"
	"errors"
)

func main() {
	_, err := os.ReadFile("baklawa.txt")
	if err!=nil {
		fmt.Print("Error", err)
	}

	numerator := 10
	denominator := 5
	result, remainder, err := intDivision(numerator, denominator)
	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0{
		fmt.Printf("The result of the integer division is %v", result)
	} else {
		fmt.Printf("\nthe result of integer division is: %v, and the remainder is: %v", result, remainder)
	}
}

// return multiple values
func intDivision(numerator, denominator int) (int, int, error)  {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot divide by Zero")
		return 0, 0, err
	}
	
	result := numerator/denominator
	remainder := numerator%denominator
	return result, remainder, err
}

