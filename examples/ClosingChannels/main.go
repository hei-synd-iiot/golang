package main

import "fmt"

func main() { // (1)!
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() { // (2)!
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ { // (3)!
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done // (4)!

	_, ok := <-jobs // (5)!
	fmt.Println("received more jobs:", ok)
}
