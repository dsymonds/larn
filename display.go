package main

func makecode(a, b, c int) int {
	return a<<16 + b<<8 + c
}

var minx, maxx, miny, maxy, k, m int
var bot1f, bot2f, bot3f bool
var always bool

/*
	bottomline()

	now for the bottom line of the display
*/
func bottomline() {
	recalc()
	bot1f = true
}

func bottomhp() {
	bot2f = true
}

func bottomspell() {
	bot3f = true
}

func bottomdo() {
	if bot1f {
		bot3f, bot1f, bot2f = false, false, false
		bot_linex()
		return
	}
	if bot2f {
		bot2f = false
		bot_hpx()
	}
	if bot3f {
		bot3f = false
		bot_spellx()
	}
}

func bot_linex() {
	if cbak[SPELLS] <= -50 || always {
		cursor(1, 18)
		if c[SPELLMAX] > 99 {
			lprintf("Spells:%3d(%3d)", c[SPELLS], c[SPELLMAX])
		} else {
			lprintf("Spells:%3d(%2d) ", c[SPELLS], c[SPELLMAX])
		}
		lprintf(" AC: %-3d  WC: %-3d  Level", c[AC], c[WCLASS])
		if c[LEVEL] > 99 {
			lprintf("%3d", c[LEVEL])
		} else {
			lprintf(" %-2d", c[LEVEL])
		}
		lprintf(" Exp: %-9d %s\n", c[EXPERIENCE], class[c[LEVEL]-1])
		lprintf("HP: %3d(%3d) STR=%-2d INT=%-2d ",
			c[HP], c[HPMAX], c[STRENGTH]+c[STREXTRA], c[INTELLIGENCE])
		lprintf("WIS=%-2d CON=%-2d DEX=%-2d CHA=%-2d LV:",
			c[WISDOM], c[CONSTITUTION], c[DEXTERITY], c[CHARISMA])

		if level == 0 || wizard {
			c[TELEFLAG] = 0
		}
		if c[TELEFLAG] != 0 {
			lprcat(" ?")
		} else {
			lprcat(levelname[level])
		}
		lprintf("  Gold: %-6d", c[GOLD])
		always = true
		botside()
		c[TMP] = c[STRENGTH] + c[STREXTRA]
		for i := 0; i < 100; i++ {
			cbak[i] = c[i]
		}
		return
	}
	botsub(makecode(SPELLS, 8, 18), "%3d")
	if c[SPELLMAX] > 99 {
		botsub(makecode(SPELLMAX, 12, 18), "%3d)")
	} else {
		botsub(makecode(SPELLMAX, 12, 18), "%2d) ")
	}
	botsub(makecode(HP, 5, 19), "%3d")
	botsub(makecode(HPMAX, 9, 19), "%3d")
	botsub(makecode(AC, 21, 18), "%-3d")
	botsub(makecode(WCLASS, 30, 18), "%-3d")
	botsub(makecode(EXPERIENCE, 49, 18), "%-9d")
	if c[LEVEL] != cbak[LEVEL] {
		cursor(59, 18)
		lprcat(class[c[LEVEL]-1])
	}
	if c[LEVEL] > 99 {
		botsub(makecode(LEVEL, 40, 18), "%3d")
	} else {
		botsub(makecode(LEVEL, 40, 18), " %-2d")
	}
	c[TMP] = c[STRENGTH] + c[STREXTRA]
	botsub(makecode(TMP, 18, 19), "%-2d")
	botsub(makecode(INTELLIGENCE, 25, 19), "%-2d")
	botsub(makecode(WISDOM, 32, 19), "%-2d")
	botsub(makecode(CONSTITUTION, 39, 19), "%-2d")
	botsub(makecode(DEXTERITY, 46, 19), "%-2d")
	botsub(makecode(CHARISMA, 53, 19), "%-2d")
	if level != cbak[CAVELEVEL] || c[TELEFLAG] != cbak[TELEFLAG] {
		if level == 0 || wizard {
			c[TELEFLAG] = 0
		}
		cbak[TELEFLAG] = c[TELEFLAG]
		cbak[CAVELEVEL] = level
		cursor(59, 19)
		if c[TELEFLAG] != 0 {
			lprcat(" ?")
		} else {
			lprcat(levelname[level])
		}
	}
	botsub(makecode(GOLD, 69, 19), "%-6d")
	botside()
}

/*
	special subroutine to update only the gold number on the bottomlines
	called from ogold()
*/
func bottomgold() {
	botsub(makecode(GOLD, 69, 19), "%-6d")
	/* botsub(GOLD,"%-6d",69,19) */
}

/*
	special routine to update hp and level fields on bottom lines
	called in monster.c hitplayer() and spattack()
*/
func bot_hpx() {
	if c[EXPERIENCE] != cbak[EXPERIENCE] {
		recalc()
		bot_linex()
	} else {
		botsub(makecode(HP, 5, 19), "%3d")
	}
}

/*
	special routine to update number of spells called from regen()
*/
func bot_spellx() {
	botsub(makecode(SPELLS, 9, 18), "%2d")
}

/*
	common subroutine for a more economical bottomline()
*/
var bot_data = [...]struct {
	typ int
	str string
}{
	{STEALTH, "stealth"},
	{UNDEADPRO, "undead pro"},
	{SPIRITPRO, "spirit pro"},
	{CHARMCOUNT, "Charm"},
	{TIMESTOP, "Time Stop"},
	{HOLDMONST, "Hold Monst"},
	{GIANTSTR, "Giant Str"},
	{FIRERESISTANCE, "Fire Resit"},
	{DEXCOUNT, "Dexterity"},
	{STRCOUNT, "Strength"},
	{SCAREMONST, "Scare"},
	{HASTESELF, "Haste Self"},
	{CANCELLATION, "Cancel"},
	{INVISIBILITY, "Invisible"},
	{ALTPRO, "Protect 3"},
	{PROTECTIONTIME, "Protect 2"},
	{WTW, "Wall-Walk"},
}

func botside() {
	for i := 0; i < 17; i++ {
		idx := bot_data[i].typ
		if always || c[idx] != cbak[idx] {
			if always || cbak[idx] == 0 {
				if c[idx] != 0 {
					cursor(70, i+1)
					lprcat(bot_data[i].str)
				}
			} else if c[idx] == 0 {
				cursor(70, i+1)
				lprcat("          ")
			}
			cbak[idx] = c[idx]
		}
	}
	always = false
}

func botsub(idx int, str string) {
	y := idx & 0xff
	x := (idx >> 8) & 0xff
	idx >>= 16
	if c[idx] != cbak[idx] {
		cbak[idx] = c[idx]
		cursor(x, y)
		lprintf(str, c[idx])
	}
}

/*
 *	subroutine to draw only a section of the screen
 *	only the top section of the screen is updated.
 *	If entire lines are being drawn, then they will be cleared first.
 */
/* for limited screen drawing */
var d_xmin, d_xmax, d_ymin, d_ymax = 0, MAXX, 0, MAXY

func draws(xmin, xmax, ymin, ymax int) {
	if xmin == 0 && xmax == MAXX { /* clear section of screen as needed */
		if ymin == 0 {
			cl_up(79, ymax)
		} else {
			for i := ymin; i < ymin; i++ {
				cl_line(1, i+1)
			}
		}
		xmin = -1
	}
	d_xmin = xmin
	d_xmax = xmax
	d_ymin = ymin
	d_ymax = ymax /* for limited screen drawing */
	drawscreen()
	if xmin <= 0 && xmax == MAXX { /* draw stuff on right side of screen as needed */
		for i := ymin; i < ymax; i++ {
			idx := bot_data[i].typ
			if c[idx] != 0 {
				cursor(70, i+1)
				lprcat(bot_data[i].str)
			}
			cbak[idx] = c[idx]
		}
	}
}

/*
	drawscreen()

	subroutine to redraw the whole screen as the player knows it
*/
var screen [MAXX][MAXY]byte /* template for the screen */
var d_flag bool

func drawscreen() {
	if d_xmin == 0 && d_xmax == MAXX && d_ymin == 0 && d_ymax == MAXY {
		d_flag = true
		clear() /* clear the screen */
	} else {
		d_flag = true
		cursor(1, 1)
	}
	if d_xmin < 0 {
		d_xmin = 0 /* d_xmin=-1 means display all without bottomline */
	}

	for i := d_ymin; i < d_ymax; i++ {
		for j := d_xmin; j < d_xmax; j++ {
			if !know[j][i] {
				screen[j][i] = ' '
			} else if mitem[j][i] != 0 {
				screen[j][i] = monstnamelist[mitem[j][i]]
			} else if item[j][i] == OWALL {
				screen[j][i] = '#'
			} else {
				screen[j][i] = ' '
			}
		}
	}

	for i := d_ymin; i < d_ymax; i++ {
		j := d_xmin
		for screen[j][i] == ' ' && j < d_xmax {
			j++
		}
		/* was m=0 */
		var m int
		if j >= d_xmax {
			m = d_xmin /* don't search backwards if blank line */
		} else { /* search backwards for end of line */
			m = d_xmax - 1
			for screen[m][i] == ' ' && m > d_xmin {
				m--
			}
			if j <= m {
				cursor(j+1, i+1)
			} else {
				continue
			}
		}
		for j <= m {
			if j <= m-3 {
				var kk int
				for kk = j; kk <= j+3; kk++ {
					if screen[kk][i] != ' ' {
						kk = 1000
					}
				}
				if kk < 1000 {
					for screen[j][i] == ' ' && j <= m {
						j++
					}
					cursor(j+1, i+1)
				}
			}
			j++
			lprc(screen[j-1][i])
		}
	}
	setbold() /* print out only bold objects now */

	//int lastx, lasty;	/* variables used to optimize the object printing */
	for lastx, lasty, i := 127, 127, d_ymin; i < d_ymax; i++ {
		for j := d_xmin; j < d_xmax; j++ {
			kk := item[j][i]
			if kk != 0 {
				if kk != OWALL {
					if know[j][i] && mitem[j][i] == 0 {
						if objnamelist[kk] != ' ' {
							if lasty != i+1 || lastx != j {
								lastx = j + 1
								lasty = i + 1
								cursor(lastx, lasty)
							} else {
								lastx++
							}
							lprc(objnamelist[kk])
						}
					}
				}
			}
		}
	}

	resetbold()
	if d_flag {
		always = true
		botside()
		always = true
		bot_linex()
	}
	oldx = 99
	d_xmin, d_xmax, d_ymin, d_ymax = 0, MAXX, 0, MAXY /* for limited screen drawing */
}

/*
	showcell(x,y)

	subroutine to display a cell location on the screen
*/
func showcell(x, y int) {
	if c[BLINDCOUNT] != 0 {
		return /* see nothing if blind		 */
	}
	if c[AWARENESS] != 0 {
		minx = x - 3
		maxx = x + 3
		miny = y - 3
		maxy = y + 3
	} else {
		minx = x - 1
		maxx = x + 1
		miny = y - 1
		maxy = y + 1
	}

	if minx < 0 {
		minx = 0
	}
	if maxx > MAXX-1 {
		maxx = MAXX - 1
	}
	if miny < 0 {
		miny = 0
	}
	if maxy > MAXY-1 {
		maxy = MAXY - 1
	}

	for j := miny; j <= maxy; j++ {
		for mm := minx; mm <= maxx; mm++ {
			if !know[mm][j] {
				cursor(mm+1, j+1)
				x = maxx
				for know[x][j] {
					x--
				}
				for i := mm; i <= x; i++ {
					kk := mitem[i][j]
					if kk != 0 {
						lprc(monstnamelist[kk])
					} else {
						kk = item[i][j]
						switch kk {
						case OWALL, 0, OIVTELETRAP, OTRAPARROWIV, OIVDARTRAP, OIVTRAPDOOR:
							lprc(objnamelist[kk])

						default:
							setbold()
							lprc(objnamelist[kk])
							resetbold()
						}
					}
					know[i][j] = true
				}
				mm = maxx
			}
		}
	}
}

/*
	this routine shows only the spot that is given it.  the spaces around
	these coordinated are not shown
	used in godirect() in monster.c for missile weapons display
*/
func show1cell(x, y int) {
	if c[BLINDCOUNT] != 0 {
		return /* see nothing if blind		 */
	}
	cursor(x+1, y+1)
	k = mitem[x][y]
	if k != 0 {
		lprc(monstnamelist[k])
	} else {
		k = item[x][y]
		switch k {
		case OWALL, 0, OIVTELETRAP, OTRAPARROWIV, OIVDARTRAP, OIVTRAPDOOR:
			lprc(objnamelist[k])

		default:
			setbold()
			lprc(objnamelist[k])
			resetbold()
		}
	}
	know[x][y] = true /* we end up knowing about it */
}

/*
	showplayer()

	subroutine to show where the player is on the screen
	cursor values start from 1 up
*/
func showplayer() {
	cursor(playerx+1, playery+1)
	oldx = playerx
	oldy = playery
}

/*
	moveplayer(dir)

	subroutine to move the player from one room to another
	returns 0 if can't move in that direction or hit a monster or on an object
	else returns 1
	nomove is set to 1 to stop the next move (inadvertent monsters hitting
	players when walking into walls) if player walks off screen or into wall
*/
var diroffx = []int{0, 0, 1, 0, -1, 1, -1, 1, -1}
var diroffy = []int{0, 1, 0, -1, 0, -1, -1, 1, 1}

func moveplayer(dir int) int {
	/* from = present room #  direction =
	 * [1-north] [2-east] [3-south] [4-west]
	 * [5-northeast] [6-northwest] [7-southeast]
	 * [8-southwest] if direction=0, don't
	 * move--just show where he is */
	if c[CONFUSE] != 0 {
		if c[LEVEL] < rnd(30) {
			dir = rund(9) /* if confused any dir */
		}
	}
	kk := playerx + diroffx[dir]
	mm := playery + diroffy[dir]
	if kk < 0 || kk >= MAXX || mm < 0 || mm >= MAXY {
		nomove = 1
		yrepcount = 0
		return 0
	}
	i := item[kk][mm]
	j := mitem[kk][mm]
	if i == OWALL && c[WTW] == 0 {
		nomove = 1
		yrepcount = 0
		return 0
	} /* hit a wall	 */
	if kk == 33 && mm == MAXY-1 && level == 1 {
		newcavelevel(0)
		for kk = 0; kk < MAXX; kk++ {
			for mm = 0; mm < MAXY; mm++ {
				if item[kk][mm] == OENTRANCE {
					playerx = kk
					playery = mm
					positionplayer()
					drawscreen()
					return 0
				}
			}
		}
	}
	if j > 0 {
		hitmonster(kk, mm)
		yrepcount = 0
		return 0
	} /* hit a monster */
	lastpx = playerx
	lastpy = playery
	playerx = kk
	playery = mm
	if i != 0 && i != OTRAPARROWIV && i != OIVTELETRAP && i != OIVDARTRAP && i != OIVTRAPDOOR {
		yrepcount = 0
		return 0
	}
	return 1
}

/*
 *	function to show what magic items have been discovered thus far
 *	enter with -1 for just spells, anything else will give scrolls & potions
 */
var lincount, count int

func seemagic(arg int) {
	count, lincount = 0, 0
	nosignal = 1

	number := 0
	if arg == -1 { /* if display spells while casting one */
		for i := 0; i < SPNUM; i++ {
			if spelknow[i] {
				number++
			}
		}
		number = (number+2)/3 + 4 /* # lines needed to display */
		cl_up(79, number)
		cursor(1, 1)
	} else {
		resetscroll()
		clear()
	}

	lprcat("The magic spells you have discovered thus far:\n\n")
	for i := 0; i < SPNUM; i++ {
		if spelknow[i] {
			lprintf("%s %-20s ", spelcode[i], spelname[i])
			seepage()
		}
	}
	if arg == -1 {
		seepage()
		more()
		nosignal = 0
		draws(0, MAXX, 0, number)
		return
	}
	lincount += 3
	if count != 0 {
		count = 2
		seepage()
	}
	lprcat("\nThe magic scrolls you have found to date are:\n\n")
	count = 0
	for i := 0; i < MAXSCROLL; i++ {
		if scrollname[i] != "" {
			if scrollname[i][1] != ' ' {
				lprintf("%-26s", scrollname[i][1:])
				seepage()
			}
		}
	}
	lincount += 3
	if count != 0 {
		count = 2
		seepage()
	}
	lprcat("\nThe magic potions you have found to date are:\n\n")
	count = 0
	for i := 0; i < MAXPOTION; i++ {
		if potionname[i] != "" {
			if potionname[i][1] != ' ' {
				lprintf("%-26s", potionname[i][1:])
				seepage()
			}
		}
	}
	if lincount != 0 {
		more()
	}
	nosignal = 0
	setscroll()
	drawscreen()
}

/*
 *	subroutine to paginate the seemagic function
 */
func seepage() {
	count++
	if count == 3 {
		lincount++
		count = 0
		lprc('\n')
		if lincount > 17 {
			lincount = 0
			more()
			clear()
		}
	}
}
