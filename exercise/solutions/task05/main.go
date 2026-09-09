package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// sensor sends up to maxReadings readings (or forever if maxReadings <= 0),
// exiting early if done is closed.
func sensor(id int, maxReadings int, out chan<- float64, done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	sent := 0
	for maxReadings <= 0 || sent < maxReadings {
		delay := time.Duration(300+rand.Intn(900)) * time.Millisecond

		select {
		case <-time.After(delay):
		case <-done:
			return
		}

		select {
		case out <- 15 + rand.Float64()*15:
			sent++
		case <-done:
			return
		}
	}

	<-done // stays "present" (like a real quiet device) until told to stop
}

// monitor watches a single sensor's channel, printing readings or an offline
// status, and safely tallies successful readings in the shared counts map.
func monitor(
	id int,
	in <-chan float64,
	deadline time.Duration,
	done <-chan struct{},
	wg *sync.WaitGroup,
	mu *sync.Mutex,
	counts map[int]int,
) {
	defer wg.Done()

	for {
		select {
		case r := <-in:
			fmt.Printf("sensor %d: %.2f\n", id, r)
			mu.Lock()
			counts[id]++
			mu.Unlock()
		case <-time.After(deadline):
			fmt.Printf("sensor %d: OFFLINE\n", id)
		case <-done:
			return
		}
	}
}

func main() {
	const numSensors = 4
	const deadline = 2 * time.Second
	const runFor = 15 * time.Second

	done := make(chan struct{})
	var wg sync.WaitGroup
	var mu sync.Mutex
	counts := make(map[int]int)

	for i := 1; i <= numSensors; i++ {
		readings := make(chan float64)

		maxReadings := 0
		if i == 3 {
			maxReadings = 3
		}

		wg.Add(2) // one sensor goroutine + one monitor goroutine
		go sensor(i, maxReadings, readings, done, &wg)
		go monitor(i, readings, deadline, done, &wg, &mu, counts)
	}

	time.Sleep(runFor)
	fmt.Println("shutting down...")
	close(done)
	wg.Wait() // wg.Wait() already synchronizes: no lock needed for this final read

	for i := 1; i <= numSensors; i++ {
		fmt.Printf("sensor %d: %d readings\n", i, counts[i])
	}
}
