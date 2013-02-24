package main

import (
	"bytes"
	"io/ioutil"
)

/*
 *	help function to display the help info
 *
 *	format of the larn.help file
 *
 *	1st character of file:	# of pages of help available (ascii digit)
 *	page (23 lines) for the introductory message (not counted in above)
 *	pages of help text (23 lines per page)
 */
func help() {
	j, lines := loadhelp()
	if j < 0 {
		return /* open the help file and get # pages */
	}
	lines = lines[23:] // skip over intro message
	for ; j > 0; j-- {
		clear()
		for _, line := range lines[:23] {
			lprcat(line + "\n")
		}
		lines = lines[23:]
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
				setscroll()
				drawscreen()
				return
			}
		}
	}
	retcont()
	drawscreen()
}

/*
 *	function to display the welcome message and background
 */
func welcome() {
	j, lines := loadhelp()
	if j < 0 {
		return
	}
	clear()
	for _, line := range lines[:23] {
		lprcat(line + "\n")
	}
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

// loadhelp loads the help file, and returns the number of pages
// and all the lines of the file.
func loadhelp() (int, []string) {
	raw, err := ioutil.ReadFile(helpfile)
	if err != nil {
		lprintf("Can't open help file %q: %v", helpfile, err)
		lflush()
		nap(4000)
		drawscreen()
		setscroll()
		return -1, nil
	}
	resetscroll()
	pages := int(raw[0] - '0')
	raw = raw[1:]
	var lines []string
	for _, l := range bytes.Split(raw, []byte("\n")) {
		lines = append(lines, string(l))
	}
	return pages, lines
}
