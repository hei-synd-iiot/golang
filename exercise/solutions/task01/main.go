package main

import (
	"fmt"
	"math/rand"
	"time"
)

// sensor repeatedly sends a random reading on out after a random delay,
// simulating a slow, irregularly-reporting field sensor.
func sensor(id int, out chan<- float64) {
	for {
		delay := time.Duration(300+rand.Intn(900)) * time.Millisecond
		time.Sleep(delay)
		out <- 15 + rand.Float64()*15
	}
}

func main() {
	readings := make(chan float64)
	go sensor(1, readings)

	for r := range readings {
		fmt.Printf("sensor 1: %.2f\n", r)
	}
}
