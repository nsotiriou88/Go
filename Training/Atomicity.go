package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// Getting ready for atomic stuff!
var counter int64
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		// race:
		// counter++
		// no race:
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Counter:", atomic.LoadInt64(&counter)) // access without race
		// it might not print in order, as there is parallelism here, but result
		// would be 40, no matter what.
	}
	wg.Done()
}

// go run -race main.go, checking if we have a race condition
// vs
// go run main.go
