package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

// Simulation of signal(3) and friends.
var sigsim = struct {
	sync.Mutex
	handlers map[os.Signal]func(os.Signal)

	c chan os.Signal
}{
	handlers: make(map[os.Signal]func(os.Signal)),
	c:        make(chan os.Signal, 20),
}

func sigwatcher() {
	for sig := range sigsim.c {
		debugf("got signal %v", sig)
		sigsim.Lock()
		h := sigsim.handlers[sig]
		sigsim.Unlock()
		if h != nil {
			h(sig)
		}
	}
}

func simsignal(sig os.Signal, handler func(os.Signal)) {
	sigsim.Lock()
	if _, ok := sigsim.handlers[sig]; !ok {
		signal.Notify(sigsim.c, sig)
	}
	sigsim.handlers[sig] = handler
	sigsim.Unlock()
}

func sigignore(os.Signal) {}

func bit(a uint) int { return 1 << (a - 1) }

func s2choose() {
	/* text to be displayed if ^C during intro screen */
	cursor(1, 24)
	lprcat("Press ")
	setbold()
	lprcat("return")
	resetbold()
	lprcat(" to continue: ")
	lflush()
}

// cntlc handles a ^C.
func cntlc(_ os.Signal) {
	debugf("()")
	if nosignal {
		return /* don't do anything if inhibited */
	}
	simsignal(os.Interrupt, sigignore)
	quit()
	if predostuff == 1 {
		s2choose()
	} else {
		showplayer()
	}
	lflush()
	simsignal(os.Interrupt, cntlc)
}

/*
 *	subroutine to issue the needed signal traps  called from main()
 */
func sigsetup() {
	debugf("()")
	go sigwatcher()
	simsignal(os.Interrupt, cntlc)
}

/*
 *	routine to process a fatal error signal
 */
func sigpanic(sig os.Signal) {
	fmt.Fprintf(os.Stderr, "\nLarn - Panic! Signal %d received [%v]", sig, sig)
	time.Sleep(2 * time.Second)
	sncbr()
	savegame(savefilename)
	os.Exit(1)
}
