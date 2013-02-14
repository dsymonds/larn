package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

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
func cntlc() {
	debugf("()")
	if nosignal {
		return /* don't do anything if inhibited */
	}
	// TODO
	//signal(SIGQUIT, SIG_IGN);
	//signal(SIGINT, SIG_IGN);
	quit()
	if predostuff == 1 {
		s2choose()
	} else {
		showplayer()
	}
	lflush()
	//signal(SIGQUIT, cntlc);
	//signal(SIGINT, cntlc);
}

/*
 *	subroutine to save the game if a hangup signal
 */
func sgam(_ int) {
	savegame(savefilename)
	wizard = true
	died(-257) /* hangup signal */
}

// tstop handles ^Y.
func tstop(n int) {
	if nosignal {
		return /* nothing if inhibited */
	}
	// TODO
	/*
			lcreat("");
			clearvt100();
			lflush();
			signal(SIGTSTP, SIG_DFL);
		#ifdef SIGVTALRM
	*/
	/*
	 * looks like BSD4.2 or higher - must clr mask for signal to take
	 * effect
	 */
	/*
			sigsetmask(sigblock(0) & ~bit(SIGTSTP));
		#endif
			kill(getpid(), SIGTSTP);

			setupvt100();
			signal(SIGTSTP, tstop);
			if (predostuff == 1)
				s2choose();
			else
				drawscreen();
			showplayer();
			lflush();
	*/
}

/*
 *	subroutine to issue the needed signal traps  called from main()
 */
func sigsetup() {
	debugf("()")
	c := make(chan os.Signal, 16)
	signal.Notify(c, syscall.SIGQUIT, syscall.SIGINT)
	go func() {
		for sig := range c {
			switch sig {
			case syscall.SIGQUIT, syscall.SIGINT:
				cntlc()
			}
		}
	}()

	// TODO
	/*
		signal(SIGQUIT, cntlc);
		signal(SIGINT, cntlc);
		signal(SIGKILL, SIG_IGN);
		signal(SIGHUP, sgam);
		signal(SIGILL, sigpanic);
		signal(SIGTRAP, sigpanic);
		signal(SIGIOT, sigpanic);
		signal(SIGEMT, sigpanic);
		signal(SIGFPE, sigpanic);
		signal(SIGBUS, sigpanic);
		signal(SIGSEGV, sigpanic);
		signal(SIGSYS, sigpanic);
		signal(SIGPIPE, sigpanic);
		signal(SIGTERM, sigpanic);
		signal(SIGTSTP, tstop);
		signal(SIGSTOP, tstop);
	*/
}

/*
 *	routine to process a fatal error signal
 */
func sigpanic(sig int) {
	// TODO
	//signal(sig, SIG_DFL);
	fmt.Fprintf(os.Stderr, "\nLarn - Panic! Signal %d received [%v]", sig, syscall.Signal(sig))
	time.Sleep(2 * time.Second)
	sncbr()
	savegame(savefilename)
	// TODO
	//kill(getpid(), sig);	/* this will terminate us */
	os.Exit(1)
}
