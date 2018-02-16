package main

import "fmt"

func main() {
	// // Set up the pipeline and consume the output.
	// for n := range sq(sq(gen(2, 3))) {
	// 	fmt.Println(n) // 16 then 81
	// }

	// // Set up the pipeline.
	// c := gen(2, 3)
	// out := sq(c)

	// // Consume the output.
	// fmt.Println(<-out) // 4
	// fmt.Println(<-out) // 9

	// Set up the pipeline and consume the output.
	for n := range sq(gen(2, 3)) {
		fmt.Println(n) // 4 then 9
	}

	// Main part of Factorial
	in := gen2()

	f := factorial(in)

	for n := range f {
		fmt.Println(n)
	}
}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// ========  FACTORIAL  ========

// The Main() part is moved up.

func gen2() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
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
