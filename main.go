package main

import (
	"fmt"
	"strconv"
)

func main() {
	var userInput string
	fmt.Println("Temperature GO, a CLI tool to convert temperatures.")
	fmt.Println("Please enter a temperature:")
	fmt.Scan(&userInput)
	convertedInput, _ := strconv.ParseFloat(userInput, 32)
	fmt.Printf("%v° celsius is %v° kelvin\n", userInput, convertCelsiusToKelvin(convertedInput))
	fmt.Printf("%v° celsius is %v° fahrenheit\n", userInput, convertCelsiusToFahrenheit(convertedInput))
}

func convertCelsiusToKelvin(temp float64) float64 {
	return temp + 273.15
}

func convertCelsiusToFahrenheit(temp float64) float64 {
	return temp*(9/5) + 32
}
