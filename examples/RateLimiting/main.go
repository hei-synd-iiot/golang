package main

import (
	"fmt"
	"time"
)

func main() {

	requests := make(chan int, 5) // (1)!
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond) // (2)!

	for req := range requests { // (3)!
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3) // (4)!

	for i := 0; i < 3; i++ { // (5)!
		burstyLimiter <- time.Now()
	}

	go func() { // (6)!
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ { // (7)!
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
