package main

import (
	"fmt"
	"time"
)

func main() {

	timer1 := time.NewTimer(2 * time.Second) // (1)!

	<-timer1.C // (2)!
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second) // (3)!
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second) // (4)!
}
