package main

import (
	"fmt"
)

func main() {
	
	// Unbufferd chennel; only one value.
	d := make(chan int)
	
	// Buffered channel; more values in.
	// d := make(chan int, 10)

	// c := make(chan int)

	// go func() {
	// 	for i := 0; i < 6; i++ {
	// 		c <- i
	// 	}
	// }()

	go func() {
		for j := 0; j < 6; j++ {
			d <- j
		}
		// Closing channel. It will not receive anything else,
		// but it can export data from channel.
		close(d)
	}()

	// 	go func() {
	// 		for {
	// 			fmt.Println(<-c)
	// 		}
	// 	}()
	// 	// We need this because there are 2 Go routines and nothing else
	// 	// executed in the main function. Therefore, it might exit before
	// 	// the routines finish
	// 	time.Sleep(2 * time.Second)

	// This time it can't exit if the channel is not closed. It waits for
	// it the channel to receive and then after unloading the value with
	// range (very useful for channels) it will ask for the next one and then
	// the code in the go rutine is going to continue and load the next value
	// to the channel and so on.
	// Only when channel is closed, the main will continue and bare i mind that
	// only one value passes at each time, because we have interuption.
	for n := range d {
		fmt.Println(n)
	}

}
