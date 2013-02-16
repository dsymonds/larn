package main

import "time"

// nap takes a nap for n milliseconds.
func nap(x int) {
	if x <= 0 {
		return
	}
	lflush()
	if len(replayActions) > 0 {
		return
	}
	time.Sleep(time.Duration(x) * time.Millisecond)
}
