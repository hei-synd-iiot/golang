package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) { // (1)!
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second) // (2)!
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup // (3)!

	for i := 1; i <= 5; i++ { // (4)!
		wg.Add(1)

		go func() { // (5)!
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait() // (6)!
}
