package main

/*
 * raiselevel()		subroutine to raise the player one level
 * loselevel()		subroutine to lower the player by one level
 * raiseexperience(x)	subroutine to increase experience points
 * loseexperience(x)	subroutine to lose experience points
 * losehp(x)		subroutine to remove hit points from the player
 * losemhp(x)		subroutine to remove max # hit points from the player
 * raisehp(x) 		subroutine to gain hit points
 * raisemhp(x)		subroutine to gain maximum hit points
 * losemspells(x)	subroutine to lose maximum spells
 * raisemspells(x)	subroutine to gain maximum spells
 * makemonst(lev)	function to return monster number for a randomly
 *			selected monster
 * positionplayer()	function to be sure player is not in a wall
 * recalc()		function to recalculate the armor class of the player
 * quit()		subroutine to ask if the player really wants to quit
 */

/*
	raiselevel()

	subroutine to raise the player one level
	uses the skill[] array to find level boundarys
	uses c[EXPERIENCE]  c[LEVEL]
*/
func raiselevel() {
	if c[LEVEL] < MAXPLEVEL {
		raiseexperience(skill[c[LEVEL]] - c[EXPERIENCE])
	}
}

/*
    loselevel()

	subroutine to lower the players character level by one
*/
func loselevel() {
	if c[LEVEL] > 1 {
		loseexperience(c[EXPERIENCE] - skill[c[LEVEL]-1] + 1)
	}
}

/*
	raiseexperience(x)

	subroutine to increase experience points
*/
func raiseexperience(x int) {
	i := c[LEVEL]
	c[EXPERIENCE] += x
	for c[EXPERIENCE] >= skill[c[LEVEL]] && c[LEVEL] < MAXPLEVEL {
		tmp := (c[CONSTITUTION] - c[HARDGAME]) >> 1
		c[LEVEL]++
		raisemhp(rnd(3) + rnd(icond(tmp > 0, tmp, 1)))
		raisemspells(rund(3))
		if c[LEVEL] < 7-c[HARDGAME] {
			raisemhp(c[CONSTITUTION] >> 2)
		}
	}
	if c[LEVEL] != i {
		cursors()
		beep()
		lprintf("\nWelcome to level %d", c[LEVEL]) /* if we changed levels	 */
	}
	bottomline()
}

/*
	loseexperience(x)

	subroutine to lose experience points
*/
func loseexperience(x int) {
	i := c[LEVEL]
	c[EXPERIENCE] -= x
	if c[EXPERIENCE] < 0 {
		c[EXPERIENCE] = 0
	}
	for c[EXPERIENCE] < skill[c[LEVEL]-1] {
		c[LEVEL]--
		if c[LEVEL] <= 1 {
			c[LEVEL] = 1 /* down one level		 */
		}
		tmp := (c[CONSTITUTION] - c[HARDGAME]) >> 1 /* lose hpoints */
		losemhp(rnd(icond(tmp > 0, tmp, 1)))        /* lose hpoints */
		if c[LEVEL] < 7-c[HARDGAME] {
			losemhp(c[CONSTITUTION] >> 2)
		}
		losemspells(rund(3)) /* lose spells		 */
	}
	if i != c[LEVEL] {
		cursors()
		beep()
		lprintf("\nYou went down to level %d!", c[LEVEL])
	}
	bottomline()
}

/*
	losehp(x)
	losemhp(x)

	subroutine to remove hit points from the player
	warning -- will kill player if hp goes to zero
*/
func losehp(x int) {
	c[HP] -= x
	if c[HP] <= 0 {
		beep()
		lprcat("\n")
		nap(3000)
		died(lastnum)
	}
}

func losemhp(x int) {
	c[HP] -= x
	if c[HP] < 1 {
		c[HP] = 1
	}
	c[HPMAX] -= x
	if c[HPMAX] < 1 {
		c[HPMAX] = 1
	}
}

/*
	raisehp(x)
	raisemhp(x)

	subroutine to gain maximum hit points
*/
func raisehp(x int) {
	c[HP] += x
	if c[HP] > c[HPMAX] {
		c[HP] = c[HPMAX]
	}
}

func raisemhp(x int) {
	c[HPMAX] += x
	c[HP] += x
}

/*
	raisemspells(x)

	subroutine to gain maximum spells
*/
func raisemspells(x int) {
	c[SPELLMAX] += x
	c[SPELLS] += x
}

/*
	losemspells(x)

	subroutine to lose maximum spells
*/
func losemspells(x int) {
	c[SPELLMAX] -= x
	if c[SPELLMAX] < 0 {
		c[SPELLMAX] = 0
	}
	c[SPELLS] -= x
	if c[SPELLS] < 0 {
		c[SPELLS] = 0
	}
}

/*
	makemonst(lev)
		int lev;

	function to return monster number for a randomly selected monster
		for the given cave level
*/
func makemonst(lev int) int {
	if lev < 1 {
		lev = 1
	}
	if lev > 12 {
		lev = 12
	}
	tmp := WATERLORD
	if lev < 5 {
		for tmp == WATERLORD {
			x := monstlevel[lev-1]
			tmp = rnd(icond(x != 0, x, 1))
		}
	} else {
		for tmp == WATERLORD {
			x := monstlevel[lev-1] - monstlevel[lev-4]
			tmp = rnd(icond(x != 0, x, 1)) + monstlevel[lev-4]
		}
	}

	for monster[tmp].genocided && tmp < MAXMONST {
		tmp++ /* genocided? */
	}
	return tmp
}

/*
	positionplayer()

	function to be sure player is not in a wall
*/
func positionplayer() {
	try := 2
	for (item[playerx][playery] != 0 || mitem[playerx][playery] != 0) && try != 0 {
		playerx++
		if playerx >= MAXX-1 {
			playerx = 1
			playery++
			if playery >= MAXY-1 {
				playery = 1
				try--
			}
		}
	}
	if try == 0 {
		lprcat("Failure in positionplayer\n")
	}
}

/*
	recalc()	function to recalculate the armor class of the player
*/
func recalc() {
	c[AC] = c[MOREDEFENSES]
	if c[WEAR] >= 0 {
		switch iven[c[WEAR]] {
		case OSHIELD:
			c[AC] += 2 + ivenarg[c[WEAR]]
		case OLEATHER:
			c[AC] += 2 + ivenarg[c[WEAR]]
		case OSTUDLEATHER:
			c[AC] += 3 + ivenarg[c[WEAR]]
		case ORING:
			c[AC] += 5 + ivenarg[c[WEAR]]
		case OCHAIN:
			c[AC] += 6 + ivenarg[c[WEAR]]
		case OSPLINT:
			c[AC] += 7 + ivenarg[c[WEAR]]
		case OPLATE:
			c[AC] += 9 + ivenarg[c[WEAR]]
		case OPLATEARMOR:
			c[AC] += 10 + ivenarg[c[WEAR]]
		case OSSPLATE:
			c[AC] += 12 + ivenarg[c[WEAR]]
		}
	}

	if c[SHIELD] >= 0 {
		if iven[c[SHIELD]] == OSHIELD {
			c[AC] += 2 + ivenarg[c[SHIELD]]
		}
	}
	if c[WIELD] < 0 {
		c[WCLASS] = 0
	} else {
		i := ivenarg[c[WIELD]]
		switch iven[c[WIELD]] {
		case ODAGGER:
			c[WCLASS] = 3 + i
		case OBELT:
			c[WCLASS] = 7 + i
		case OSHIELD:
			c[WCLASS] = 8 + i
		case OSPEAR:
			c[WCLASS] = 10 + i
		case OFLAIL:
			c[WCLASS] = 14 + i
		case OBATTLEAXE:
			c[WCLASS] = 17 + i
		case OLANCE:
			c[WCLASS] = 19 + i
		case OLONGSWORD:
			c[WCLASS] = 22 + i
		case O2SWORD:
			c[WCLASS] = 26 + i
		case OSWORD:
			c[WCLASS] = 32 + i
		case OSWORDofSLASHING:
			c[WCLASS] = 30 + i
		case OHAMMER:
			c[WCLASS] = 35 + i
		default:
			c[WCLASS] = 0
		}
	}
	c[WCLASS] += c[MOREDAM]

	/* now for regeneration abilities based on rings	 */
	c[REGEN] = 1
	c[ENERGY] = 0
	j := 0
	for k := 25; k > 0; k-- {
		if iven[k] != 0 {
			j = k
			k = 0
		}
	}
	for i := 0; i <= j; i++ {
		switch iven[i] {
		case OPROTRING:
			c[AC] += ivenarg[i] + 1
		case ODAMRING:
			c[WCLASS] += ivenarg[i] + 1
		case OBELT:
			c[WCLASS] += (ivenarg[i] << 1) + 2

		case OREGENRING:
			c[REGEN] += ivenarg[i] + 1
		case ORINGOFEXTRA:
			c[REGEN] += 5 * (ivenarg[i] + 1)
		case OENERGYRING:
			c[ENERGY] += ivenarg[i] + 1
		}
	}
}

/*
	quit()

	subroutine to ask if the player really wants to quit
*/
func quit() {
	cursors()
	lastmonst = ""
	lprcat("\n\nDo you really want to quit?")
	for {
		i := ttgetch()
		if i == 'y' {
			died(300)
			return
		}
		if i == 'n' || i == '\033' {
			lprcat(" no")
			lflush()
			return
		}
		lprcat("\n")
		setbold()
		lprcat("Yes")
		resetbold()
		lprcat(" or ")
		setbold()
		lprcat("No")
		resetbold()
		lprcat(" please?   Do you want to quit? ")
	}
}

/*
	function to ask --more-- then the user must enter a space
*/
func more() {
	lprcat("\n  --- press ")
	standout("space")
	lprcat(" to continue --- ")
	for ttgetch() != ' ' {
	}
}

/*
	function to put something in the players inventory
	returns 0 if success, 1 if a failure
*/
func take(theitem, arg int) int {
	/* cursors(); */
	limit := 15 + (c[LEVEL] >> 1)
	if limit > 26 {
		limit = 26
	}
	for i := 0; i < limit; i++ {
		if iven[i] == 0 {
			iven[i] = theitem
			ivenarg[i] = arg
			limit = 0
			switch theitem {
			case OPROTRING, ODAMRING, OBELT:
				limit = 1
			case ODEXRING:
				c[DEXTERITY] += ivenarg[i] + 1
				limit = 1
			case OSTRRING:
				c[STREXTRA] += ivenarg[i] + 1
				limit = 1
			case OCLEVERRING:
				c[INTELLIGENCE] += ivenarg[i] + 1
				limit = 1
			case OHAMMER:
				c[DEXTERITY] += 10
				c[STREXTRA] += 10
				c[INTELLIGENCE] -= 10
				limit = 1

			case OORBOFDRAGON:
				c[SLAYING]++
			case OSPIRITSCARAB:
				c[NEGATESPIRIT]++
			case OCUBEofUNDEAD:
				c[CUBEofUNDEAD]++
			case ONOTHEFT:
				c[NOTHEFT]++
			case OSWORDofSLASHING:
				c[DEXTERITY] += 5
				limit = 1
			}
			lprcat("\nYou pick up:")
			srcount = 0
			show3(i)
			if limit != 0 {
				bottomline()
			}
			return 0
		}
	}
	lprcat("\nYou can't carry anything else")
	return 1
}

/*
	subroutine to drop an object
	returns 1 if something there already else 0
*/
func drop_object(k int) int {
	if k < 0 || k > 25 {
		return 0
	}
	theitem := iven[k]
	cursors()
	if theitem == 0 {
		lprintf("\nYou don't have item %c! ", k+'a')
		return 1
	}
	if item[playerx][playery] != 0 {
		beep()
		lprcat("\nThere's something here already")
		return 1
	}
	if playery == MAXY-1 && playerx == 33 {
		return 1 /* not in entrance */
	}
	item[playerx][playery] = theitem
	iarg[playerx][playery] = ivenarg[k]
	srcount = 0
	lprcat("\n  You drop:")
	show3(k) /* show what item you dropped */
	know[playerx][playery] = false
	iven[k] = 0
	if c[WIELD] == k {
		c[WIELD] = -1
	}
	if c[WEAR] == k {
		c[WEAR] = -1
	}
	if c[SHIELD] == k {
		c[SHIELD] = -1
	}
	adjustcvalues(theitem, ivenarg[k])
	dropflag = 1 /* say dropped an item so wont ask to pick it
	 * up right away */
	return 0
}

/*
	function to enchant armor player is currently wearing
*/
func enchantarmor() {
	if c[WEAR] < 0 {
		if c[SHIELD] < 0 {
			cursors()
			beep()
			lprcat("\nYou feel a sense of loss")
			return
		} else {
			tmp := iven[c[SHIELD]]
			if tmp != OSCROLL {
				if tmp != OPOTION {
					ivenarg[c[SHIELD]]++
					bottomline()
				}
			}
		}
	}
	tmp := iven[c[WEAR]]
	if tmp != OSCROLL {
		if tmp != OPOTION {
			ivenarg[c[WEAR]]++
			bottomline()
		}
	}
}

/*
	function to enchant a weapon presently being wielded
*/
func enchweapon() {
	if c[WIELD] < 0 {
		cursors()
		beep()
		lprcat("\nYou feel a sense of loss")
		return
	}
	tmp := iven[c[WIELD]]
	if tmp != OSCROLL {
		if tmp != OPOTION {
			ivenarg[c[WIELD]]++
			if tmp == OCLEVERRING {
				c[INTELLIGENCE]++
			} else if tmp == OSTRRING {
				c[STREXTRA]++
			} else if tmp == ODEXRING {
				c[DEXTERITY]++
			}
			bottomline()
		}
	}
}

/*
	routine to tell if player can carry one more thing
	returns true if pockets are full, else false
*/
func pocketfull() bool {
	limit := 15 + (c[LEVEL] >> 1)
	if limit > 26 {
		limit = 26
	}
	for i := 0; i < limit; i++ {
		if iven[i] == 0 {
			return false
		}
	}
	return true
}

/*
	function to return true if a monster is next to the player else returns false
*/
func nearbymonst() bool {
	for x := max(playerx-1, 0); x < min(playerx+2, MAXX-1); x++ {
		for y := max(playery-1, 0); y < min(playery+2, MAXY-1); y++ {
			if mitem[x][y] != 0 {
				return true /* if monster nearby */
			}
		}
	}
	return false
}

/*
	function to steal an item from the players pockets
	returns 1 if steals something else returns 0
*/
func stealsomething() int {
	j := 100
	for {
		i := rund(26)
		if iven[i] != 0 {
			if c[WEAR] != i {
				if c[WIELD] != i {
					if c[SHIELD] != i {
						srcount = 0
						show3(i)
						adjustcvalues(iven[i], ivenarg[i])
						iven[i] = 0
						return 1
					}
				}
			}
		}
		j--
		if j <= 0 {
			return 0
		}
	}
	panic("unreachable")
}

/*
	function to return 1 is player carrys nothing else return 0
*/
func emptyhanded() int {
	for i := 0; i < 26; i++ {
		if iven[i] != 0 {
			if i != c[WIELD] {
				if i != c[WEAR] {
					if i != c[SHIELD] {
						return 0
					}
				}
			}
		}
	}
	return 1
}

/*
	function to create a gem on a square near the player
*/
func creategem() {
	var i, j int
	switch rnd(4) {
	case 1:
		i = ODIAMOND
		j = 50
	case 2:
		i = ORUBY
		j = 40
	case 3:
		i = OEMERALD
		j = 30
	default:
		i = OSAPPHIRE
		j = 20
	}
	createitem(i, rnd(j)+j/10)
}

/*
	function to change character levels as needed when dropping an object
	that affects these characteristics
*/
func adjustcvalues(theitem, arg int) {
	flag := false
	switch theitem {
	case ODEXRING:
		c[DEXTERITY] -= arg + 1
		flag = true
	case OSTRRING:
		c[STREXTRA] -= arg + 1
		flag = true
	case OCLEVERRING:
		c[INTELLIGENCE] -= arg + 1
		flag = true
	case OHAMMER:
		c[DEXTERITY] -= 10
		c[STREXTRA] -= 10
		c[INTELLIGENCE] += 10
		flag = true
	case OSWORDofSLASHING:
		c[DEXTERITY] -= 5
		flag = true
	case OORBOFDRAGON:
		c[SLAYING]--
		return
	case OSPIRITSCARAB:
		c[NEGATESPIRIT]--
		return
	case OCUBEofUNDEAD:
		c[CUBEofUNDEAD]--
		return
	case ONOTHEFT:
		c[NOTHEFT]--
		return
	case OLANCE:
		c[LANCEDEATH] = 0
		return
	case OPOTION, OSCROLL:
		return

	default:
		flag = true
	}
	if flag {
		bottomline()
	}
}

/*
	function to ask user for a password (no echo)
	returns true if entered correctly, false if not
*/
func getpassword() bool {
	scbr() /* system("stty -echo cbreak"); */
	lprcat("\nEnter Password: ")
	lflush()
	i := len(password)
	var gpwbuf string
	for j := 0; j < i; j++ {
		gpwbuf += string(ttgetch())
	}
	sncbr() /* system("stty echo -cbreak"); */
	if gpwbuf != password {
		lprcat("\nSorry\n")
		lflush()
		return false
	}
	return true
}

/*
	subroutine to get a yes or no response from the user
	returns y or n
*/
func getyn() int {
	i := 0
	for i != 'y' && i != 'n' && i != '\033' {
		i = ttgetch()
	}
	return i
}

/*
	function to calculate the pack weight of the player
	returns the number of pounds the player is carrying
*/
func packweight() int {
	k := c[GOLD] / 1000
	j := 25
	for iven[j] == 0 && j > 0 {
		j--
	}
	for i := 0; i <= j; i++ {
		switch iven[i] {
		case 0:
		case OSSPLATE, OPLATEARMOR:
			k += 40
		case OPLATE:
			k += 35
		case OHAMMER:
			k += 30
		case OSPLINT:
			k += 26
		case OSWORDofSLASHING, OCHAIN, OBATTLEAXE, O2SWORD:
			k += 23
		case OLONGSWORD, OSWORD, ORING, OFLAIL:
			k += 20
		case OLANCE, OSTUDLEATHER:
			k += 15
		case OLEATHER, OSPEAR:
			k += 8
		case OORBOFDRAGON, OBELT:
			k += 4
		case OSHIELD:
			k += 7
		case OCHEST:
			k += 30 + ivenarg[i]
		default:
			k++
		}
	}
	return k
}

/* macros to generate random numbers   1<=rnd(N)<=N   0<=rund(N)<=N-1 */
func rnd(x int) int {
	randx = randx*1103515245 + 12345
	return int(((randx >> 7) % uint32(x)) + 1)
}

func rund(x int) int {
	randx = randx*1103515245 + 12345
	return int((randx >> 7) % uint32(x))
}
