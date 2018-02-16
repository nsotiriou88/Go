package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}

// When we need to range over a channel, we need to
// close it also, otherwise we might end with a deadlock.

// If you put something on the channel, you need to have
// something to receive it, otherwise we have a deadlock.

// Always watchout for deadlocks and if exiting main()
// before we finish all of our tasks.

// ==================================

// 	c1 := incrementor("Foo:")
// 	c2 := incrementor("Bar:")
// 	c3 := puller(c1)
// 	c4 := puller(c2)
// 	fmt.Println("Final Counter:", <-c3+<-c4)
// }

// func incrementor(s string) chan int {
// 	out := make(chan int)
// 	go func() {
// 		for i := 0; i < 20; i++ {
// 			out <- 1
// 			fmt.Println(s, i)
// 		}
// 		close(out)
// 	}()
// 	return out
// }

// func puller(c chan int) chan int {
// 	out := make(chan int)
// 	go func() {
// 		var sum int
// 		for n := range c {
// 			sum += n
// 		}
// 		out <- sum
// 		close(out)
// 	}()
// 	return out
// }
