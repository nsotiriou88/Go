package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// mutual exclusive. Needs to be global!
var mutex sync.Mutex
var wg sync.WaitGroup
var counter int

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		// We lock the counter and only when it exits, the next one processes
		mutex.Lock()
		// since we access once per cycle, the counter is safe like this:
		counter++
		fmt.Println(s, i, "Counter:", counter)
		mutex.Unlock()
	}
	wg.Done()
}

// go run -race main.go, checking if we have a race condition
// vs
// go run main.go
