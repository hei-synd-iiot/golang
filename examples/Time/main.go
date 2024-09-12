package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now() // (1)!
	p(now)

	then := time.Date( // (2)!
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	p(then.Year()) // (3)!
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday()) // (4)!

	p(then.Before(now)) // (5)!
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then) // (6)!
	p(diff)

	p(diff.Hours()) // (7)!
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff)) // (8)!
	p(then.Add(-diff))

	p(now.Unix()) // (9)!
	p(now.UnixMilli())
	p(now.UnixNano())

	p(time.Unix(now.Unix(), 0)) // (10)!
	p(time.Unix(0, now.UnixNano()))
}
