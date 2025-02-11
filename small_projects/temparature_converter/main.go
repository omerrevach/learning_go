package main

import (
	"fmt"
	"strings"
)

func main() {
	var temp float64
	var fromUnit, toUnit string

	fmt.Print("Enter temperature: ")
	fmt.Scanln(&temp)

	fmt.Print("Enter current unit (C, F, K): ")
	fmt.Scanln(&fromUnit)
	fromUnit = strings.ToUpper(fromUnit)

	fmt.Print("Enter a unit to convert to (C, F, K): ")
	fmt.Scanln(&toUnit)
	toUnit = strings.ToUpper(toUnit) // Fix here

	convertedTemp := 0.0

	switch fromUnit {
	case "C":
		if toUnit == "F" {
			convertedTemp = temp*9/5 + 32 // Celsius to Fahrenheit
		} else if toUnit == "K" { // Fix here
			convertedTemp = temp + 273.15 // Celsius to Kelvin
		}
	case "F":
		if toUnit == "C" {
			convertedTemp = (temp - 32) * 5 / 9
		} else if toUnit == "K" {
			convertedTemp = (temp - 32)*5/9 + 273.15
		}
	case "K":
		if toUnit == "C" {
			convertedTemp = temp - 273.15
		} else if toUnit == "F" {
			convertedTemp = (temp - 273.15)*9/5 + 32
		}
	default:
		fmt.Println("Invalid unit entered.")
		return
	}

	fmt.Printf("Converted temperature: %.2f %s\n", convertedTemp, toUnit)
}
