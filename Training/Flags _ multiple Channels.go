package main

import (
	"fmt"
	// "sync"
)

func main() {

	// Without synce, using flags us a second channel.
	// Now, there is no guarantee about the order of the
	// values coming out, but they will all be printed.
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 4; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 6; i++ {
			c <- i
		}
		done <- true
	}()

	// Notice the structure here.
	go func() {
		// Just throwing values off the channel; no blank
		// identifier "_" needed.
		<-done
		<-done
		close(c)
	}()

	// If we go like this, the program will hang.
	// <-done
	// <-done
	// close(c)

	for n := range c {
		fmt.Println(n)
	}

	// //////////////////////////////////////////////
	// For launching many routines; more advanced! //
	// //////////////////////////////////////////////

	// n := 10
	// c := make(chan int)
	// done := make(chan bool)

	// for i := 0; i < n; i++ {
	// 	go func() {
	// 		for i := 0; i < 10; i++ {
	// 			c <- i
	// 		}
	// 		done <- true
	// 	}()
	// }

	// go func() {
	// 	for i := 0; i < n; i++ {
	// 		<-done
	// 	}
	// 	close(c)
	// }()

	// for n := range c {
	// 	fmt.Println(n)
	// }

	// /////////////////////////
	// Using sync and waitgroups
	// /////////////////////////

	// c := make(chan int)

	// var wg sync.WaitGroup
	// wg.Add(2)

	// go func() {
	// 	for i := 0; i < 4; i++ {
	// 		c <- i
	// 	}
	// 	wg.Done()
	// }()

	// go func() {
	// 	for i := 0; i < 6; i++ {
	// 		c <- i
	// 	}
	// 	wg.Done()
	// }()

	// go func() {
	// 	wg.Wait()
	// 	close(c)
	// }()

	// for n := range c {
	// 	fmt.Println(n)
	// }
}
