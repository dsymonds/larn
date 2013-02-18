package main

import (
	"fmt"
)

/*
 * This file contains the following functions:
 * ----------------------------------------------------------------------------
 *
 * createmonster(monstno) 	Function to create a monster next to the player
 * int monstno;
 *
 * int cgood(x,y,itm,monst)Function to check location for emptiness
 * int x,y,itm,monst;
 *
 * createitem(it,arg) 		Routine to place an item next to the player
 * int it,arg;
 *
 * cast() 			Subroutine called by parse to cast a spell for the user
 *
 * speldamage(x) 	Function to perform spell functions cast by the player
 * int x;
 *
 * loseint()		Routine to decrement your int (intelligence) if > 3
 *
 * isconfuse() 	Routine to check to see if player is confused
 *
 * nospell(x,monst)Routine to return 1 if a spell doesn't affect a monster
 * int x,monst;
 *
 * fullhit(xx)		Function to return full damage against a monst (aka web)
 * int xx;
 *
 * direct(spnum,dam,str,arg)Routine to direct spell damage 1 square in 1 dir
 * int spnum,dam,arg;
 * char *str;
 *
 * godirect(spnum,dam,str,delay,cshow)	Function to perform missile attacks
 * int spnum,dam,delay;
 * char *str,cshow;
 *
 * ifblind(x,y)Routine to put "monster" or the monster name into lastmosnt
 * int x,y;
 *
 * tdirect(spnum)		Routine to teleport away a monster
 * int spnum;
 *
 * omnidirect(sp,dam,str)  Routine to damage all monsters 1 square from player
 * int sp,dam;
 * char *str;
 *
 * dirsub(x,y)		Routine to ask for direction, then modify x,y for it
 * int *x,*y;
 *
 * vxy(x,y)	  	Routine to verify/fix (*x,*y) for being within bounds
 * int *x,*y;
 *
 * dirpoly(spnum)	Routine to ask for a direction and polymorph a monst
 * int spnum;
 *
 * hitmonster(x,y) Function to hit a monster at the designated coordinates
 * int x,y;
 *
 * hitm(x,y,amt)	Function to just hit a monster at a given coordinates
 * int x,y,amt;
 *
 * hitplayer(x,y) 	Function for the monster to hit the player from (x,y)
 * int x,y;
 *
 * dropsomething(monst) Function to create an object when a monster dies
 * int monst;
 *
 * dropgold(amount) 	Function to drop some gold around player
 * int amount;
 *
 * something(level) 	Function to create a random item around player
 * int level;
 *
 * newobject(lev,i) 	Routine to return a randomly selected new object
 * int lev,*i;
 *
 *  spattack(atckno,xx,yy)  Function to process special attacks from monsters
 *   int atckno,xx,yy;
 *
 * checkloss(x) Routine to subtract hp from user and flag bottomline display
 * int x;
 *
 * annihilate()   Routine to annihilate monsters around player, playerx,playery
 *
 * newsphere(x,y,dir,lifetime)  Function to create a new sphere of annihilation
 * int x,y,dir,lifetime;
 *
 * rmsphere(x,y)	Function to delete a sphere of annihilation from list
 * int x,y;
 *
 * sphboom(x,y)	Function to perform the effects of a sphere detonation
 * int x,y;
 *
 * genmonst()		Function to ask for monster and genocide from game
 *
 */

type isave struct { /* used for altar reality */
	typ bool /* false=item,  true=monster */
	id  int  /* item number or monster number */
	arg int  /* the type of item or hitpoints of monster */
}

/*
 * createmonster(monstno)	Function to create a monster next to the player
 * 	int monstno;
 *
 * Enter with the monster number (1 to MAXMONST+8)
 * Returns no value.
 */
func createmonster(mon int) {
	if mon < 1 || mon > MAXMONST+8 { /* check for monster number out of bounds */
		beep()
		lprintf("\ncan't createmonster(%d)\n", mon)
		nap(3000)
		return
	}
	for monster[mon].genocided && mon < MAXMONST {
		mon++ /* genocided? */
	}
	for k, i := rnd(8), -8; i < 0; i++ { /* choose direction, then try all */
		if k > 8 {
			k = 1 /* wraparound the diroff arrays */
		}
		x := playerx + diroffx[k]
		y := playery + diroffy[k]
		if cgood(x, y, 0, 1) { /* if we can create here */
			mitem[x][y] = mon
			hitp[x][y] = monster[mon].hitpoints
			stealth[x][y], know[x][y] = 0, false
			switch mon {
			case ROTHE, POLTERGEIST, VAMPIRE:
				stealth[x][y] = 1
			}
			return
		}
		k++
	}
}

/*
 * int cgood(x,y,itm,monst)	  Function to check location for emptiness
 * 	int x,y,itm,monst;
 *
 * Routine to return TRUE if a location does not have itm or monst there
 * returns FALSE (0) otherwise
 * Enter with itm or monst TRUE or FALSE if checking it
 * Example:  if itm==TRUE check for no item at this location
 * 		  if monst==TRUE check for no monster at this location
 * This routine will return FALSE if at a wall or the dungeon exit on level 1
 */
func cgood(x, y, theitem, monst int) bool {
	if y >= 0 && y <= MAXY-1 && x >= 0 && x <= MAXX-1 {
		/* within bounds? */
		if item[x][y] != OWALL { /* can't make anything on walls */
			/* is it free of items? */
			if theitem == 0 || item[x][y] == 0 {
				/* is it free of monsters? */
				if monst == 0 || mitem[x][y] == 0 {
					if level != 1 || x != 33 || y != MAXY-1 {
						/* not exit to level 1 */
						return true
					}
				}
			}
		}
	}
	return false
}

/*
 * createitem(it,arg) 	Routine to place an item next to the player
 * 	int it,arg;
 *
 * Enter with the item number and its argument (iven[], ivenarg[])
 * Returns no value, thus we don't know about createitem() failures.
 */
func createitem(it, arg int) {
	if it >= MAXOBJ {
		return /* no such object */
	}
	for k, i := rnd(8), -8; i < 0; i++ { /* choose direction, then try all */
		if k > 8 {
			k = 1 /* wraparound the diroff arrays */
		}
		x := playerx + diroffx[k]
		y := playery + diroffy[k]
		if cgood(x, y, 1, 0) { /* if we can create here */
			item[x][y] = it
			know[x][y] = false
			iarg[x][y] = arg
			return
		}
		k++
	}
}

/*
 * cast() 		Subroutine called by parse to cast a spell for the user
 *
 * No arguments and no return value.
 */
func cast() {
	cursors()
	if c[SPELLS] <= 0 {
		lprcat("\nYou don't have any spells!")
		return
	}
	const eys = "\nEnter your spell: "
	lprcat(eys)
	c[SPELLS]--
	var a int
	for {
		a = ttgetch()
		if a != 'D' {
			break
		}
		seemagic(-1)
		cursors()
		lprcat(eys)
	}
	if a == '\033' {
		/* to escape casting a spell	 */
		lprcat(" aborted")
		c[SPELLS]++
		return
	}
	b := ttgetch()
	if b == '\033' {
		/* to escape casting a spell	 */
		lprcat(" aborted")
		c[SPELLS]++
		return
	}
	d := ttgetch()
	if d == '\033' {
		/* to escape casting a spell	 */
		lprcat(" aborted")
		c[SPELLS]++
		return
	}
	c[SPELLSCAST]++
	lprc('\n')
	spell := fmt.Sprintf("%c%c%c", a, b, d)
	var i, j int
	for j, i = -1, 0; i < SPNUM; i++ { /* seq search for his spell, hash? */
		if spelcode[i] == spell {
			if spelknow[i] {
				speldamage(i)
				j = 1
				i = SPNUM
			}
		}
	}
	if j == -1 {
		lprcat("  Nothing Happened ")
	}
	bottomline()
}

func icond(x bool, a, b int) int {
	if x {
		return a
	}
	return b
}

func scond(x bool, a, b string) string {
	if x {
		return a
	}
	return b
}

/*
 * speldamage(x) 		Function to perform spell functions cast by the player
 * 	int x;
 *
 * Enter with the spell number, returns no value.
 * Please insure that there are 2 spaces before all messages here
 */
func speldamage(x int) {
	if x >= SPNUM {
		return /* no such spell */
	}
	if c[TIMESTOP] != 0 {
		lprcat("  It didn't seem to work")
		return
	} /* not if time stopped */
	clev := c[LEVEL]
	if rnd(23) == 7 || rnd(18) > c[INTELLIGENCE] {
		lprcat("  It didn't work!")
		return
	}
	if clev*3+2 < x {
		lprcat("  Nothing happens.  You seem inexperienced at this")
		return
	}
	//int    i, j
	//int    xl, xh, yl, yh;
	//u_char *p, *kn, *pm;
	switch x {
	/* ----- LEVEL 1 SPELLS ----- */

	case 0:
		if c[PROTECTIONTIME] == 0 {
			c[MOREDEFENSES] += 2 /* protection field +2 */
		}
		c[PROTECTIONTIME] += 250
		return

	case 1:
		i := rnd((clev+1)<<1) + clev + 3
		godirect(x, i, scond(clev >= 2, "  Your missiles hit the %s", "  Your missile hit the %s"), 100, '+') /* magic missile */

		return

	case 2:
		if c[DEXCOUNT] == 0 {
			c[DEXTERITY] += 3 /* dexterity	 */
		}
		c[DEXCOUNT] += 400
		return

	case 3: /* sleep		 */
		i := rnd(3) + 1
		direct(x, fullhit(i),
			"  While the %s slept, you smashed it %d times", i)
		return

	case 4: /* charm monster	 */
		c[CHARMCOUNT] += c[CHARISMA] << 1
		return

	case 5:
		godirect(x, rnd(10)+15+clev, "  The sound damages the %s", 70, '@') /* sonic spear */
		return

		/* ----- LEVEL 2 SPELLS ----- */

	case 6: /* web 			*/
		i := rnd(3) + 2
		direct(x, fullhit(i),
			"  While the %s is entangled, you hit %d times", i)
		return

	case 7:
		if c[STRCOUNT] == 0 {
			c[STREXTRA] += 3 /* strength	 */
		}
		c[STRCOUNT] += 150 + rnd(100)
		return

	case 8:
		yl := playery - 5 /* enlightenment */
		yh := playery + 6
		xl := playerx - 15
		xh := playerx + 16
		vxy(&xl, &yl)
		vxy(&xh, &yh)               /* check bounds */
		for i := yl; i <= yh; i++ { /* enlightenment	 */
			for j := xl; j <= xh; j++ {
				know[j][i] = true
			}
		}
		draws(xl, xh+1, yl, yh+1)
		return

	case 9:
		raisehp(20 + (clev << 1))
		return /* healing */

	case 10:
		c[BLINDCOUNT] = 0
		return /* cure blindness	 */

	case 11:
		createmonster(makemonst(level+1) + 8)
		return

	case 12:
		if rnd(11)+7 <= c[WISDOM] {
			direct(x, rnd(20)+20+clev, "  The %s believed!", 0)
		} else {
			lprcat("  It didn't believe the illusions!")
		}
		return

	case 13: /* if he has the amulet of invisibility then add more time */
		j := 0
		for i := 0; i < 26; i++ {
			if iven[i] == OAMULET {
				j += 1 + ivenarg[i]
			}
		}
		c[INVISIBILITY] += (j << 7) + 12
		return

		/* ----- LEVEL 3 SPELLS ----- */

	case 14:
		godirect(x, rnd(25+clev)+25+clev, "  The fireball hits the %s", 40, '*')
		return /* fireball */

	case 15:
		godirect(x, rnd(25)+20+clev, "  Your cone of cold strikes the %s", 60, 'O') /* cold */
		return

	case 16:
		dirpoly(x)
		return /* polymorph */

	case 17:
		c[CANCELLATION] += 5 + clev
		return /* cancellation	 */

	case 18:
		c[HASTESELF] += 7 + clev
		return /* haste self	 */

	case 19:
		omnidirect(x, 30+rnd(10), "  The %s gasps for air") /* cloud kill */
		return

	case 20:
		xh := min(playerx+1, MAXX-2)
		yh := min(playery+1, MAXY-2)
		for i := max(playerx-1, 1); i <= xh; i++ { /* vaporize rock */
			for j := max(playery-1, 1); j <= yh; j++ {
				kn := &know[i][j]
				pm := &mitem[i][j]
				switch p := &item[i][j]; *p {
				case OWALL:
					if level < MAXLEVEL+MAXVLEVEL-1 {
						*p, *kn = 0, false
					}

				case OSTATUE:
					if c[HARDGAME] < 3 {
						*p = OBOOK
						iarg[i][j] = level
						*kn = false
					}

				case OTHRONE:
					*pm = GNOMEKING
					*kn = false
					*p = OTHRONE2
					hitp[i][j] = monster[GNOMEKING].hitpoints

				case OALTAR:
					*pm = DEMONPRINCE
					*kn = false
					hitp[i][j] = monster[DEMONPRINCE].hitpoints
				}
				switch *pm {
				case XORN:
					ifblind(i, j)
					hitm(i, j, 200)
					/* Xorn takes damage from vpr */
				}
			}
		}
		return

		/* ----- LEVEL 4 SPELLS ----- */

	case 21:
		direct(x, 100+clev, "  The %s shrivels up", 0) /* dehydration */
		return

	case 22:
		godirect(x, rnd(25)+20+(clev<<1), "  A lightning bolt hits the %s", 1, '~') /* lightning */
		return

	case 23:
		i := min(c[HP]-1, c[HPMAX]/2) /* drain life */
		direct(x, i+i, "", 0)
		c[HP] -= i
		return

	case 24:
		if c[GLOBE] == 0 {
			c[MOREDEFENSES] += 10
		}
		c[GLOBE] += 200
		loseint() /* globe of invulnerability */
		return

	case 25:
		omnidirect(x, 32+clev, "  The %s struggles for air in your flood!") /* flood */
		return

	case 26:
		if rnd(151) == 63 {
			beep()
			lprcat("\nYour heart stopped!\n")
			nap(4000)
			died(270)
			return
		}
		if c[WISDOM] > rnd(10)+10 {
			direct(x, 2000, "  The %s's heart stopped", 0) /* finger of death */
		} else {
			lprcat("  It didn't work")
		}
		return

		/* ----- LEVEL 5 SPELLS ----- */

	case 27:
		c[SCAREMONST] += rnd(10) + clev
		return /* scare monster */

	case 28:
		c[HOLDMONST] += rnd(10) + clev
		return /* hold monster */

	case 29:
		c[TIMESTOP] += rnd(20) + (clev << 1)
		return /* time stop */

	case 30:
		tdirect(x)
		return /* teleport away */

	case 31:
		omnidirect(x, 35+rnd(10)+clev, "  The %s cringes from the flame") /* magic fire */
		return

		/* ----- LEVEL 6 SPELLS ----- */

	case 32:
		if rnd(23) == 5 && !wizard { /* sphere of annihilation */
			beep()
			lprcat("\nYou have been enveloped by the zone of nothingness!\n")
			nap(4000)
			died(258)
			return
		}
		xl := playerx
		yl := playery
		loseint()
		i := dirsub(&xl, &yl)            /* get direction of sphere */
		newsphere(xl, yl, i, rnd(20)+11) /* make a sphere */
		return

	case 33:
		genmonst()
		spelknow[33] = false /* genocide */
		loseint()
		return

	case 34: /* summon demon */
		if rnd(100) > 30 {
			direct(x, 150, "  The demon strikes at the %s", 0)
			return
		}
		if rnd(100) > 15 {
			lprcat("  Nothing seems to have happened")
			return
		}
		lprcat("  The demon turned on you and vanished!")
		beep()
		i := rnd(40) + 30
		lastnum = 277
		losehp(i) /* must say killed by a demon */
		return

	case 35: /* walk through walls */
		c[WTW] += rnd(10) + 5
		return

	case 36: /* alter reality */
		sc := 0 // # items saved
		save := make([]isave, MAXX*MAXY*2)
		for j := 0; j < MAXY; j++ {
			for i := 0; i < MAXX; i++ { // save all items and monsters
				xl := item[i][j]
				if xl != 0 && xl != OWALL && xl != OANNIHILATION {
					save[sc].typ = false
					save[sc].id = item[i][j]
					save[sc].arg = iarg[i][j]
					sc++
				}
				if mitem[i][j] != 0 {
					save[sc].typ = true
					save[sc].id = mitem[i][j]
					save[sc].arg = hitp[i][j]
					sc++
				}
				item[i][j] = OWALL
				mitem[i][j] = 0
				if wizard {
					know[i][j] = true
				} else {
					know[i][j] = false
				}
			}
		}
		eat(1, 1)
		if level == 1 {
			item[33][MAXY-1] = 0
		}
		for j, i := rnd(MAXY-2), 1; i < MAXX-1; i++ {
			item[i][j] = 0
		}
		for sc > 0 { // put objects back in level
			sc--
			if !save[sc].typ {
				trys := 100
				i, j := 1, 1
				for trys > 1 && item[i][j] != 0 {
					trys--
					i = rnd(MAXX - 1)
					j = rnd(MAXY - 1)
				}
				if trys > 1 {
					item[i][j] = save[sc].id
					iarg[i][j] = save[sc].arg
				}
			} else { // put monsters back in
				trys := 100
				i, j := 1, 1
				for trys > 1 && (item[i][j] == OWALL || mitem[i][j] != 0) {
					trys--
					i = rnd(MAXX - 1)
					j = rnd(MAXY - 1)
				}
				if trys > 1 {
					mitem[i][j] = save[sc].id
					hitp[i][j] = save[sc].arg
				}
			}
		}
		loseint()
		draws(0, MAXX, 0, MAXY)
		if !wizard {
			spelknow[36] = false
		}
		positionplayer()
		return

	case 37: /* permanence */
		adjusttime(-99999)
		spelknow[37] = false /* forget */
		loseint()
		return

	default:
		lprintf("  spell %d not available!", x)
		beep()
		return
	}
}

/*
 * loseint()		Routine to subtract 1 from your int (intelligence) if > 3
 *
 * No arguments and no return value
 */
func loseint() {
	if c[INTELLIGENCE] > 3 {
		c[INTELLIGENCE]--
	}
}

/*
 * isconfuse() 		Routine to check to see if player is confused
 *
 * This routine prints out a message saying "You can't aim your magic!"
 * returns 0 if not confused, non-zero (time remaining confused) if confused
 */
func isconfuse() bool {
	if c[CONFUSE] != 0 {
		lprcat(" You can't aim your magic!")
		beep()
		return true
	}
	return false
}

/*
 * nospell(x,monst)	Routine to return 1 if a spell doesn't affect a monster
 * 	int x,monst;
 *
 * Subroutine to return 1 if the spell can't affect the monster
 *   otherwise returns 0
 * Enter with the spell number in x, and the monster number in monst.
 */
func nospell(x, monst int) bool {
	if x >= SPNUM || monst >= MAXMONST+8 || monst < 0 || x < 0 {
		return false /* bad spell or monst */
	}
	tmp := spelweird[monst-1][x]
	if tmp == 0 {
		return false
	}
	cursors()
	lprc('\n')
	lprintf(spelmes[tmp], monster[monst].name)
	return true
}

/*
 * fullhit(xx)		Function to return full damage against a monster (aka web)
 * 	int xx;
 *
 * Function to return hp damage to monster due to a number of full hits
 * Enter with the number of full hits being done
 */
func fullhit(xx int) int {
	if xx < 0 || xx > 20 {
		return 0 /* fullhits are out of range */
	}
	if c[LANCEDEATH] != 0 {
		return 10000 /* lance of death */
	}
	i := xx * ((c[WCLASS] >> 1) + c[STRENGTH] + c[STREXTRA] - c[HARDGAME] - 12 + c[MOREDAM])
	return icond(i >= 1, i, xx)
}

/*
 * direct(spnum,dam,str,arg)	Routine to direct spell damage 1 square in 1 dir
 * 	int spnum,dam,arg;
 * 	char *str;
 *
 * Routine to ask for a direction to a spell and then hit the monster
 * Enter with the spell number in spnum, the damage to be done in dam,
 *   lprintf format string in str, and lprintf's argument in arg.
 * Returns no value.
 */
func direct(spnum, dam int, str string, arg int) {
	if spnum < 0 || spnum >= SPNUM || str == "" {
		return /* bad arguments */
	}
	if isconfuse() {
		return
	}
	var x, y int
	dirsub(&x, &y)
	m := mitem[x][y]
	if item[x][y] == OMIRROR {
		fool := func() {
			beep()
			arg += 2
			for {
				arg--
				if arg < 0 {
					break
				}
				parse2()
				nap(1000)
			}
		}
		if spnum == 3 { /* sleep */
			lprcat("You fall asleep! ")
			fool() // closure above
			return
		} else if spnum == 6 { /* web */
			lprcat("You get stuck in your own web! ")
			fool() // closure above
			return
		} else {
			lastnum = 278
			lprintf(str, "spell caster (that's you)", arg)
			beep()
			losehp(dam)
			return
		}
	}
	if m == 0 {
		lprcat("  There wasn't anything there!")
		return
	}
	ifblind(x, y)
	if nospell(spnum, m) {
		lasthx = x
		lasthy = y
		return
	}
	lprintf(str, lastmonst, arg)
	hitm(x, y, dam)
}

/*
 * godirect(spnum,dam,str,delay,cshow)		Function to perform missile attacks
 * 	int spnum,dam,delay;
 * 	char *str,cshow;
 *
 * Function to hit in a direction from a missile weapon and have it keep
 * on going in that direction until its power is exhausted
 * Enter with the spell number in spnum, the power of the weapon in hp,
 *   lprintf format string in str, the # of milliseconds to delay between
 *   locations in delay, and the character to represent the weapon in cshow.
 * Returns no value.
 */
func godirect(spnum, dam int, str string, delay, cshow int) {
	if spnum < 0 || spnum >= SPNUM || str == "" || delay < 0 {
		return /* bad args */
	}
	if isconfuse() {
		return
	}
	var dx, dy int
	dirsub(&dx, &dy)
	x := dx
	y := dy
	dx = x - playerx
	dy = y - playery
	x = playerx
	y = playery
	for dam > 0 {
		x += dx
		y += dy
		if x > MAXX-1 || y > MAXY-1 || x < 0 || y < 0 {
			dam = 0
			break /* out of bounds */
		}
		if x == playerx && y == playery { /* if energy hits player */
			cursors()
			lprcat("\nYou are hit by your own magic!")
			beep()
			lastnum = 278
			losehp(dam)
			return
		}
		if c[BLINDCOUNT] == 0 { /* if not blind show effect */
			cursor(x+1, y+1)
			lprc(byte(cshow))
			nap(delay)
			show1cell(x, y)
		}
		if m := mitem[x][y]; m != 0 { /* is there a monster there? */
			ifblind(x, y)
			if nospell(spnum, m) {
				lasthx = x
				lasthy = y
				return
			}
			cursors()
			lprc('\n')
			lprintf(str, lastmonst)
			dam -= hitm(x, y, dam)
			show1cell(x, y)
			nap(1000)
			x -= dx
			y -= dy
		} else {
			switch p := &item[x][y]; *p {
			case OWALL:
				cursors()
				lprc('\n')
				lprintf(str, "wall")
				if dam >= 50+c[HARDGAME] { /* enough damage? */
					if level < MAXLEVEL+MAXVLEVEL-1 { /* not on V3 */
						if x < MAXX-1 && y < MAXY-1 && x != 0 && y != 0 {
							lprcat("  The wall crumbles")
							*p = 0
							know[x][y] = false
							show1cell(x, y)
						}
					}
				}
				dam = 0

			case OCLOSEDDOOR:
				cursors()
				lprc('\n')
				lprintf(str, "door")
				if dam >= 40 {
					lprcat("  The door is blasted apart")
					*p = 0
					know[x][y] = false
					show1cell(x, y)
				}
				dam = 0

			case OSTATUE:
				cursors()
				lprc('\n')
				lprintf(str, "statue")
				if c[HARDGAME] < 3 {
					if dam > 44 {
						lprcat("  The statue crumbles")
						*p = OBOOK
						iarg[x][y] = level
						know[x][y] = false
						show1cell(x, y)
					}
				}
				dam = 0

			case OTHRONE:
				cursors()
				lprc('\n')
				lprintf(str, "throne")
				if dam > 39 {
					mitem[x][y] = GNOMEKING
					hitp[x][y] = monster[GNOMEKING].hitpoints
					*p = OTHRONE2
					know[x][y] = false
					show1cell(x, y)
				}
				dam = 0

			case OMIRROR:
				dx *= -1
				dy *= -1
			}
		}
		dam -= 3 + (c[HARDGAME] >> 1)
	}
}

/*
 * ifblind(x,y)	Routine to put "monster" or the monster name into lastmosnt
 * 	int x,y;
 *
 * Subroutine to copy the word "monster" into lastmonst if the player is blind
 * Enter with the coordinates (x,y) of the monster
 * Returns no value.
 */
func ifblind(x, y int) {
	vxy(&x, &y) /* verify correct x,y coordinates */
	if c[BLINDCOUNT] != 0 {
		lastnum = 279
		lastmonst = "monster"
	} else {
		lastnum = mitem[x][y]
		lastmonst = monster[lastnum].name
	}
}

/*
 * tdirect(spnum)		Routine to teleport away a monster
 * 	int spnum;
 *
 * Routine to ask for a direction to a spell and then teleport away monster
 * Enter with the spell number that wants to teleport away
 * Returns no value.
 */
func tdirect(spnum int) {
	if spnum < 0 || spnum >= SPNUM {
		return /* bad args */
	}
	if isconfuse() {
		return
	}
	var x, y int
	dirsub(&x, &y)
	m := mitem[x][y]
	if m == 0 {
		lprcat("  There wasn't anything there!")
		return
	}
	ifblind(x, y)
	if nospell(spnum, m) {
		lasthx = x
		lasthy = y
		return
	}
	fillmonst(m)
	mitem[x][y], know[x][y] = 0, false
}

/*
 * omnidirect(sp,dam,str)   Routine to damage all monsters 1 square from player
 * 	int sp,dam;
 * 	char *str;
 *
 * Routine to cast a spell and then hit the monster in all directions
 * Enter with the spell number in sp, the damage done to wach square in dam,
 *   and the lprintf string to identify the spell in str.
 * Returns no value.
 */
func omnidirect(spnum, dam int, str string) {
	if spnum < 0 || spnum >= SPNUM || str == "" {
		return /* bad args */
	}
	for x := playerx - 1; x < playerx+2; x++ {
		for y := playery - 1; y < playery+2; y++ {
			if m := mitem[x][y]; m != 0 {
				if !nospell(spnum, m) {
					ifblind(x, y)
					cursors()
					lprc('\n')
					lprintf(str, lastmonst)
					hitm(x, y, dam)
					nap(800)
				} else {
					lasthx = x
					lasthy = y
				}
			}
		}
	}
}

/*
 * static dirsub(x,y)		Routine to ask for direction, then modify x,y for it
 * 	int *x,*y;
 *
 * Function to ask for a direction and modify an x,y for that direction
 * Enter with the origination coordinates in (x,y).
 * Returns index into diroffx[] (0-8).
 */
func dirsub(x, y *int) int {
	var i int
	lprcat("\nIn What Direction? ")
	for {
		switch ttgetch() {
		case 'b':
			i++
			fallthrough
		case 'n':
			i++
			fallthrough
		case 'y':
			i++
			fallthrough
		case 'u':
			i++
			fallthrough
		case 'h':
			i++
			fallthrough
		case 'k':
			i++
			fallthrough
		case 'l':
			i++
			fallthrough
		case 'j':
			i++
			goto out
		}
	}
out:
	*x = playerx + diroffx[i]
	*y = playery + diroffy[i]
	vxy(x, y)
	return i
}

/*
 * vxy(x,y)	   Routine to verify/fix coordinates for being within bounds
 * 	int *x,*y;
 *
 * Function to verify x & y are within the bounds for a level
 * If *x or *y is not within the absolute bounds for a level, fix them so that
 *   they are on the level.
 * Returns TRUE if it was out of bounds, and the *x & *y in the calling
 * routine are affected.
 */
func vxy(x, y *int) bool {
	flag := false
	if *x < 0 {
		*x = 0
		flag = true
	}
	if *y < 0 {
		*y = 0
		flag = true
	}
	if *x >= MAXX {
		*x = MAXX - 1
		flag = true
	}
	if *y >= MAXY {
		*y = MAXY - 1
		flag = true
	}
	return flag
}

/*
 * dirpoly(spnum)	Routine to ask for a direction and polymorph a monst
 * 	int spnum;
 *
 * Subroutine to polymorph a monster and ask for the direction its in
 * Enter with the spell number in spmun.
 * Returns no value.
 */
func dirpoly(spnum int) {
	if spnum < 0 || spnum >= SPNUM {
		return /* bad args */
	}
	if isconfuse() {
		return /* if he is confused, he can't aim his magic */
	}
	var x, y int
	dirsub(&x, &y)
	if mitem[x][y] == 0 {
		lprcat("  There wasn't anything there!")
		return
	}
	ifblind(x, y)
	if nospell(spnum, mitem[x][y]) {
		lasthx = x
		lasthy = y
		return
	}
	var m int
	for {
		m = rnd(MAXMONST + 7)
		mitem[x][y] = m
		if !monster[m].genocided {
			break
		}
	}
	hitp[x][y] = monster[m].hitpoints
	show1cell(x, y) /* show the new monster */
}

/*
 * hitmonster(x,y) 	Function to hit a monster at the designated coordinates
 * 	int x,y;
 *
 * This routine is used for a bash & slash type attack on a monster
 * Enter with the coordinates of the monster in (x,y).
 * Returns no value.
 */
func hitmonster(x, y int) {
	if c[TIMESTOP] != 0 {
		return /* not if time stopped */
	}
	vxy(&x, &y) /* verify coordinates are within range */
	monst := mitem[x][y]
	if monst == 0 {
		return
	}
	hit3flag = true
	ifblind(x, y)
	tmp := monster[monst].armorclass + c[LEVEL] + c[DEXTERITY] + c[WCLASS]/4 - 12
	cursors()
	/* need at least random chance to hit */
	var damag int
	var flag bool
	if rnd(20) < tmp-c[HARDGAME] || rnd(71) < 5 {
		lprcat("\nYou hit")
		flag = true
		damag = fullhit(1)
		if damag < 9999 {
			damag = rnd(damag) + 1
		}
	} else {
		lprcat("\nYou missed")
		flag = false
	}
	lprcat(" the ")
	lprcat(lastmonst)
	if flag { /* if the monster was hit */
		if monst == RUSTMONSTER || monst == DISENCHANTRESS || monst == CUBE {
			if c[WIELD] > 0 {
				if ivenarg[c[WIELD]] > -10 {
					lprintf("\nYour weapon is dulled by the %s", lastmonst)
					beep()
					ivenarg[c[WIELD]]--
				}
			}
		}
	}
	if flag {
		hitm(x, y, damag)
	}
	if monst == VAMPIRE {
		if hitp[x][y] < 25 {
			mitem[x][y] = BAT
			know[x][y] = false
		}
	}
}

/*
 * hitm(x,y,amt)	Function to just hit a monster at a given coordinates
 * 	int x,y,amt;
 *
 * Returns the number of hitpoints the monster absorbed
 * This routine is used to specifically damage a monster at a location (x,y)
 * Called by hitmonster(x,y)
 */
func hitm(x, y, amt int) int {
	vxy(&x, &y) /* verify coordinates are within range */
	amt2 := amt /* save initial damage so we can return it */
	monst := mitem[x][y]
	if c[HALFDAM] != 0 {
		amt >>= 1 /* if half damage curse adjust damage points */
	}
	if amt <= 0 {
		amt2, amt = 1, 1
	}
	lasthx = x
	lasthy = y
	stealth[x][y] = 1 /* make sure hitting monst breaks stealth condition */
	c[HOLDMONST] = 0  /* hit a monster breaks hold monster spell	 */
	switch monst {    /* if a dragon and orbs of dragon slaying	 */
	case WHITEDRAGON, REDDRAGON, GREENDRAGON, BRONZEDRAGON, PLATINUMDRAGON, SILVERDRAGON:
		amt *= 1 + (c[SLAYING] << 1)
	}
	/* invincible monster fix is here */
	if hitp[x][y] > monster[monst].hitpoints {
		hitp[x][y] = monster[monst].hitpoints
	}
	hpoints := hitp[x][y]
	if hpoints <= amt {
		c[MONSTKILLED]++
		lprintf("\nThe %s died!", lastmonst)
		raiseexperience(monster[monst].experience)
		amt = monster[monst].gold
		if amt > 0 {
			dropgold(rnd(amt) + amt)
		}
		dropsomething(monst)
		disappear(x, y)
		bottomline()
		return hpoints
	}
	hitp[x][y] = hpoints - amt
	return amt2
}

/*
 * hitplayer(x,y) 	Function for the monster to hit the player from (x,y)
 * 	int x,y;
 *
 * Function for the monster to hit the player with monster at location x,y
 * Returns nothing of value.
 */
func hitplayer(x, y int) {
	vxy(&x, &y) /* verify coordinates are within range */
	mster := mitem[x][y]
	lastnum = mster
	/*
	 * spirit nagas and poltergeists do nothing if scarab of negate
	 * spirit
	 */
	if c[NEGATESPIRIT] != 0 || c[SPIRITPRO] != 0 {
		if mster == POLTERGEIST || mster == SPIRITNAGA {
			return
		}
	}
	/* if undead and cube of undead control	 */
	if c[CUBEofUNDEAD] != 0 || c[UNDEADPRO] != 0 {
		if mster == VAMPIRE || mster == WRAITH || mster == ZOMBIE {
			return
		}
	}
	if !know[x][y] {
		know[x][y] = true
		show1cell(x, y)
	}
	bias := c[HARDGAME] + 1
	hitflag, hit2flag, hit3flag = true, true, true
	yrepcount = 0
	cursors()
	ifblind(x, y)
	if c[INVISIBILITY] != 0 {
		if rnd(33) < 20 {
			lprintf("\nThe %s misses wildly", lastmonst)
			return
		}
	}
	if c[CHARMCOUNT] != 0 {
		if rnd(30)+5*monster[mster].level-c[CHARISMA] < 30 {
			lprintf("\nThe %s is awestruck at your magnificence!", lastmonst)
			return
		}
	}
	var dam int
	if mster == BAT {
		dam = 1
	} else {
		dam = monster[mster].damage
		dam += rnd(icond(dam < 1, 1, dam)) + monster[mster].level
	}
	tmp := 0
	if monster[mster].attack > 0 {
		if dam+bias+8 > c[AC] || rnd(icond(c[AC] > 0, c[AC], 1)) == 1 {
			if spattack(monster[mster].attack, x, y) {
				flushall()
				return
			}
			tmp = 1
			bias -= 2
			cursors()
		}
	}
	if dam+bias > c[AC] || rnd(icond(c[AC] > 0, c[AC], 1)) == 1 {
		lprintf("\n  The %s hit you ", lastmonst)
		tmp = 1
		dam -= c[AC]
		if dam < 0 {
			dam = 0
		}
		if dam > 0 {
			losehp(dam)
			bottomhp()
			flushall()
		}
	}
	if tmp == 0 {
		lprintf("\n  The %s missed ", lastmonst)
	}
}

/*
 * dropsomething(monst) 	Function to create an object when a monster dies
 * 	int monst;
 *
 * Function to create an object near the player when certain monsters are killed
 * Enter with the monster number
 * Returns nothing of value.
 */
func dropsomething(monst int) {
	switch monst {
	case ORC, NYMPH, ELF, TROGLODYTE, TROLL, ROTHE, VIOLETFUNGI, PLATINUMDRAGON, GNOMEKING, REDDRAGON:
		something(level)
		return

	case LEPRECHAUN:
		if rnd(101) >= 75 {
			creategem()
		}
		if rnd(5) == 1 {
			dropsomething(LEPRECHAUN)
		}
		return
	}
}

/*
 * dropgold(amount) 	Function to drop some gold around player
 * 	int amount;
 *
 * Enter with the number of gold pieces to drop
 * Returns nothing of value.
 */
func dropgold(amount int) {
	if amount > 250 {
		createitem(OMAXGOLD, amount/100)
	} else {
		createitem(OGOLDPILE, amount)
	}
}

/*
 * something(level) 	Function to create a random item around player
 * 	int level;
 *
 * Function to create an item from a designed probability around player
 * Enter with the cave level on which something is to be dropped
 * Returns nothing of value.
 */
func something(cavelevel int) {
	if cavelevel < 0 || cavelevel > MAXLEVEL+MAXVLEVEL {
		return /* correct level? */
	}
	if rnd(101) < 8 {
		something(cavelevel) /* possibly more than one item */
	}
	var i int
	j := newobject(cavelevel, &i)
	createitem(j, i)
}

/*
 * newobject(lev,i) 	Routine to return a randomly selected new object
 * 	int lev,*i;
 *
 * Routine to return a randomly selected object to be created
 * Returns the object number created, and sets *i for its argument
 * Enter with the cave level and a pointer to the items arg
 */
var nobjtab = [...]int{
	0, OSCROLL, OSCROLL, OSCROLL, OSCROLL, OPOTION, OPOTION,
	OPOTION, OPOTION, OGOLDPILE, OGOLDPILE, OGOLDPILE, OGOLDPILE,
	OBOOK, OBOOK, OBOOK, OBOOK, ODAGGER, ODAGGER, ODAGGER,
	OLEATHER, OLEATHER, OLEATHER, OREGENRING, OPROTRING,
	OENERGYRING, ODEXRING, OSTRRING, OSPEAR, OBELT, ORING,
	OSTUDLEATHER, OSHIELD, OFLAIL, OCHAIN, O2SWORD, OPLATE,
	OLONGSWORD}

func newobject(lev int, i *int) int {
	if level < 0 || level > MAXLEVEL+MAXVLEVEL {
		return 0 /* correct level? */
	}
	tmp := 32
	if lev > 6 {
		tmp = 37
	} else if lev > 4 {
		tmp = 35
	}
	tmp = rnd(tmp)
	j := nobjtab[tmp] /* the object type */
	switch tmp {
	case 1, 2, 3, 4:
		*i = newscroll()
	case 5, 6, 7, 8:
		*i = newpotion()
	case 9, 10, 11, 12:
		*i = rnd((lev+1)*10) + lev*10 + 10
	case 13, 14, 15, 16:
		*i = lev
	case 17, 18, 19:
		*i = newdagger()
		if *i == 0 {
			return 0
		}
	case 20, 21, 22:
		*i = newleather()
		if *i == 0 {
			return 0
		}
	case 23, 32, 35:
		*i = rund(lev/3 + 1)
		break
	case 24, 26:
		*i = rnd(lev/4 + 1)
	case 25:
		*i = rund(lev/4 + 1)
	case 27:
		*i = rnd(lev/2 + 1)
	case 30, 33:
		*i = rund(lev/2 + 1)
	case 28:
		*i = rund(lev/3 + 1)
		if *i == 0 {
			return 0
		}
	case 29, 31:
		*i = rund(lev/2 + 1)
		if *i == 0 {
			return 0
		}
	case 34:
		*i = newchain()
	case 36:
		*i = newplate()
	case 37:
		*i = newsword()
	}
	return j
}

/*
 *  spattack(atckno,xx,yy) Function to process special attacks from monsters
 *  	int atckno,xx,yy;
 *
 * Enter with the special attack number, and the coordinates (xx,yy)
 * 	of the monster that is special attacking
 * Returns 1 if must do a show1cell(xx,yy) upon return, 0 otherwise
 *
 * atckno   monster     effect
 * ---------------------------------------------------
 * 0	none
 * 1	rust monster	eat armor
 * 2	hell hound	breathe light fire
 * 3	dragon		breathe fire
 * 4	giant centipede	weakening sing
 * 5	white dragon	cold breath
 * 6	wraith		drain level
 * 7	waterlord	water gusher
 * 8	leprechaun	steal gold
 * 9	disenchantress	disenchant weapon or armor
 * 10	ice lizard	hits with barbed tail
 * 11	umber hulk	confusion
 * 12	spirit naga	cast spells	taken from special attacks
 * 13	platinum dragon	psionics
 * 14	nymph		steal objects
 * 15	bugbear		bite
 * 16	osequip		bite
 *
 * char rustarm[ARMORTYPES][2];
 * special array for maximum rust damage to armor from rustmonster
 * format is: { armor type , minimum attribute
 */
const ARMORTYPES = 6

var rustarm = [ARMORTYPES][2]int{
	{OSTUDLEATHER, -2},
	{ORING, -4},
	{OCHAIN, -5},
	{OSPLINT, -6},
	{OPLATE, -8},
	{OPLATEARMOR, -9},
}
var spsel = []int{1, 2, 3, 5, 6, 8, 9, 11, 13, 14}

func spattack(x, xx, yy int) bool {
	var p string

	if c[CANCELLATION] != 0 {
		return false
	}
	vxy(&xx, &yy) /* verify x & y coordinates */
	var i int
	spout2 := func() {
		if p != "" {
			lprintf(p, lastmonst)
			beep()
		}
	}
	switch x {
	case 1: /* rust your armor, j=1 when rusting has occurred */
		k := c[WEAR]
		m := k
		i = c[SHIELD]
		var j int
		if i != -1 {
			ivenarg[i]--
			if ivenarg[i] < -1 {
				ivenarg[i] = -1
			} else {
				j = 1
			}
		}
		if j == 0 && k != -1 {
			m = iven[k]
			for i = 0; i < ARMORTYPES; i++ {
				/* find his armor in table */
				if m == rustarm[i][0] {
					ivenarg[k]--
					if ivenarg[k] < rustarm[i][1] {
						ivenarg[k] = rustarm[i][1]
					} else {
						j = 1
					}
					break
				}
			}
		}
		if j == 0 { /* if rusting did not occur */
			switch m {
			case OLEATHER:
				p = "\nThe %s hit you -- You're lucky you have leather on"
			case OSSPLATE:
				p = "\nThe %s hit you -- You're fortunate to have stainless steel armor!"
			}
		} else {
			beep()
			p = "\nThe %s hit you -- your armor feels weaker"
		}

	case 2, 3:
		if x == 2 {
			i = rnd(15) + 8 - c[AC]
		} else if x == 3 {
			i = rnd(20) + 25 - c[AC]
		}
		p = "\nThe %s breathes fire at you!"
		if c[FIRERESISTANCE] != 0 {
			p = "\nThe %s's flame doesn't faze you!"
		} else {
			spout2()
		}
		checkloss(i)
		return false

	case 4:
		if c[STRENGTH] > 3 {
			p = "\nThe %s stung you!  You feel weaker"
			beep()
			c[STRENGTH]--
		} else {
			p = "\nThe %s stung you!"
		}

	case 5:
		p = "\nThe %s blasts you with his cold breath"
		i = rnd(15) + 18 - c[AC]
		spout2()
		checkloss(i)
		return false

	case 6:
		lprintf("\nThe %s drains you of your life energy!", lastmonst)
		loselevel()
		beep()
		return false

	case 7:
		p = "\nThe %s got you with a gusher!"
		i = rnd(15) + 25 - c[AC]
		spout2()
		checkloss(i)
		return false

	case 8:
		if c[NOTHEFT] != 0 {
			return false /* he has a device of no theft */
		}
		if c[GOLD] > 0 {
			p = "\nThe %s hit you -- Your purse feels lighter"
			if c[GOLD] > 32767 {
				c[GOLD] >>= 1
			} else {
				c[GOLD] -= rnd(1 + (c[GOLD] >> 1))
			}
			if c[GOLD] < 0 {
				c[GOLD] = 0
			}
		} else {
			p = "\nThe %s couldn't find any gold to steal"
		}
		lprintf(p, lastmonst)
		disappear(xx, yy)
		beep()
		bottomgold()
		return true

	case 9:
		j := 50
		for { /* disenchant */
			i = rund(26)
			m := iven[i] /* randomly select item */
			if m > 0 && ivenarg[i] > 0 && m != OSCROLL && m != OPOTION {
				ivenarg[i] -= 3
				if ivenarg[i] < 0 {
					ivenarg[i] = 0
				}
				lprintf("\nThe %s hits you -- you feel a sense of loss", lastmonst)
				srcount = 0
				beep()
				show3(i)
				bottomline()
				return false
			}
			j--
			if j <= 0 {
				p = "\nThe %s nearly misses"
				break
			}
			break
		}

	case 10:
		p = "\nThe %s hit you with his barbed tail"
		i = rnd(25) - c[AC]
		spout2()
		checkloss(i)
		return false

	case 11:
		p = "\nThe %s has confused you"
		beep()
		c[CONFUSE] += 10 + rnd(10)

	case 12: /* performs any number of other special attacks	 */
		return spattack(spsel[rund(10)], xx, yy)

	case 13:
		p = "\nThe %s flattens you with his psionics!"
		i = rnd(15) + 30 - c[AC]
		spout2()
		checkloss(i)
		return false

	case 14:
		if c[NOTHEFT] != 0 {
			return false /* he has device of no theft */
		}
		if emptyhanded() == 1 {
			p = "\nThe %s couldn't find anything to steal"
			break
		}
		lprintf("\nThe %s picks your pocket and takes:", lastmonst)
		beep()
		if stealsomething() == 0 {
			lprcat("  nothing")
		}
		disappear(xx, yy)
		bottomline()
		return true

	case 15, 16:
		if x == 15 {
			i = rnd(10) + 5 - c[AC]
		} else if x == 16 {
			i = rnd(15) + 10 - c[AC]
		}
		p = "\nThe %s bit you!"
		spout2()
		checkloss(i)
		return false
	}
	if p != "" {
		lprintf(p, lastmonst)
		bottomline()
	}
	return false
}

/*
 * checkloss(x) Routine to subtract hp from user and flag bottomline display
 * 	int x;
 *
 * Routine to subtract hitpoints from the user and flag the bottomline display
 * Enter with the number of hit points to lose
 * Note: if x > c[HP] this routine could kill the player!
 */
func checkloss(x int) {
	if x > 0 {
		losehp(x)
		bottomhp()
	}
}

/*
 * annihilate() 	Routine to annihilate all monsters around player (playerx,playery)
 *
 * Gives player experience, but no dropped objects
 * Returns the experience gained from all monsters killed
 */
func annihilate() int {
	var k int
	for i := playerx - 1; i <= playerx+1; i++ {
		for j := playery - 1; j <= playery+1; j++ {
			if !vxy(&i, &j) { /* if not out of bounds */
				p := &mitem[i][j]
				if *p != 0 { /* if a monster there */
					if *p < DEMONLORD+2 {
						k += monster[*p].experience
						*p, know[i][j] = 0, false
					} else {
						lprintf("\nThe %s barely escapes being annihilated!", monster[*p].name)
						hitp[i][j] = (hitp[i][j] >> 1) + 1 /* lose half hit points */
					}
				}
			}
		}
	}
	if k > 0 {
		lprcat("\nYou hear loud screams of agony!")
		raiseexperience(k)
	}
	return k
}

/*
 * newsphere(x,y,dir,lifetime)  Function to create a new sphere of annihilation
 * 	int x,y,dir,lifetime;
 *
 * Enter with the coordinates of the sphere in x,y
 *   the direction (0-8 diroffx format) in dir, and the lifespan of the
 *   sphere in lifetime (in turns)
 * Returns the number of spheres currently in existence
 */
func newsphere(x, y, dir, life int) int {
	sp := new(sphere)
	if dir >= 9 {
		dir = 0 // no movement if direction not found
	}
	if level == 0 {
		vxy(&x, &y) // don't go out of bounds
	} else {
		if x < 1 {
			x = 1
		}
		if x >= MAXX-1 {
			x = MAXX - 2
		}
		if y < 1 {
			y = 1
		}
		if y >= MAXY-1 {
			y = MAXY - 2
		}
	}
	m := mitem[x][y]
	if m >= DEMONLORD+4 { // demons dispel spheres
		know[x][y] = true
		show1cell(x, y) // show the demon (ha ha)
		cursors()
		lprintf("\nThe %s dispels the sphere!", monster[m].name)
		beep()
		rmsphere(x, y) // remove any spheres that are here
		return c[SPHCAST]
	}
	if m == DISENCHANTRESS { // disenchantress cancels spheres
		cursors()
		lprintf("\nThe %s causes cancellation of the sphere!", monster[m].name)
		beep()
		goto boom
	}
	if c[CANCELLATION] != 0 { // cancellation cancels spheres
		cursors()
		lprcat("\nAs the cancellation takes effect, you hear a great earth shaking blast!")
		beep()
		goto boom
	}
	if item[x][y] == OANNIHILATION { // collision of spheres detonates spheres
		cursors()
		lprcat("\nTwo spheres of annihilation collide! You hear a great earth shaking blast!")
		beep()
		rmsphere(x, y)
		goto boom
	}
	if playerx == x && playery == y { // collision of sphere and player!
		cursors()
		lprcat("\nYou have been enveloped by the zone of nothingness!\n")
		beep()
		rmsphere(x, y) // remove any spheres that are here
		nap(4000)
		died(258)
	}
	item[x][y] = OANNIHILATION
	mitem[x][y] = 0
	know[x][y] = true
	show1cell(x, y) // show the new sphere
	sp.x = x
	sp.y = y
	sp.lev = level
	sp.dir = dir
	sp.lifetime = life
	sp.p = nil
	if spheres == nil {
		// if first node in the sphere list
		spheres = sp
	} else {
		// add sphere to beginning of linked list
		sp.p = spheres
		spheres = sp
	}
	c[SPHCAST]++ // one more sphere in the world
	return c[SPHCAST]

boom:
	sphboom(x, y) // blow up stuff around sphere
	rmsphere(x, y) // remove any spheres that are here
	return c[SPHCAST]
}

/*
 * rmsphere(x,y)		Function to delete a sphere of annihilation from list
 * 	int x,y;
 *
 * Enter with the coordinates of the sphere (on current level)
 * Returns the number of spheres currently in existence
 */
func rmsphere(x, y int) int {
	for sp, sp2 := spheres, (*sphere)(nil); sp != nil; sp2, sp = sp, sp.p {
		if level == sp.lev { // is sphere on this level?
			if x == sp.x && y == sp.y { // locate sphere at this location
				item[x][y], mitem[x][y] = 0, 0
				know[x][y] = true
				show1cell(x, y) // show the now missing sphere
				c[SPHCAST]--
				if sp == spheres {
					sp2 = sp
					spheres = sp.p
				} else {
					if sp2 != nil {
						sp2.p = sp.p
					}
				}
				break
			}
		}
	}
	return c[SPHCAST] /* return number of spheres in the world */
}

/*
 * sphboom(x,y)	Function to perform the effects of a sphere detonation
 * 	int x,y;
 *
 * Enter with the coordinates of the blast, Returns no value
 */
func sphboom(x, y int) {
	if c[HOLDMONST] != 0 {
		c[HOLDMONST] = 1
	}
	if c[CANCELLATION] != 0 {
		c[CANCELLATION] = 1
	}
	for j := max(1, x-2); j < min(x+3, MAXX-1); j++ {
		for i := max(1, y-2); i < min(y+3, MAXY-1); i++ {
			item[j][i], mitem[j][i] = 0, 0
			show1cell(j, i)
			if playerx == j && playery == i {
				cursors()
				beep()
				lprcat("\nYou were too close to the sphere!")
				nap(3000)
				died(283) /* player killed in explosion */
			}
		}
	}
}

func isalpha(x int) bool {
	return ('A' <= x && x <= 'Z') || ('a' <= x && x <= 'z')
}

/*
 * genmonst()		Function to ask for monster and genocide from game
 *
 * This is done by setting a flag in the monster[] structure
 */
func genmonst() {
	cursors()
	lprcat("\nGenocide what monster? ")
	var i int
	for {
		i = ttgetch()
		if isalpha(i) || i == ' ' {
			break
		}
	}
	lprc(byte(i))
	for j := 0; j < MAXMONST; j++ { /* search for the monster type */
		if monstnamelist[j] == byte(i) { /* have we found it? */
			monster[j].genocided = true /* genocided from game */
			lprintf("  There will be no more %s's", monster[j].name)
			/* now wipe out monsters on this level */
			newcavelevel(level)
			draws(0, MAXX, 0, MAXY)
			bot_linex()
			return
		}
	}
	lprcat("  You sense failure!")
}
