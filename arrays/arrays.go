package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays")

	var a [10]string // array initialization
	fmt.Println(a)
	a[1] = "jaya"
	fmt.Println(a)

	primes := [...]int{2, 3, 5, 7, 11}
	fmt.Println(primes[2])
	fmt.Println(primes[1:4]) //slicing

	names := [4]string{
		"bgmi",
		"coc",
		"cricket",
		"vscode",
	}
	fmt.Println(names)
	b := names[1:3]
	b[0] = "hello" //slices can also act as the references
	fmt.Println(b)
	fmt.Println(names)
	fmt.Println(primes[1:])

	fmt.Printf("len=%d cap=%d %v\n", len(primes[1:]), cap(primes), primes)
	// make function it is a built-in function which allocates a zeroed array
	z := make([]int, 10)
	fmt.Println(z, "hi", "hello this is jaya")
}
