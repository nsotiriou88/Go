package main

import "fmt"

func main() {
	c := incrementor()
	cSum := puller(c)     // Remove this and replace it below. Same results
	for n := range cSum { // for n := range puller(c) => use this as a refactor
		fmt.Println(n)
	}
}

// func incrementor() <-chan int {  // Can only be used to receive int
func incrementor() chan int {
	out := make(chan int)
	// The routine dissapears in the background once launched, therefore there is no
	// "Dead Lock" condition.
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

// func puller(c <-chan int) <-chan int {  // Can only be used to receive int
func puller(c chan int) chan int {
	out := make(chan int)
	// Same applies for here for the receiving end. Go routine dissapears in the background.
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}

/*
The optional <- operator specifies the channel direction, send or receive.
If no direction is given, the channel is bidirectional.
https://golang.org/ref/spec#Channel_types
*/
