package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var ops atomic.Uint64 // (1)!

	var wg sync.WaitGroup // (2)!

	for i := 0; i < 50; i++ { // (3)!
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {

				ops.Add(1) // (4)!
			}

			wg.Done()
		}()
	}

	wg.Wait() // (5)!

	fmt.Println("ops:", ops.Load()) // (6)!
}
