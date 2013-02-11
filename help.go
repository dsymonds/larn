package main

import (
	"time"
)

/*
 *	help function to display the help info
 *
 *	format of the .larn.help file
 *
 *	1st character of file:	# of pages of help available (ascii digit)
 *	page (23 lines) for the introductory message (not counted in above)
 *	pages of help text (23 lines per page)
 */
func help() {
	j := openhelp()
	if j < 0 {
		return /* open the help file and get # pages */
	}
	for i := 0; i < 23; i++ {
		lgetl() /* skip over intro message */
	}
	for ; j > 0; j-- {
		clear()
		for i := 0; i < 23; i++ {
			lprcat(lgetl()) /* print out each line that we read in */
		}
		if j > 1 {
			lprcat("    ---- Press ")
			standout("return")
			lprcat(" to exit, ")
			standout("space")
			lprcat(" for more help ---- ")
			i := 0
			for i != ' ' && i != '\n' && i != '\033' {
				i = ttgetch()
			}
			if i == '\n' || i == '\033' {
				lrclose()
				setscroll()
				drawscreen()
				return
			}
		}
	}
	lrclose()
	retcont()
	drawscreen()
}

/*
 *	function to display the welcome message and background
 */
func welcome() {
	if openhelp() < 0 {
		return /* open the help file */
	}
	clear()
	for i := 0; i < 23; i++ {
		lprcat(lgetl()) /* print out each line that we read in */
	}
	lrclose()
	retcont() /* press return to continue */
}

/*
 *	function to say press return to continue and reset scroll when done
 */
func retcont() {
	cursor(1, 24)
	lprcat("Press ")
	standout("return")
	lprcat(" to continue: ")
	for ttgetch() != '\n' {
	}
	setscroll()
}

/*
 *	routine to open the help file and return the first character - '0'
 */
func openhelp() int {
	if lopen(helpfile) < 0 {
		lprintf("Can't open help file \"%s\" ", helpfile)
		lflush()
		time.Sleep(4 * time.Second)
		drawscreen()
		setscroll()
		return -1
	}
	resetscroll()
	return lgetc() - '0'
}
