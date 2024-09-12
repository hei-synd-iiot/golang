package main

import (
	"fmt"
	"sync"
)

type Container struct { // (1)!
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) { // (2)!
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() { // (3)!
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) { // (4)!
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3) // (5)!
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait() // (6)!
	fmt.Println(c.counters)
}
