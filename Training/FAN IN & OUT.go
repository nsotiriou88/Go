package main

import (
	"fmt"
	"sync"
)

func main() {

	in := gen()

	// FAN OUT
	// Multiple functions reading from the same channel until that channel is closed
	// Distribute work across multiple functions (ten goroutines) that all read from in.
	c0 := factorial(in)
	c1 := factorial(in)
	c2 := factorial(in)
	c3 := factorial(in)
	c4 := factorial(in)
	c5 := factorial(in)
	c6 := factorial(in)
	c7 := factorial(in)
	c8 := factorial(in)
	c9 := factorial(in)

	// FAN IN
	// multiplex multiple channels onto a single channel
	// merge the channels from c0 through c9 onto a single channel
	var y int
	for n := range merge(c0, c1, c2, c3, c4, c5, c6, c7, c8, c9) {
		y++
		fmt.Println(y, "\t", n)
	}

}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// This is a func expression; it gets assigned to a variable (output)
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// =====================================================

// FAN IN & OUT Explanations and exapmples

// func main() {
// 	in := gen(2, 3)

// 	// FAN OUT
// 	// Distribute the sq work across two goroutines that both read from in.
// 	c1 := sq(in)
// 	c2 := sq(in)

// 	// FAN IN
// 	// Consume the merged output from multiple channels.
// 	for n := range merge(c1, c2) {
// 		fmt.Println(n) // 4 then 9, or 9 then 4
// 	}
// }

// func gen(nums ...int) chan int {
// 	fmt.Printf("TYPE OF NUMS %T\n", nums) // just FYI

// 	out := make(chan int)
// 	go func() {
// 		for _, n := range nums {
// 			out <- n
// 		}
// 		close(out)
// 	}()
// 	return out
// }

// func sq(in chan int) chan int {
// 	out := make(chan int)
// 	go func() {
// 		for n := range in {
// 			out <- n * n
// 		}
// 		close(out)
// 	}()
// 	return out
// }

// func merge(cs ...chan int) chan int {
// 	fmt.Printf("TYPE OF CS: %T\n", cs) // just FYI

// 	out := make(chan int)
// 	var wg sync.WaitGroup
// 	wg.Add(len(cs))

// 	for _, c := range cs {
// 		go func(ch chan int) {
// 			for n := range ch {
// 				out <- n
// 			}
// 			wg.Done()
// 		}(c)
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(out)
// 	}()

// 	return out
// }

/*
FAN OUT
Multiple functions reading from the same channel until that channel is closed

FAN IN
A function can read from multiple inputs and proceed until all are closed by
multiplexing the input channels onto a single channel that's closed when
all the inputs are closed.

PATTERN
there's a pattern to our pipeline functions:
-- stages close their outbound channels when all the send operations are done.
-- stages keep receiving values from inbound channels until those channels are closed.

source:
https://blog.golang.org/pipelines
*/

// =================================================

// func main() {
// 	c := fanIn(boring("Joe"), boring("Ann"))
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(<-c)
// 	}
// 	fmt.Println("You're both boring; I'm leaving.")
// }

// // It is not the FAN OUT. We need channels to come in and distribute
// // accross the functions.
// func boring(msg string) <-chan string {
// 	c := make(chan string)
// 	go func() {
// 		for i := 0; ; i++ {
// 			c <- fmt.Sprintf("%s %d", msg, i)
// 			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
// 		}
// 	}()
// 	return c
// }

// // FAN IN
// func fanIn(input1, input2 <-chan string) <-chan string {
// 	c := make(chan string)
// 	go func() {
// 		for {
// 			// First take the value off from input1 and then load it to c.
// 			c <- <-input1
// 		}
// 	}()
// 	go func() {
// 		for {
// 			c <- <-input2
// 		}
// 	}()
// 	return c
// }
