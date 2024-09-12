package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) { // (1)!
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5 // (2)!
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ { // (3)!
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ { // (4)!
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ { // (5)!
		<-results
	}
}
