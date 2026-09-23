package main

import (
	"fmt"
	"math/rand"
	"sync"
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
// exiting early if done is closed.
func sensor(id int, maxReadings int, out chan<- reading, done <-chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	sent := 0
	for maxReadings <= 0 || sent < maxReadings {
		delay := time.Duration(300+rand.Intn(900)) * time.Millisecond

		select {
		case <-time.After(delay):
			out <- reading{id: id, value: 15 + rand.Float64()*15}
			sent++
		case <-done:
			return
		}
	}

	<-done // stays "present" (like a real quiet device) until told to stop
}

// monitor watches a single sensor's channel, safely tallying successful
// readings in the shared counts map and forwarding a report of each reading
// (or offline status) to the shared results channel for display.
func monitor(
	id int,
	in <-chan reading,
	deadline time.Duration,
	done <-chan struct{},
	wg *sync.WaitGroup,
	mu *sync.Mutex,
	counts map[int]int,
	results chan<- report,
) {
	wg.Add(1)
	defer wg.Done()

	for {
		select {
		case r := <-in:
			mu.Lock()
			counts[id]++
			mu.Unlock()
			results <- report{reading: r, offline: false}
		case <-time.After(deadline):
			results <- report{reading: reading{id: id}, offline: true}
		case <-done:
			return
		}
	}
}

// reporter prints out the reports from the sensor monitor.
// It exist when the report channel is closed and consequently closes the done channel.
func reporter(reports chan report, done chan struct{}) {
	for r := range reports {
		if r.offline {
			fmt.Printf("sensor %d: OFFLINE\n", r.id)
		} else {
			fmt.Printf("sensor %d: %.2f\n", r.id, r.value)
		}
	}
	close(done)
}

func main() {
	const numSensors = 4
	const deadline = 2 * time.Second
	const runFor = 15 * time.Second

	done := make(chan struct{})
	var wg sync.WaitGroup
	var mu sync.Mutex
	counts := make(map[int]int)
	reports := make(chan report)

	for i := 1; i <= numSensors; i++ {
		readings := make(chan reading)

		maxReadings := 0
		if i == 3 {
			maxReadings = 3
		}

		go sensor(i, maxReadings, readings, done, &wg)
		go monitor(i, readings, deadline, done, &wg, &mu, counts, reports)
	}

	reportsDone := make(chan struct{})
	go reporter(reports, reportsDone)

	time.Sleep(runFor)
	fmt.Println("shutting down...")

	close(done)    // closing the done channel forces all sensors and monitors to exit
	wg.Wait()      // wait for all sensors and monitors to be stopped
	close(reports) // closing the reports channel drains the channel and exits the reporting goroutine
	<-reportsDone  // wait for last report to be printed before the summary

	for i := 1; i <= numSensors; i++ {
		fmt.Printf("sensor %d: %d readings\n", i, counts[i])
	}
}
