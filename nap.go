package main

import "time"

// nap takes a nap for n milliseconds.
// TODO: Eliminate this.
func nap(x int) {
	if x <= 0 {
		return
	}
	lflush()
	time.Sleep(time.Duration(x) * time.Millisecond)
}
