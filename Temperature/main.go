package main

import (
	"fmt"
)

func main() {
	var temp float64
	var conversion string
	fmt.Println("Enter Temperature to convert:")
	fmt.Scanln(&temp)
	fmt.Println("Click c for celsius or f for farenheit ")
	fmt.Scanln(&conversion)
	switch conversion {
	case "f":
		fmt.Println(celsiusToFarenheit(temp))
	case "c":
		fmt.Println(farenheitToCelcius(temp))
	default:
		fmt.Println("Invalid coversion choose c or f ")
	}
}

func celsiusToFarenheit(num float64) float64 {
	return (num * 9 / 5) + 32
}
func farenheitToCelcius(num float64) float64 {
	return (num - 32) * 5 / 9
}
