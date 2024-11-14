package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	defer fmt.Println("defer says hello") // it runs after the execution of all the statements
	fmt.Println("Loops")
	sum := 1
	mul := 1
	for i := 0; i < 10; i++ {
		sum += i
	}
	for j := 1; j < 5; j++ {
		mul *= j
	}
	fmt.Println(sum)
	fmt.Println(mul)
	//without initializer and the increments it act's as while loop
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
	//with the fixed range
	for i := range 3 {
		fmt.Println("range", i)
	}
	var n int
	for n = range 8 {
		if n%2 == 0 {
			fmt.Println(n, "is even number")
		}
	}
	//switch cases these are evaluated from top to bottom, stopping when a case succeeds
	switch os := runtime.GOARCH; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Println(runtime.GOARCH)
	}
	today := time.Now().Weekday()
	fmt.Println(today)
	switch time.Saturday {
	case today + 1:
		fmt.Println(today+1, "is tomorrow")
	case today + 2:
		fmt.Println(today+2, "is day after tomorrow")
	}
	//switch without a condition
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good night")
	}
}
