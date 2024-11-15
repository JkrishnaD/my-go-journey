package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("Enter first number:")
	fmt.Scanln(&num1)
	fmt.Println("Enter the operator:")
	fmt.Scanln(&operator)
	fmt.Println("Enter second number:")
	fmt.Scanln(&num2)
	switch operator {
	case "+":
		fmt.Println(add(num1, num2))
	case "-":
		fmt.Println(sub(num1, num2))
	case "*":
		fmt.Println(mul(num1, num2))
	case "/":
		fmt.Println(div(num1, num2))
	default:
		fmt.Println("Invalid operation")
	}
}

func add(a, b float64) float64 { return a + b }
func sub(a, b float64) float64 { return a - b }
func mul(a, b float64) float64 { return a * b }
func div(a, b float64) float64 { return a / b }
