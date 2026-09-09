package main

import (
	"fmt"
	"math/rand"
	"time"
)

type reading struct {
	id    int
	value float64
}

type report struct {
	reading
	offline bool
}

// sensor sends up to maxReadings readings (or forever if maxReadings <= 0),
// simulating a device that can go quiet partway through a run.
func sensor(id int, maxReadings int, out chan<- reading) {
	sent := 0
	for maxReadings <= 0 || sent < maxReadings {
		delay := time.Duration(300+rand.Intn(900)) * time.Millisecond
		time.Sleep(delay)
		out <- reading{id: id, value: 15 + rand.Float64()*15}
		sent++
	}
	// stops sending, but the channel is left open: a real device that has gone
	// quiet doesn't politely close anything on its way out.
}

// monitor watches a single sensor's channel and reports either a fresh reading
// or an "offline" status if nothing arrived within the deadline.
func monitor(id int, in <-chan reading, deadline time.Duration, results chan<- report) {
	for {
		select {
		case r := <-in:
			results <- report{reading: r, offline: false}
		case <-time.After(deadline):
			results <- report{reading: reading{id: id}, offline: true}
		}
	}
}

func main() {
	const numSensors = 4
	const deadline = 2 * time.Second

	reports := make(chan report)

	for i := 1; i <= numSensors; i++ {
		readings := make(chan reading)

		maxReadings := 0
		if i == 3 {
			maxReadings = 3 // sensor 3 goes quiet early, on purpose
		}

		go sensor(i, maxReadings, readings)
		go monitor(i, readings, deadline, reports)
	}

	for r := range reports {
		if r.offline {
			fmt.Printf("sensor %d: OFFLINE\n", r.id)
		} else {
			fmt.Printf("sensor %d: %.2f\n", r.id, r.value)
		}
	}
}
