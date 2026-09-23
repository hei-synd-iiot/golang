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
	const numSensors = 4

	channels := make([]chan float64, numSensors)
	for i := range channels {
		channels[i] = make(chan float64)
		go sensor(i+1, channels[i])
	}

	for {
		select {
		case r := <-channels[0]:
			fmt.Printf("sensor 1: %.2f\n", r)
		case r := <-channels[1]:
			fmt.Printf("sensor 2: %.2f\n", r)
		case r := <-channels[2]:
			fmt.Printf("sensor 3: %.2f\n", r)
		case r := <-channels[3]:
			fmt.Printf("sensor 4: %.2f\n", r)
		}
	}
}
