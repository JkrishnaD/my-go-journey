package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

// Pointers holds the memory address of the value
// pointers help to avoid copying the huge data by just passing the reference instead
func main() {
	fmt.Println("Pointers")
	var x int = 32
	var p *int = &x
	fmt.Println(p)  // gives the address of the value x
	fmt.Println(*p) // gives the values of the x present in the p addresses
	if x == *p {
		fmt.Println("we are same")
	}
	v := Vertex{1, 2}
	fmt.Println(&v.X, v.Y) // structs can be accesed by using dots
	q := &v
	fmt.Println(q.X)
	//we can allocate the value to the struct
	v1 := Vertex{X: 1}
	v2 := Vertex{Y: 1}
	v3 := Vertex{}
	fmt.Println(v1, v2, v3)
}
