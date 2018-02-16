package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

// "init" is a special function that allows you to do some
// initial setup. Here we set all cores available to usage.
// After Go 1.5, using all cores is by default.
// You can have many "init" and it executes first and once.
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// Concurrency
	// Adding go, is creating concurrencies. Including the
	// main, we have 3 threads running (plus foo & bar).
	wg.Add(2)
	go foo()
	println()
	go bar()
	wg.Wait()

	// Parallelism

}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(time.Duration(3 * time.Millisecond)) // try without
	}
	// adding waitgroup for checking functionality. Without
	// wg, the program will execute and there will be no output,
	// as we see only the first thread (main func).
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar:", i)
		// adding some delay to see the execution of the program.
		// Also, it can cause deviation in pace of each function,
		// if we set different sleeping time.
		time.Sleep(time.Duration(10 * time.Millisecond)) // try without
	}
	wg.Done()
}

// Concurrency is different from parallelism.
//
// Concurrency is about dealing with lots of things at once, while
// Parallelism is about doing lots of things simultaneously at once.
// This means that one processor is doing everything in Concurrencies,
// a little bit of each function.
// Parallelism is/has Concurrency, but Concurrency has no Parallelism.
