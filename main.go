package main

import (
	"fmt"
	"strconv"
)

func main() {
	var userInput string
	fmt.Println("Temperature GO, a CLI tool to convert temperatures.")
	fmt.Printf("Please enter a temperature: ")
	_, err := fmt.Scan(&userInput)
	if err != nil {
		return
	}
	convertedInput, _ := strconv.ParseFloat(userInput, 32)
	fmt.Printf("%v° celsius is %v° kelvin\n", userInput, convertCelsiusToKelvin(convertedInput))
	fmt.Printf("%v° celsius is %v° fahrenheit\n", userInput, convertCelsiusToFahrenheit(convertedInput))
	fmt.Printf("%v Kelvin us %v Celsisus", userInput, convertKelvinToFahrenheit(convertedInput))
}

func convertCelsiusToKelvin(temp float64) float64 {
	return temp + 273.15
}

func convertCelsiusToFahrenheit(temp float64) float64 {
	return (temp * 9 / 5) + 32
}

func convertKelvinToCelsius(temp float64) float64 {
	return temp - 273.15
}

func convertKelvinToFahrenheit(temp float64) float64 {
	return (temp-273.15)*9/5 + 32
}
