package main

import (
	"time"
)

/*
 * act_remove_gems
 *
 * Remove gems from a throne.
 *
 * arg is zero if there is a gnome king associated with the throne.
 *
 * Assumes that cursors() has been called previously, and that a check
 * has been made that the throne actually has gems.
 */
func act_remove_gems(arg int) {
	k := rnd(101)
	if k < 25 {
		for i := 0; i < rnd(4); i++ {
			creategem() /* gems pop off the throne */
		}
		item[playerx][playery] = ODEADTHRONE
		know[playerx][playery] = false
	} else if k < 40 && arg == 0 {
		createmonster(GNOMEKING)
		item[playerx][playery] = OTHRONE2
		know[playerx][playery] = false
	} else {
		lprcat("\nnothing happens")
	}
}

/*
 * act_sit_throne
 *
 * Sit on a throne.
 *
 * arg is zero if there is a gnome king associated with the throne
 *
 * Assumes that cursors() has been called previously.
 */
func act_sit_throne(arg int) {
	k := rnd(101)
	if k < 30 && arg == 0 {
		createmonster(GNOMEKING)
		item[playerx][playery] = OTHRONE2
		know[playerx][playery] = false
	} else if k < 35 {
		lprcat("\nZaaaappp!  You've been teleported!\n")
		beep()
		oteleport(0)
	} else {
		lprcat("\nnothing happens")
	}
}

/*
 * Code to perform the action of drinking at a fountain.  Assumes that
 * cursors() has already been called, and that a check has been made
 * that the player is actually standing at a live fountain.
 */
func act_drink_fountain() {
	if rnd(1501) < 2 {
		lprcat("\nOops!  You seem to have caught the dreadful sleep!")
		beep()
		lflush()
		time.Sleep(3 * time.Second)
		died(280)
		return
	}
	x := rnd(100)
	if x < 7 {
		c[HALFDAM] += 200 + rnd(200)
		lprcat("\nYou feel a sickness coming on")
	} else if x < 13 {
		quaffpotion(23) /* see invisible */
	} else if x < 45 {
		lprcat("\nnothing seems to have happened")
	} else if rnd(3) != 2 {
		fntchange(1) /* change char levels upward	 */
	} else {
		fntchange(-1) /* change char levels downward	 */
	}
	if rnd(12) < 3 {
		lprcat("\nThe fountains bubbling slowly quiets")
		item[playerx][playery] = ODEADFOUNTAIN /* dead fountain */
		know[playerx][playery] = false
	}
}

/*
 * Code to perform the action of washing at a fountain.  Assumes that
 * cursors() has already been called and that a check has been made
 * that the player is actually standing at a live fountain.
 */
func act_wash_fountain() {
	if rnd(100) < 11 {
		x := rnd((level << 2) + 2)
		lprintf("\nOh no!  The water was foul!  You suffer %d hit points!", x)
		lastnum = 273
		losehp(x)
		bottomline()
		cursors()
	} else if rnd(100) < 29 {
		lprcat("\nYou got the dirt off!")
	} else if rnd(100) < 31 {
		lprcat("\nThis water seems to be hard water!  The dirt didn't come off!")
	} else if rnd(100) < 34 {
		createmonster(WATERLORD) /* make water lord		 */
	} else {
		lprcat("\nnothing seems to have happened")
	}
}

/*
 * Perform the actions associated with altar desecration.
 */
func act_desecrate_altar() {
	if rnd(100) < 60 {
		createmonster(makemonst(level+2) + 8)
		c[AGGRAVATE] += 2500
	} else if rnd(101) < 30 {
		lprcat("\nThe altar crumbles into a pile of dust before your eyes")
		forget() /* remember to destroy the altar	 */
	} else {
		lprcat("\nnothing happens")
	}
}

/*
 * Perform the actions associated with praying at an altar and giving
 * a donation.
 */
func act_donation_pray() {
	for {
		lprcat("\n\n")
		cursor(1, 24)
		cltoeoln()
		cursor(1, 23)
		cltoeoln()
		lprcat("how much do you donate? ")
		amt := readnum(c[GOLD])
		if amt < 0 || c[GOLD] < amt {
			// Loop around to prevent the player from escaping the altar for free.
			lprcat("\nYou don't have that much!")
			continue
		}
		min := c[GOLD] / 10
		c[GOLD] -= amt
		bottomline()
		if amt < min || amt < rnd(50) {
			// Player gave less than 10% of their gold; summon a monster.
			createmonster(makemonst(level + 1))
			c[AGGRAVATE] += 200
		} else if rnd(101) > 50 {
			ohear()
		} else if rnd(43) == 5 {
			if c[WEAR] != 0 {
				lprcat("\nYou feel your armor vibrate for a moment")
			}
			enchantarmor()
		} else if rnd(43) == 8 {
			if c[WIELD] != 0 {
				lprcat("\nYou feel your weapon vibrate for a moment")
			}
			enchweapon()
		} else {
			lprcat("\nThank You.")
		}
		return
	}
}

/*
 *  Performs the actions associated with 'just praying' at the altar.  Called
 *  when the user responds 'just pray' when in prompt mode, or enters 0 to
 *  the money prompt when praying.
 *
 *  Assumes cursors(), and that any leading \n have been printed.
 */
func act_just_pray() {
	if rnd(100) < 75 {
		lprcat("\nnothing happens")
	} else if rnd(13) < 4 {
		ohear()
	} else if rnd(43) == 10 {
		if c[WEAR] != 0 {
			lprcat("\nYou feel your armor vibrate for a moment")
		}
		enchantarmor()
		return
	} else if rnd(43) == 10 {
		if c[WIELD] != 0 {
			lprcat("\nYou feel your weapon vibrate for a moment")
		}
		enchweapon()
		return
	} else {
		createmonster(makemonst(level + 1))
	}
}

/*
 * Function to cast a +3 protection on the player
 */
func ohear() {
	lprcat("\nYou have been heard!")
	if c[ALTPRO] == 0 {
		c[MOREDEFENSES] += 3
	}
	c[ALTPRO] += 500 /* protection field */
	bottomline()
}

/*
 * Performs the act of ignoring an altar.
 *
 * Assumptions:  cursors() has been called.
 */
func act_ignore_altar() {
	if rnd(100) < 30 {
		createmonster(makemonst(level + 1))
		c[AGGRAVATE] += rnd(450)
	} else {
		lprcat("\nnothing happens")
	}
}

/*
 * Performs the act of opening a chest.
 *
 * Parameters:   x,y location of the chest to open.
 * Assumptions:  cursors() has been called previously
 */
func act_open_chest(x, y int) {
	k := rnd(101)
	if k < 40 {
		lprcat("\nThe chest explodes as you open it")
		beep()
		i := rnd(10)
		lastnum = 281 /* in case he dies */
		lprintf("\nYou suffer %d hit points damage!", i)
		checkloss(i)
		switch rnd(10) { /* see if he gets a curse */
		case 1:
			c[ITCHING] += rnd(1000) + 100
			lprcat("\nYou feel an irritation spread over your skin!")
			beep()

		case 2:
			c[CLUMSINESS] += rnd(1600) + 200
			lprcat("\nYou begin to lose hand to eye coordination!")
			beep()

		case 3:
			c[HALFDAM] += rnd(1600) + 200
			beep()
			lprcat("\nA sickness engulfs you!")
		}
		item[x][y], know[x][y] = 0, false
		if rnd(100) < 69 {
			creategem() /* gems from the chest */
		}
		dropgold(rnd(110*iarg[x][y] + 200))
		for i := 0; i < rnd(4); i++ {
			something(iarg[x][y] + 2)
		}
	} else {
		lprcat("\nnothing happens")
	}
}
