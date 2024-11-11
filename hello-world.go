package main

// these are the packages
import (
	"fmt"
	"math"
	"time"
)

// simillar to other languages go has the values like string,integers,float and boolean
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

// variables can be both package and function level
var c, python, java bool              // these are booleans
var ram, lakshman = "jaya", "krishna" //type can be omitted when it has the initializers

// main function called automatically in a go file all other functions should be called
func main() {
	var i, j = 1, 2 // variables inside the function is known as short variables
	k := 4          // short representaion for both declaring and initializing
	fmt.Println("hello world")
	fmt.Println("jaya" + "Krishna") //adding the string
	fmt.Println(time.Now())
	packages()
	fmt.Println(add(2, 3))
	fmt.Println(swap("krishna", "jaya"))
	fmt.Println(c, python)
	fmt.Println(ram, lakshman, i, j, k)
}
