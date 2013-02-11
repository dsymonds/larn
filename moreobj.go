package main

/*
 * Routines in this file:
 *
 * oaltar() othrone() ochest() ofountain()
 */

/*
 *	subroutine to process an altar object
 */
func oaltar() {

	lprcat("\nDo you (p) pray  (d) desecrate")
	iopts()
	for {
		switch ttgetch() {
		case 'p':
			lprcat(" pray\nDo you (m) give money or (j) just pray? ")
			for {
				switch ttgetch() {
				case 'j':
					act_just_pray()
					return

				case 'm':
					act_donation_pray()
					return

				case '\033':
					return
				}
			}

		case 'd':
			lprcat(" desecrate")
			act_desecrate_altar()
			return

		case 'i', '\033':
			ignore()
			act_ignore_altar()
			return
		}
	}
}

/*
	subroutine to process a throne object
*/
func othrone(arg int) {

	lprcat("\nDo you (p) pry off jewels, (s) sit down")
	iopts()
	for {
		switch ttgetch() {
		case 'p':
			lprcat(" pry off")
			act_remove_gems(arg)
			return

		case 's':
			lprcat(" sit down")
			act_sit_throne(arg)
			return

		case 'i', '\033':
			ignore()
			return
		}
	}
}

func odeadthrone() {
	lprcat("\nDo you (s) sit down")
	iopts()
	for {
		switch ttgetch() {
		case 's':
			lprcat(" sit down")
			k := rnd(101)
			if k < 35 {
				lprcat("\nZaaaappp!  You've been teleported!\n")
				beep()
				oteleport(0)
			} else {
				lprcat("\nnothing happens")
			}
			return

		case 'i', '\033':
			ignore()
			return
		}
	}
}

/*
	subroutine to process a throne object
*/
func ochest() {

	lprcat("\nDo you (t) take it, (o) try to open it")
	iopts()
	for {
		switch ttgetch() {
		case 'o':
			lprcat(" open it")
			act_open_chest(playerx, playery)
			return

		case 't':
			lprcat(" take")
			if take(OCHEST, iarg[playerx][playery]) == 0 {
				item[playerx][playery], know[playerx][playery] = 0, false
			}
			return

		case 'i', '\033':
			ignore()
			return
		}
	}
}

/*
	process a fountain object
*/
func ofountain() {

	cursors()
	lprcat("\nDo you (d) drink, (w) wash yourself")
	iopts()
	for {
		switch ttgetch() {
		case 'd':
			lprcat("drink")
			act_drink_fountain()
			return

		case '\033', 'i':
			ignore()
			return

		case 'w':
			lprcat("wash yourself")
			act_wash_fountain()
			return
		}
	}
}

/*
	***
	FCH
	***

	subroutine to process an up/down of a character attribute for ofountain
*/
func fch(how int, x *int32) {
	if how < 0 {
		lprcat(" went down by one!")
		(*x)--
	} else {
		lprcat(" went up by one!")
		(*x)++
	}
	bottomline()
}

/*
	a subroutine to raise or lower character levels
	if x > 0 they are raised   if x < 0 they are lowered
*/
func fntchange(how int) {
	lprc('\n')
	switch rnd(9) {
	case 1:
		lprcat("Your strength")
		fch(how, &c[0])
	case 2:
		lprcat("Your intelligence")
		fch(how, &c[1])
	case 3:
		lprcat("Your wisdom")
		fch(how, &c[2])
	case 4:
		lprcat("Your constitution")
		fch(how, &c[3])
	case 5:
		lprcat("Your dexterity")
		fch(how, &c[4])
	case 6:
		lprcat("Your charm")
		fch(how, &c[5])
	case 7:
		j := rnd(level + 1)
		if how < 0 {
			lprintf("You lose %d hit point", j)
			if j > 1 {
				lprcat("s!")
			} else {
				lprc('!')
			}
			losemhp(j)
		} else {
			lprintf("You gain %d hit point", j)
			if j > 1 {
				lprcat("s!")
			} else {
				lprc('!')
			}
			raisemhp(j)
		}
		bottomline()

	case 8:
		j := rnd(level + 1)
		if how > 0 {
			lprintf("You just gained %d spell", j)
			raisemspells(j)
			if j > 1 {
				lprcat("s!")
			} else {
				lprc('!')
			}
		} else {
			lprintf("You just lost %d spell", j)
			losemspells(j)
			if j > 1 {
				lprcat("s!")
			} else {
				lprc('!')
			}
		}
		bottomline()

	case 9:
		j := 5 * rnd((level+1)*(level+1))
		if how < 0 {
			lprintf("You just lost %d experience point", j)
			if j > 1 {
				lprcat("s!")
			} else {
				lprc('!')
			}
			loseexperience(j)
		} else {
			lprintf("You just gained %d experience point", j)
			if j > 1 {
				lprcat("s!")
			} else {
				lprc('!')
			}
			raiseexperience(j)
		}
	}
	cursors()
}
