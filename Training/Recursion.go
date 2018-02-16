package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
}

func factorial(x int) int {
	if x == 0 { // We can use 1 to save one cycle.
		return 1
	}
	return x * factorial(x-1)
}

// Recursive function that stops at some point.
// When we hit return, it returns the value to
// the function and exits immediately from it.
