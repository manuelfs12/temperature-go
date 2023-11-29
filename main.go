package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var userInput string
	fmt.Println("Temperature GO, a CLI tool to convert temperatures.")
	fmt.Println("#------------------------------------------------#")
	fmt.Println("Accepted inputs: <number><c | f | k>, Eg: 50c ")
	fmt.Println("To exit the application, type \"exit\"")
	for !strings.Contains(userInput, "exit") {
		fmt.Printf("Please enter a temperature: ")
		_, err := fmt.Scan(&userInput)
		if err != nil {
			return
		}
		userInput = strings.ToLower(userInput)
		if strings.Contains(userInput, "c") {
			userInput = strings.Trim(userInput, "c")
			convertedInput, err := strconv.ParseFloat(userInput, 32)
			if err != nil {
				fmt.Println("Invalid input")
				continue
			}
			convertCelsius(userInput, convertedInput)
		} else if strings.Contains(userInput, "f") {
			userInput = strings.Trim(userInput, "f")
			convertedInput, err := strconv.ParseFloat(userInput, 32)
			if err != nil {
				fmt.Println("Invalid input")
				continue
			}
			convertFahrenheit(userInput, convertedInput)
		} else if strings.Contains(userInput, "k") {
			userInput = strings.Trim(userInput, "k")
			convertedInput, err := strconv.ParseFloat(userInput, 32)
			if err != nil {
				fmt.Println("Invalid input")
				continue
			}
			convertKelvin(userInput, convertedInput)
		} else if strings.Contains(userInput, "exit") {
			return
		} else {
			fmt.Println("Invalid input")
		}
	}
}

func convertCelsiusToKelvin(temp float64) string {
	return fmt.Sprintf("%.2f", temp+273.15)
}

func convertCelsiusToFahrenheit(temp float64) string {
	return fmt.Sprintf("%.2f", (temp*9/5)+32)
}

func convertKelvinToCelsius(temp float64) string {
	return fmt.Sprintf("%.2f", temp-273.15)
}

func convertKelvinToFahrenheit(temp float64) string {
	return fmt.Sprintf("%.2f", (temp-273.15)*9/5+32)
}

func convertFahrenheitToCelsius(temp float64) string {
	return fmt.Sprintf("%.2f", (temp-32)*5/9)
}

func convertFahrenheitToKelvin(temp float64) string {
	return fmt.Sprintf("%.2f", (temp-32)*5/9+273.15)
}

func convertCelsius(userInput string, convertedInput float64) {
	fmt.Printf("%v° Celsius is %v° Kelvin\n", userInput, convertCelsiusToKelvin(convertedInput))
	fmt.Printf("%v° Celsius is %v° Fahrenheit\n", userInput, convertCelsiusToFahrenheit(convertedInput))
}

func convertKelvin(userInput string, convertedInput float64) {
	fmt.Printf("%v° Kelvin is %v° Celsius\n", userInput, convertKelvinToCelsius(convertedInput))
	fmt.Printf("%v° Kelvin is %v° Fahrenheit\n", userInput, convertKelvinToFahrenheit(convertedInput))
}

func convertFahrenheit(userInput string, convertedInput float64) {
	fmt.Printf("%v° Fahrenheit is %v° Celsius\n", userInput, convertFahrenheitToCelsius(convertedInput))
	fmt.Printf("%v° Fahrenheit is %v° Kelvin\n", userInput, convertFahrenheitToKelvin(convertedInput))
}
