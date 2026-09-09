package main

import (
	"fmt"
	"math/rand"
	"time"
)

// reading is a single sensor's report: which sensor it came from, and what it read.
type reading struct {
	id    int
	value float64
}

// sensor repeatedly sends a random reading on out after a random delay,
// simulating a slow, irregularly-reporting field sensor.
func sensor(id int, out chan<- reading) {
	for {
		delay := time.Duration(300+rand.Intn(900)) * time.Millisecond
		time.Sleep(delay)
		out <- reading{id: id, value: 15 + rand.Float64()*15}
	}
}

func main() {
	const numSensors = 4

	readings := make(chan reading)
	for i := 1; i <= numSensors; i++ {
		go sensor(i, readings)
	}

	for r := range readings {
		fmt.Printf("sensor %d: %.2f\n", r.id, r.value)
	}
}
