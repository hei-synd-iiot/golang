package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) { // (1)!
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true // (2)!
}

func main() {

	done := make(chan bool, 1)
	go worker(done) // (3)!

	<-done // (4)!
}
