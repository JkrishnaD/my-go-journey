package main

// these are the packages
import (
	"fmt"
	"math"
	"time"
)

// only one main function should present for one package

/*
  go has the basic types like string
  int-int8,int16,int32,int64.
  float-float32,float64.
  uint-unit8,uint16,uint32,uint64,uintptr.
  byte alias for uint32.
  rune alias for int32.
  complex64,complex128.
*/

// variables which are declared without any explicit value then zero for numbers false for boolean and "" for string
// functions can't be called in the package scope which means outside the function

// type conversions
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

// when value is assigned to variable without type based on the value type is assigned

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
var c, python string
var java bool
var ram, lakshman = "ram", "lakshman" //type can be omitted when it has the initializers
const pi = 3.14                       // we can't chage the value of the const like var
// main function called automatically in a go file all other functions should be called
func main() {
	var i, j int                    // variables inside the function is known as short variables
	k := 4                          // short representaion for both declaring and initializing only used inside the function
	fmt.Println("jaya" + "Krishna") //adding the string
	fmt.Println(time.Now())
	packages()
	fmt.Println(add(2, 3))
	fmt.Println(swap("krishna", "jaya"))
	fmt.Println(c, python, java)
	fmt.Println(ram, lakshman, i, j, k)
	fmt.Printf("i is of type %T\n", i)
	fmt.Printf("f is of type %T\n", f)
	fmt.Printf("u is of type %T\n", u)
}
