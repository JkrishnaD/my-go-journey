package main

// these are the packages

import (
	"fmt"
	"math"
	"time"
)

// main function called automatically in a go file all other functions should be called
// functions can't be called in the package scope which means outside the function

func packages() {
	fmt.Println(math.Pi) // every exported name must start with the capital
}

// contrast to other languages type comes after the variable name
func add(x, y int) int {
	return x + y // if function returns something it's type should be mentioned
}

func swap(x, y string) (string, string) {
	return y, x // a function can return multiple values
}

func main() {
	fmt.Println("hello world")
	fmt.Println("jaya" + "Krishna") //adding the string
	fmt.Println(time.Now())
	packages()
	fmt.Println(add(2, 3))
	fmt.Println(swap("krishna", "jaya"))
}

// simillar to other languages go has the values like string,integers,float and boolean
