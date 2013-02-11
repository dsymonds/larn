package main

/*
	lookforobject

	subroutine to look for an object and give the player his options
	if an object was found.
*/
func lookforobject() {
	if c[TIMESTOP] {
		return /* can't find objects if time is stopped	 */
	}
	i := item[playerx][playery]
	if i == 0 {
		return
	}
	showcell(playerx, playery)
	cursors()
	yrepcount = 0
	switch i {
	case OGOLDPILE, OMAXGOLD, OKGOLD, ODGOLD:
		lprcat("\n\nYou have found some gold!")
		ogold(i)

	case OPOTION:
		lprcat("\n\nYou have found a magic potion")
		i = iarg[playerx][playery]
		if potionname[i] != "" {
			lprintf(" of%s", potionname[i])
		}
		opotion(i)

	case OSCROLL:
		lprcat("\n\nYou have found a magic scroll")
		i = iarg[playerx][playery]
		if scrollname[i] != "" {
			lprintf(" of%s", scrollname[i])
		}
		oscroll(i)

	case OALTAR:
		if nearbymonst() {
			return
		}
		lprcat("\n\nThere is a Holy Altar here!")
		oaltar()

	case OBOOK:
		lprcat("\n\nYou have found a book.")
		obook()

	case OCOOKIE:
		lprcat("\n\nYou have found a fortune cookie.")
		ocookie()

	case OTHRONE:
		if nearbymonst() {
			return
		}
		lprintf("\n\nThere is %s here!", objectname[i])
		othrone(0)

	case OTHRONE2:
		if nearbymonst() {
			return
		}
		lprintf("\n\nThere is %s here!", objectname[i])
		othrone(1)

	case ODEADTHRONE:
		lprintf("\n\nThere is %s here!", objectname[i])
		odeadthrone()

	case OORB:
		lprcat("\n\nYou have found the Orb!!!!!")
		oorb()

	case OPIT:
		lprcat("\n\nYou're standing at the top of a pit.")
		opit()

	case OSTAIRSUP:
		lprcat("\n\nThere is a circular staircase here")
		ostairs(1) /* up */

	case OELEVATORUP:
		lprcat("\n\nYou feel heavy for a moment, but the feeling disappears")
		oelevator(1) /* up  */

	case OFOUNTAIN:
		if nearbymonst() {
			return
		}
		lprcat("\n\nThere is a fountain here")
		ofountain()

	case OSTATUE:
		if nearbymonst() {
			return
		}
		lprcat("\n\nYou are standing in front of a statue")
		ostatue()

	case OCHEST:
		lprcat("\n\nThere is a chest here")
		ochest()

	case OIVTELETRAP:
		if rnd(11) < 6 {
			return
		}
		item[playerx][playery] = OTELEPORTER
		know[playerx][playery] = true
		fallthrough

	case OTELEPORTER:
		lprcat("\nZaaaappp!  You've been teleported!\n")
		beep()
		nap(3000)
		oteleport(0)

	case OSCHOOL:
		if nearbymonst() {
			return
		}
		lprcat("\n\nYou have found the College of Larn.")
		lprcat("\nDo you (g) go inside, or (i) stay here? ")
		i = 0
		for i != 'g' && i != 'i' && i != '\033' {
			i = ttgetch()
		}
		if i == 'g' {
			oschool() /* the college of larn	 */
		} else {
			lprcat(" stay here")
		}

	case OMIRROR:
		if nearbymonst() {
			return
		}
		lprcat("\n\nThere is a mirror here")
		omirror()

	case OBANK2, OBANK:
		if nearbymonst() {
			return
		}
		if i == OBANK {
			lprcat("\n\nYou have found the bank of Larn.")
		} else {
			lprcat("\n\nYou have found a branch office of the bank of Larn.")
		}
		lprcat("\nDo you (g) go inside, or (i) stay here? ")
		j = 0
		for j != 'g' && j != 'i' && j != '\033' {
			j = ttgetch()
		}
		if j == 'g' {
			if i == OBANK {
				obank()
			} else {
				obank2() /* the bank of larn  */
			}
		} else {
			lprcat(" stay here")
		}

	case ODEADFOUNTAIN:
		if nearbymonst() {
			return
		}
		lprcat("\n\nThere is a dead fountain here")

	case ODNDSTORE:
		if nearbymonst() {
			return
		}
		lprcat("\n\nThere is a DND store here.")
		lprcat("\nDo you (g) go inside, or (i) stay here? ")
		i = 0
		for i != 'g' && i != 'i' && i != '\033' {
			i = ttgetch()
		}
		if i == 'g' {
			dndstore() /* the dnd adventurers store  */
		} else {
			lprcat(" stay here")
		}

	case OSTAIRSDOWN:
		lprcat("\n\nThere is a circular staircase here")
		ostairs(-1) /* down */

	case OELEVATORDOWN:
		lprcat("\n\nYou feel light for a moment, but the feeling disappears")
		oelevator(-1) /* down	 */

	case OOPENDOOR:
		lprintf("\n\nYou have found %s", objectname[i])
		lprcat("\nDo you (c) close it")
		iopts()
		i = 0
		for i != 'c' && i != 'i' && i != '\033' {
			i = ttgetch()
		}
		if i == '\033' || i == 'i' {
			ignore()
			break
		}
		lprcat("close")
		forget()
		item[playerx][playery] = OCLOSEDDOOR
		iarg[playerx][playery] = 0
		playerx = lastpx
		playery = lastpy

	case OCLOSEDDOOR:
		lprintf("\n\nYou have found %s", objectname[i])
		lprcat("\nDo you (o) try to open it")
		iopts()
		i = 0
		for i != 'o' && i != 'i' && i != '\033' {
			i = ttgetch()
		}
		if i == '\033' || i == 'i' {
			ignore()
			playerx = lastpx
			playery = lastpy
			break
		} else {
			lprcat("open")
			if rnd(11) < 7 {
				switch iarg[playerx][playery] {
				case 6:
					c[AGGRAVATE] += rnd(400)

				case 7:
					lprcat("\nYou are jolted by an electric shock ")
					lastnum = 274
					losehp(rnd(20))
					bottomline()

				case 8:
					loselevel()

				case 9:
					lprcat("\nYou suddenly feel weaker ")
					if c[STRENGTH] > 3 {
						c[STRENGTH]--
					}
					bottomline()
				}
				playerx = lastpx
				playery = lastpy
			} else {
				forget()
				item[playerx][playery] = OOPENDOOR
			}
		}

	case OENTRANCE:
		lprcat("\nYou have found ")
		lprcat(objectname[OENTRANCE])
		lprcat("\nDo you (g) go inside")
		iopts()
		i = 0
		for i != 'g' && i != 'i' && i != '\033' {
			i = ttgetch()
		}
		if i == 'g' {
			newcavelevel(1)
			playerx = 33
			playery = MAXY - 2
			item[33][MAXY-1], know[33][MAXY-1], mitem[33][MAXY-1] = 0, false, 0
			draws(0, MAXX, 0, MAXY)
			bot_linex()
			return
		}
		ignore()

	case OVOLDOWN:
		lprcat("\nYou have found ")
		lprcat(objectname[OVOLDOWN])
		lprcat("\nDo you (c) climb down")
		iopts()
		i = 0
		for i != 'c' && i != 'i' && i != '\033' {
			i = ttgetch()
		}
		if i == '\033' || i == 'i' {
			ignore()
			break
		}
		if level != 0 {
			lprcat("\nThe shaft only extends 5 feet downward!")
			return
		}
		if packweight() > 45+3*(c[STRENGTH]+c[STREXTRA]) {
			lprcat("\nYou slip and fall down the shaft")
			beep()
			lastnum = 275
			losehp(30 + rnd(20))
			bottomhp()
		} else {
			lprcat("climb down")
		}
		nap(3000)
		newcavelevel(MAXLEVEL)
		for i = 0; i < MAXY; i++ {
			for j = 0; j < MAXX; j++ { /* put player near volcano shaft */
				if item[j][i] == OVOLUP {
					playerx = j
					playery = i
					j = MAXX
					i = MAXY
					positionplayer()
				}
			}
		}
		draws(0, MAXX, 0, MAXY)
		bot_linex()
		return

	case OVOLUP:
		lprcat("\nYou have found ")
		lprcat(objectname[OVOLUP])
		lprcat("\nDo you (c) climb up")
		iopts()
		i = 0
		for i != 'c' && i != 'i' && i != '\033' {
			i = ttgetch()
		}
		if i == '\033' || i == 'i' {
			ignore()
			break
		}
		if level != 11 {
			lprcat("\nThe shaft only extends 8 feet upwards before you find a blockage!")
			return
		}
		if packweight() > 45+5*(c[STRENGTH]+c[STREXTRA]) {
			lprcat("\nYou slip and fall down the shaft")
			beep()
			lastnum = 275
			losehp(15 + rnd(20))
			bottomhp()
			return
		}
		lprcat("climb up")
		lflush()
		nap(3000)
		newcavelevel(0)
		for i = 0; i < MAXY; i++ {
			for j = 0; j < MAXX; j++ { /* put player near volcano shaft */
				if item[j][i] == OVOLDOWN {
					playerx = j
					playery = i
					j = MAXX
					i = MAXY
					positionplayer()
				}
			}
		}
		draws(0, MAXX, 0, MAXY)
		bot_linex()
		return

	case OTRAPARROWIV:
		if rnd(17) < 13 {
			return /* for an arrow trap */
		}
		item[playerx][playery] = OTRAPARROW
		know[playerx][playery] = false
		fallthrough

	case OTRAPARROW:
		lprcat("\nYou are hit by an arrow")
		beep() /* for an arrow trap */
		lastnum = 259
		losehp(rnd(10) + level)
		bottomhp()
		return

	case OIVDARTRAP:
		if rnd(17) < 13 {
			return /* for a dart trap */
		}
		item[playerx][playery] = ODARTRAP
		know[playerx][playery] = false
		fallthrough

	case ODARTRAP:
		lprcat("\nYou are hit by a dart")
		beep() /* for a dart trap */
		lastnum = 260
		losehp(rnd(5))
		c[STRENGTH]--
		if c[STRENGTH] < 3 {
			c[STRENGTH] = 3
		}
		bottomline()
		return

	case OIVTRAPDOOR:
		if rnd(17) < 13 {
			return /* for a trap door */
		}
		item[playerx][playery] = OTRAPDOOR
		know[playerx][playery] = true
		fallthrough

	case OTRAPDOOR:
		lastnum = 272 /* a trap door */
		if level == MAXLEVEL-1 || level == MAXLEVEL+MAXVLEVEL-1 {
			lprcat("\nYou fell through a bottomless trap door!")
			beep()
			nap(3000)
			died(271)
		}
		lprcat("\nYou fall through a trap door!")
		beep() /* for a trap door */
		losehp(rnd(5 + level))
		nap(2000)
		newcavelevel(level + 1)
		draws(0, MAXX, 0, MAXY)
		bot_linex()
		return

	case OTRADEPOST:
		if nearbymonst() {
			return
		}
		lprcat("\nYou have found the Larn trading Post.")
		lprcat("\nDo you (g) go inside, or (i) stay here? ")
		i = 0
		for i != 'g' && i != 'i' && i != '\033' {
			i = ttgetch()
		}
		if i == 'g' {
			otradepost()
		} else {
			lprcat("stay here")
		}
		return

	case OHOME:
		if nearbymonst() {
			return
		}
		lprcat("\nYou have found your way home.")
		lprcat("\nDo you (g) go inside, or (i) stay here? ")
		i = 0
		for i != 'g' && i != 'i' && i != '\033' {
			i = ttgetch()
		}
		if i == 'g' {
			ohome()
		} else {
			lprcat("stay here")
		}
		return

	case OWALL:

	case OANNIHILATION:
		died(283)
		return /* annihilated by sphere of annihilation */

	case OLRS:
		if nearbymonst() {
			return
		}
		lprcat("\n\nThere is an LRS office here.")
		lprcat("\nDo you (g) go inside, or (i) stay here? ")
		i = 0
		for i != 'g' && i != 'i' && i != '\033' {
			i = ttgetch()
		}
		if i == 'g' {
			olrs() /* the larn revenue service */
		} else {
			lprcat(" stay here")
		}
		break

	default:
		finditem(i)
		break
	}
}

/*
	function to say what object we found and ask if player wants to take it
*/
func finditem(theitem int) {
	lprintf("\n\nYou have found %s ", objectname[theitem])
	tmp := iarg[playerx][playery]
	switch theitem {
	case ODIAMOND, ORUBY, OEMERALD, OSAPPHIRE, OSPIRITSCARAB, OORBOFDRAGON, OCUBEofUNDEAD, ONOTHEFT:

	default:
		if tmp > 0 {
			lprintf("+ %d", tmp)
		} else if tmp < 0 {
			lprintf(" %d", tmp)
		}
	}
	lprcat("\nDo you want to (t) take it")
	iopts()
	i := 0
	for i != 't' && i != 'i' && i != '\033' {
		i = ttgetch()
	}
	if i == 't' {
		lprcat("take")
		if take(theitem, tmp) == 0 {
			forget()
		}
		return
	}
	ignore()
}

/*
	subroutine to process the stair cases
	if dir > 0 the up else down
*/
func ostairs(dir int) {
	lprcat("\nDo you (s) stay here  ")
	if dir > 0 {
		lprcat("(u) go up  ")
	} else {
		lprcat("(d) go down  ")
	}
	lprcat("or (f) kick stairs? ")

	for {
		switch ttgetch() {
		case '\033', 's', 'i':
			lprcat("stay here")
			return

		case 'f':
			lprcat("kick stairs")
			if rnd(2) == 1 {
				lprcat("\nI hope you feel better.  Showing anger rids you of frustration.")
			} else {
				k := rnd((level + 1) << 1)
				lprintf("\nYou hurt your foot dumb dumb!  You suffer %d hit points", k)
				lastnum = 276
				losehp(k)
				bottomline()
			}
			return

		case 'u':
			lprcat("go up")
			if dir < 0 {
				lprcat("\nThe stairs don't go up!")
			} else if level >= 2 && level != 11 {
				newcavelevel(level - 1)
				draws(0, MAXX, 0, MAXY)
				bot_linex()
			} else {
				lprcat("\nThe stairs lead to a dead end!")
			}
			return

		case 'd':
			lprcat("go down")
			if dir > 0 {
				lprcat("\nThe stairs don't go down!")
			} else if level != 0 && level != 10 && level != 13 {
				newcavelevel(level + 1)
				draws(0, MAXX, 0, MAXY)
				bot_linex()
			} else {
				lprcat("\nThe stairs lead to a dead end!")
			}
			return
		}
	}
}

/*
	subroutine to handle a teleport trap +/- 1 level maximum
*/
func oteleport(err int) {
	if err != 0 {
		if rnd(151) < 3 {
			died(264) /* stuck in a rock */
		}
	}
	c[TELEFLAG] = 1 /* show ?? on bottomline if been teleported	 */
	var tmp int
	if level == 0 {
		tmp = 0
	} else if level < MAXLEVEL {
		tmp = rnd(5) + level - 3
		if tmp >= MAXLEVEL {
			tmp = MAXLEVEL - 1
		}
		if tmp < 1 {
			tmp = 1
		}
	} else {
		tmp = rnd(3) + level - 2
		if tmp >= MAXLEVEL+MAXVLEVEL {
			tmp = MAXLEVEL + MAXVLEVEL - 1
		}
		if tmp < MAXLEVEL {
			tmp = MAXLEVEL
		}
	}
	playerx = rnd(MAXX - 2)
	playery = rnd(MAXY - 2)
	if level != tmp {
		newcavelevel(tmp)
	}
	positionplayer()
	draws(0, MAXX, 0, MAXY)
	bot_linex()
}

/*
	function to process a potion
*/
func opotion(pot int) {
	lprcat("\nDo you (d) drink it, (t) take it")
	iopts()
	for {
		switch ttgetch() {
		case '\033', 'i':
			ignore()
			return

		case 'd':
			lprcat("drink\n")
			forget() /* destroy potion	 */
			quaffpotion(pot)
			return

		case 't':
			lprcat("take\n")
			if take(OPOTION, pot) == 0 {
				forget()
			}
			return
		}
	}
}

/*
	function to drink a potion
*/
func quaffpotion(pot int) {
	if pot < 0 || pot >= MAXPOTION {
		return /* check for within bounds */
	}
	potionname[pot] = potionhide[pot]
	switch pot {
	case 9:
		lprcat("\nYou feel greedy . . .")
		nap(2000)
		for i := 0; i < MAXY; i++ {
			for j := 0; j < MAXX; j++ {
				if item[j][i] == OGOLDPILE || item[j][i] == OMAXGOLD {
					know[j][i] = true
					show1cell(j, i)
				}
			}
		}
		showplayer()
		return

	case 19:
		lprcat("\nYou feel greedy . . .")
		nap(2000)
		for i := 0; i < MAXY; i++ {
			for j := 0; j < MAXX; j++ {
				k := item[j][i]
				if k == ODIAMOND || k == ORUBY || k == OEMERALD || k == OMAXGOLD || k == OSAPPHIRE || k == OLARNEYE || k == OGOLDPILE {
					know[j][i] = true
					show1cell(j, i)
				}
			}
		}
		showplayer()
		return

	case 20:
		c[HP] = c[HPMAX] /* instant healing */

	case 1:
		lprcat("\nYou feel better")
		if c[HP] == c[HPMAX] {
			raisemhp(1)
		} else {
			c[HP] += rnd(20) + 20 + c[LEVEL]
			if c[HP] > c[HPMAX] {
				c[HP] = c[HPMAX]
			}
		}

	case 2:
		lprcat("\nSuddenly, you feel much more skillful!")
		raiselevel()
		raisemhp(1)
		return

	case 3:
		lprcat("\nYou feel strange for a moment")
		c[rund(6)]++

	case 4:
		lprcat("\nYou feel more self confident!")
		c[WISDOM] += rnd(2)

	case 5:
		lprcat("\nWow!  You feel great!")
		if c[STRENGTH] < 12 {
			c[STRENGTH] = 12
		} else {
			c[STRENGTH]++
		}

	case 6:
		lprcat("\nYour charm went up by one!")
		c[CHARISMA]++

	case 8:
		lprcat("\nYour intelligence went up by one!")
		c[INTELLIGENCE]++

	case 10:
		for i := 0; i < MAXY; i++ {
			for j := 0; j < MAXX; j++ {
				if mitem[j][i] {
					know[j][i] = true
					show1cell(j, i)
				}
			}
		}
		/* monster detection	 */
		return

	case 12:
		lprcat("\nThis potion has no taste to it")
		return

	case 15:
		lprcat("\nWOW!!!  You feel Super-fantastic!!!")
		if c[HERO] == 0 {
			for i := 0; i < 6; i++ {
				c[i] += 11
			}
		}
		c[HERO] += 250

	case 16:
		lprcat("\nYou have a greater intestinal constitude!")
		c[CONSTITUTION]++

	case 17:
		lprcat("\nYou now have incredibly bulging muscles!!!")
		if c[GIANTSTR] == 0 {
			c[STREXTRA] += 21
		}
		c[GIANTSTR] += 700

	case 18:
		lprcat("\nYou feel a chill run up your spine!")
		c[FIRERESISTANCE] += 1000

	case 0:
		lprcat("\nYou fall asleep. . .")
		i := rnd(11) - (c[CONSTITUTION] >> 2) + 2
		for {
			i--
			if i <= 0 {
				break
			}
			parse2()
			nap(1000)
		}
		cursors()
		lprcat("\nYou woke up!")
		return

	case 7:
		lprcat("\nYou become dizzy!")
		c[STRENGTH]--
		if c[STRENGTH] < 3 {
			c[STRENGTH] = 3
		}

	case 11:
		lprcat("\nYou stagger for a moment . .")
		for i := 0; i < MAXY; i++ {
			for j := 0; j < MAXX; j++ {
				know[j][i] = false
			}
		}
		nap(2000)
		draws(0, MAXX, 0, MAXY) /* potion of forgetfulness */
		return

	case 13:
		lprcat("\nYou can't see anything!") /* blindness */
		c[BLINDCOUNT] += 500
		return

	case 14:
		lprcat("\nYou feel confused")
		c[CONFUSE] += 20 + rnd(9)
		return

	case 21:
		lprcat("\nYou don't seem to be affected")
		return /* cure dianthroritis */

	case 22:
		lprcat("\nYou feel a sickness engulf you") /* poison */
		c[HALFDAM] += 200 + rnd(200)
		return

	case 23:
		lprcat("\nYou feel your vision sharpen") /* see invisible */
		c[SEEINVISIBLE] += rnd(1000) + 400
		monstnamelist[INVISIBLESTALKER] = 'I'
		return
	}
	bottomline() /* show new stats		 */
	return
}

/*
	function to process a magic scroll
*/
func oscroll(typ int) {
	lprcat("\nDo you ")
	if c[BLINDCOUNT] == 0 {
		lprcat("(r) read it, ")
	}
	lprcat("(t) take it")
	iopts()
	for {
		switch ttgetch() {
		case '\033', 'i':
			ignore()
			return

		case 'r':
			if c[BLINDCOUNT] {
				break
			}
			lprcat("read")
			forget()
			if typ == 2 || typ == 15 {
				show1cell(playerx, playery)
				cursors()
			}
			/* destroy it	 */
			read_scroll(typ)
			return

		case 't':
			lprcat("take")
			if take(OSCROLL, typ) == 0 {
				forget() /* destroy it	 */
			}
			return
		}
	}
}

/*
	data for the function to read a scroll
*/
var curse = [...]byte{
	BLINDCOUNT, CONFUSE, AGGRAVATE, HASTEMONST, ITCHING,
	LAUGHING, DRAINSTRENGTH, CLUMSINESS, INFEEBLEMENT, HALFDAM,
}

var exten = [...]byte{
	PROTECTIONTIME, DEXCOUNT, STRCOUNT, CHARMCOUNT, INVISIBILITY,
	CANCELLATION, HASTESELF, GLOBE, SCAREMONST, HOLDMONST, TIMESTOP,
}

var time_change = [...]byte{
	HASTESELF, HERO, ALTPRO, PROTECTIONTIME, DEXCOUNT, STRCOUNT,
	GIANTSTR, CHARMCOUNT, INVISIBILITY, CANCELLATION, HASTESELF,
	AGGRAVATE, SCAREMONST, STEALTH, AWARENESS, HOLDMONST,
	HASTEMONST, FIRERESISTANCE, GLOBE, SPIRITPRO, UNDEADPRO,
	HALFDAM, SEEINVISIBLE, ITCHING, CLUMSINESS, WTW,
}

/*
 *	function to adjust time when time warping and taking courses in school
 */
func adjusttime(tim int32) {
	for j := 0; j < 26; j++ { /* adjust time related parameters */
		if c[time_change[j]] {
			c[time_change[j]] -= tim
			if c[time_change[j]] < 1 {
				c[time_change[j]] = 1
			}
		}
	}
	regen()
}

/*
	function to read a scroll
*/
func read_scroll(typ int) {
	if typ < 0 || typ >= MAXSCROLL {
		return /* be sure we are within bounds */
	}
	scrollname[typ] = scrollhide[typ]
	switch typ {
	case 0:
		lprcat("\nYour armor glows for a moment")
		enchantarmor()
		return

	case 1:
		lprcat("\nYour weapon glows for a moment")
		enchweapon()
		return /* enchant weapon */

	case 2:
		lprcat("\nYou have been granted enlightenment!")
		yh := min(playery+7, MAXY)
		xh := min(playerx+25, MAXX)
		yl := max(playery-7, 0)
		xl := max(playerx-25, 0)
		for i := yl; i < yh; i++ {
			for j := xl; j < xh; j++ {
				know[j][i] = true
			}
		}
		nap(2000)
		draws(xl, xh, yl, yh)
		return

	case 3:
		lprcat("\nThis scroll seems to be blank")
		return

	case 4:
		createmonster(makemonst(level + 1))
		return /* this one creates a monster  */

	case 5:
		something(level) /* create artifact		 */
		return

	case 6:
		c[AGGRAVATE] += 800
		return /* aggravate monsters */

	case 7:
		i := rnd(1000) - 850
		gltime += i /* time warp */
		if i >= 0 {
			lprintf("\nYou went forward in time by %d mobuls", (i+99)/100)
		} else {
			lprintf("\nYou went backward in time by %d mobuls", -(i+99)/100)
		}
		adjusttime(i) /* adjust time for time warping */
		return

	case 8:
		oteleport(0)
		return /* teleportation */

	case 9:
		c[AWARENESS] += 1800
		return /* expanded awareness	 */

	case 10:
		c[HASTEMONST] += rnd(55) + 12
		return /* haste monster */

	case 11:
		for i := 0; i < MAXY; i++ {
			for j := 0; j < MAXX; j++ {
				if mitem[j][i] {
					hitp[j][i] = monster[mitem[j][i]].hitpoints
				}
			}
		}
		return /* monster healing */
	case 12:
		c[SPIRITPRO] += 300 + rnd(200)
		bottomline()
		return /* spirit protection */

	case 13:
		c[UNDEADPRO] += 300 + rnd(200)
		bottomline()
		return /* undead protection */

	case 14:
		c[STEALTH] += 250 + rnd(250)
		bottomline()
		return /* stealth */

	case 15:
		lprcat("\nYou have been granted enlightenment!") /* magic mapping */
		for i := 0; i < MAXY; i++ {
			for j := 0; j < MAXX; j++ {
				know[j][i] = true
			}
		}
		nap(2000)
		draws(0, MAXX, 0, MAXY)
		return

	case 16:
		c[HOLDMONST] += 30
		bottomline()
		return /* hold monster */

	case 17:
		for i = 0; i < 26; i++ { /* gem perfection */
			switch iven[i] {
			case ODIAMOND, ORUBY, OEMERALD, OSAPPHIRE:
				j := ivenarg[i]
				j &= 255
				j <<= 1
				if j > 255 {
					j = 255 /* double value */
				}
				ivenarg[i] = j
			}
		}

	case 18:
		for i := 0; i < 11; i++ {
			c[exten[i]] <<= 1 /* spell extension */
		}

	case 19:
		for i := 0; i < 26; i++ { /* identify */
			if iven[i] == OPOTION {
				potionname[ivenarg[i]] = potionhide[ivenarg[i]]
			}
			if iven[i] == OSCROLL {
				scrollname[ivenarg[i]] = scrollhide[ivenarg[i]]
			}
		}

	case 20:
		for i := 0; i < 10; i++ { /* remove curse */
			if c[curse[i]] {
				c[curse[i]] = 1
			}
		}

	case 21:
		annihilate()
		break /* scroll of annihilation */

	case 22:
		godirect(22, 150, "The ray hits the %s", 0, ' ') /* pulverization */
	case 23:
		c[LIFEPROT]++
		break /* life protection */
	}
}

func oorb() {
}

func opit() {
	if rnd(101) < 81 {
		if rnd(70) > 9*c[DEXTERITY]-packweight() || rnd(101) < 5 {
			if level == MAXLEVEL-1 {
				obottomless()
			} else if level == MAXLEVEL+MAXVLEVEL-1 {
				obottomless()
			} else {
				var i int
				if rnd(101) < 20 {
					i = 0
					lprcat("\nYou fell into a pit!  Your fall is cushioned by an unknown force\n")
				} else {
					i = rnd(level*3 + 3)
					lprintf("\nYou fell into a pit!  You suffer %d hit points damage", i)
					lastnum = 261 /* if he dies scoreboard will say so */
				}
				losehp(i)
				nap(2000)
				newcavelevel(level + 1)
				draws(0, MAXX, 0, MAXY)
			}
		}
	}
}

func obottomless() {
	lprcat("\nYou fell into a bottomless pit!")
	beep()
	nap(3000)
	died(262)
}

func oelevator(dir int) {
}

func ostatue() {
}

func omirror() {
}

func obook() {
	lprcat("\nDo you ")
	if c[BLINDCOUNT] == 0 {
		lprcat("(r) read it, ")
	}
	lprcat("(t) take it")
	iopts()
	for {
		switch ttgetch() {
		case '\033', 'i':
			ignore()
			return

		case 'r':
			if c[BLINDCOUNT] {
				break
			}
			lprcat("read")
			/* no more book	 */ readbook(iarg[playerx][playery])
			forget()
			return

		case 't':
			lprcat("take")
			if take(OBOOK, iarg[playerx][playery]) == 0 {
				forget() /* no more book	 */
			}
			return
		}
	}
}

/*
	function to read a book
*/
func readbook(lev int) {
	var i, tmp int
	if lev <= 3 {
		tmp = splev[lev]
		i = rund(icond(tmp, tmp, 1))
	} else {
		tmp = splev[lev] - 9
		i = rnd(icond(tmp, tmp, 1)) + 9
	}
	spelknow[i] = true
	lprintf("\nSpell \"%s\":  %s\n%s", spelcode[i], spelname[i], speldescript[i])
	if rnd(10) == 4 {
		lprcat("\nYour int went up by one!")
		c[INTELLIGENCE]++
		bottomline()
	}
}

func ocookie() {
	lprcat("\nDo you (e) eat it, (t) take it")
	iopts()
	for {
		switch ttgetch() {
		case '\033', 'i':
			ignore()
			return

		case 'e':
			lprcat("eat\nThe cookie tasted good.")
			forget() /* no more cookie	 */
			if c[BLINDCOUNT] {
				return
			}
			p := fortune()
			if p == "" {
				return
			}
			lprcat("  A message inside the cookie reads:\n")
			lprcat(p)
			return

		case 't':
			lprcat("take")
			if take(OCOOKIE, 0) == 0 {
				forget() /* no more book	 */
			}
			return
		}
	}
}

/*
 * routine to pick up some gold -- if arg==OMAXGOLD then the pile is worth
 * 100* the argument
 */
func ogold(arg int) {
	i := iarg[playerx][playery]
	if arg == OMAXGOLD {
		i *= 100
	} else if arg == OKGOLD {
		i *= 1000
	} else if arg == ODGOLD {
		i *= 10
	}
	lprintf("\nIt is worth %d!", i)
	c[GOLD] += i
	bottomgold()
	item[playerx][playery], know[playerx][playery] = 0, false /* destroy gold	 */
}

func ohome() {
	nosignal = 1 /* disable signals */
	for i := 0; i < 26; i++ {
		if iven[i] == OPOTION {
			if ivenarg[i] == 21 {
				iven[i] = 0 /* remove the potion of cure dianthroritis from inventory */
				clear()
				lprcat("Congratulations.  You found a potion of cure dianthroritis.\n")
				lprcat("\nFrankly, No one thought you could do it.  Boy!  Did you surprise them!\n")
				if gltime > TIMELIMIT {
					lprcat("\nThe doctor has the sad duty to inform you that your daughter died!\n")
					lprcat("You didn't make it in time.  In your agony, you kill the doctor,\nyour wife, and yourself!  Too bad!\n")
					nap(5000)
					died(269)
				} else {
					lprcat("\nThe doctor is now administering the potion, and in a few moments\n")
					lprcat("Your daughter should be well on her way to recovery.\n")
					nap(6000)
					lprcat("\nThe potion is")
					nap(3000)
					lprcat(" working!  The doctor thinks that\n")
					lprcat("your daughter will recover in a few days.  Congratulations!\n")
					beep()
					nap(5000)
					died(263)
				}
			}
		}
	}
	for {
		clear()
		lprintf("Welcome home %s.  Latest word from the doctor is not good.\n", logname)

		if gltime > TIMELIMIT {
			lprcat("\nThe doctor has the sad duty to inform you that your daughter died!\n")
			lprcat("You didn't make it in time.  In your agony, you kill the doctor,\nyour wife, and yourself!  Too bad!\n")
			nap(5000)
			died(269)
		}
		lprcat("\nThe diagnosis is confirmed as dianthroritis.  He guesses that\n")
		lprintf("your daughter has only %d mobuls left in this world.  It's up to you,\n", (TIMELIMIT-gltime+99)/100)
		lprintf("%s, to find the only hope for your daughter, the very rare\n", logname)
		lprcat("potion of cure dianthroritis.  It is rumored that only deep in the\n")
		lprcat("depths of the caves can this potion be found.\n\n\n")
		lprcat("\n     ----- press ")
		standout("return")
		lprcat(" to continue, ")
		standout("escape")
		lprcat(" to leave ----- ")
		i := ttgetch()
		for i != '\033' && i != '\n' {
			i = ttgetch()
		}
		if i == '\033' {
			drawscreen()
			nosignal = 0 /* enable signals */
			return
		}
	}
}

/* routine to save program space	 */
func iopts() {
	lprcat(", or (i) ignore it? ")
}

func ignore() {
	lprcat("ignore\n")
}
