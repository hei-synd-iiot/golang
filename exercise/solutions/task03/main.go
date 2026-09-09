package main

import (
	"fmt"
	"math/rand"
	"time"
)

type report struct {
	id      int
	reading float64
	offline bool
}

// sensor sends up to maxReadings readings (or forever if maxReadings <= 0),
// simulating a device that can go quiet partway through a run.
func sensor(id int, maxReadings int, out chan<- float64) {
	sent := 0
	for maxReadings <= 0 || sent < maxReadings {
		delay := time.Duration(300+rand.Intn(900)) * time.Millisecond
		time.Sleep(delay)
		out <- 15 + rand.Float64()*15
		sent++
	}
	// stops sending, but the channel is left open: a real device that has gone
	// quiet doesn't politely close anything on its way out.
}

// monitor watches a single sensor's channel and reports either a fresh reading
// or an "offline" status if nothing arrived within the deadline.
func monitor(id int, in <-chan float64, deadline time.Duration, results chan<- report) {
	for {
		select {
		case r := <-in:
			results <- report{id: id, reading: r}
		case <-time.After(deadline):
			results <- report{id: id, offline: true}
		}
	}
}

func main() {
	const numSensors = 4
	const deadline = 2 * time.Second

	results := make(chan report)

	for i := 1; i <= numSensors; i++ {
		readings := make(chan float64)

		maxReadings := 0
		if i == 3 {
			maxReadings = 3 // sensor 3 goes quiet early, on purpose
		}

		go sensor(i, maxReadings, readings)
		go monitor(i, readings, deadline, results)
	}

	for r := range results {
		if r.offline {
			fmt.Printf("sensor %d: OFFLINE\n", r.id)
		} else {
			fmt.Printf("sensor %d: %.2f\n", r.id, r.reading)
		}
	}
}
