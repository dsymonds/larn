package main

import (
	"log"
)

/*
	makeplayer()

	subroutine to create the player and the players attributes
	this is called at the beginning of a game and at no other time
*/
func makeplayer() {
	scbr()
	clear()
	c[HPMAX], c[HP] = 10, 10      /* start player off with 15 hit points	 */
	c[LEVEL] = 1                  /* player starts at level one		 */
	c[SPELLMAX], c[SPELLS] = 1, 1 /* total # spells starts off as 3 */
	c[REGENCOUNTER] = 16
	c[ECOUNTER] = 96 /* start regeneration correctly */
	c[SHIELD], c[WEAR], c[WIELD] = -1, -1, -1
	for i := 0; i < 26; i++ {
		iven[i] = 0
	}
	spelknow[0], spelknow[1] = true, true /* he knows protection, magic missile */
	if c[HARDGAME] <= 0 {
		iven[0] = OLEATHER
		iven[1] = ODAGGER
		ivenarg[1], ivenarg[0], c[WEAR] = 0, 0, 0
		c[WIELD] = 1
	}
	playerx = rnd(MAXX - 2)
	playery = rnd(MAXY - 2)
	oldx = 0
	oldy = 25
	gltime = 0 /* time clock starts at zero	 */
	cbak[SPELLS] = -50
	for i := 0; i < 6; i++ {
		c[i] = 12 /* make the attributes, ie str, int, etc. */
	}
	recalc()
}

/*
	newcavelevel(level)
	int level;

	function to enter a new level.  This routine must be called anytime the
	player changes levels.  If that level is unknown it will be created.
	A new set of monsters will be created for a new level, and existing
	levels will get a few more monsters.
	Note that it is here we remove genocided monsters from the present level.
*/
func newcavelevel(x int) {
	if beenhere[level] {
		savelevel() /* put the level back into storage	 */
	}
	level = x /* get the new level and put in working storage */
	if beenhere[x] {
		getlevel()
		sethp(0)
		checkgen()
		return
	}

	/* fill in new level */
	for i := 0; i < MAXY; i++ {
		for j := 0; j < MAXX; j++ {
			know[j][i], mitem[j][i] = false, 0
		}
	}
	makemaze(x)
	makeobject(x)
	beenhere[x] = true
	sethp(1)
	checkgen() /* wipe out any genocided monsters */

	if (WIZID && wizard) || x == 0 {
		for j := 0; j < MAXY; j++ {
			for i := 0; i < MAXX; i++ {
				know[i][j] = true
			}
		}
	}
}

/*
	makemaze(level)
	int level;

	subroutine to make the caverns for a given level.  only walls are made.
*/
var mx, mxl, mxh, my, myl, myh int // static

func makemaze(k int) {
	if k > 1 && (rnd(17) <= 4 || k == MAXLEVEL-1 || k == MAXLEVEL+MAXVLEVEL-1) {
		if cannedlevel(k) {
			return /* read maze from data file */
		}
	}
	var tmp int
	if k == 0 {
		tmp = 0
	} else {
		tmp = OWALL
	}
	for i := 0; i < MAXY; i++ {
		for j := 0; j < MAXX; j++ {
			item[j][i] = tmp
		}
	}
	if k == 0 {
		return
	}
	eat(1, 1)
	if k == 1 {
		item[33][MAXY-1] = 0 /* exit from dungeon */
	}

	/* now for open spaces -- not on level 10	 */
	if k != MAXLEVEL-1 {
		tmp2 := rnd(3) + 3
		for tmp = 0; tmp < tmp2; tmp++ {
			my = rnd(11) + 2
			myl = my - rnd(2)
			myh = my + rnd(2)
			var z int
			if k < MAXLEVEL {
				mx = rnd(44) + 5
				mxl = mx - rnd(4)
				mxh = mx + rnd(12) + 3
				z = 0
			} else {
				mx = rnd(60) + 3
				mxl = mx - rnd(2)
				mxh = mx + rnd(2)
				z = makemonst(k)
			}
			for i := mxl; i < mxh; i++ {
				for j := myl; j < myh; j++ {
					item[i][j] = 0
					mitem[i][j] = z
					if mitem[i][j] != 0 {
						hitp[i][j] = monster[z].hitpoints
					}
				}
			}
		}
	}
	if k != MAXLEVEL-1 {
		my = rnd(MAXY - 2)
		for i := 1; i < MAXX-1; i++ {
			item[i][my] = 0
		}
	}
	if k > 1 {
		treasureroom(k)
	}
}

/*
	function to eat away a filled in maze
*/
func eat(xx, yy int) {
	dir := rnd(4)
	try := 2
	for try != 0 {
		switch dir {
		case 1:
			if xx <= 2 {
				break /* west	 */
			}
			if item[xx-1][yy] != OWALL || item[xx-2][yy] != OWALL {
				break
			}
			item[xx-1][yy], item[xx-2][yy] = 0, 0
			eat(xx-2, yy)

		case 2:
			if xx >= MAXX-3 {
				break /* east	 */
			}
			if item[xx+1][yy] != OWALL || item[xx+2][yy] != OWALL {
				break
			}
			item[xx+1][yy], item[xx+2][yy] = 0, 0
			eat(xx+2, yy)

		case 3:
			if yy <= 2 {
				break /* south	 */
			}
			if item[xx][yy-1] != OWALL || item[xx][yy-2] != OWALL {
				break
			}
			item[xx][yy-1], item[xx][yy-2] = 0, 0
			eat(xx, yy-2)

		case 4:
			if yy >= MAXY-3 {
				break /* north	 */
			}
			if item[xx][yy+1] != OWALL || item[xx][yy+2] != OWALL {
				break
			}
			item[xx][yy+1], item[xx][yy+2] = 0, 0
			eat(xx, yy+2)
		}
		dir++
		if dir > 4 {
			dir = 1
			try--
		}
	}
}

/*
 *	function to read in a maze from a data file
 *
 *	Format of maze data file:  1st character = # of mazes in file (ascii digit)
 *				For each maze: 18 lines (1st 17 used) 67 characters per line
 *
 *	Special characters in maze data file:
 *
 *		#	wall			D	door			.	random monster
 *		~	eye of larn		!	cure dianthroritis
 *		-	random object
 */
func cannedlevel(k int) bool {
	if !lopen(larnlevels) {
		log.Print("Can't open the maze data file")
		died(-282)
		return false
	}
	i := lgetc()
	if i <= '0' {
		died(-282)
		return false
	}
	for i = 18 * rund(i-'0'); i > 0; i-- {
		lgetl() /* advance to desired maze */
	}
	for i := 0; i < MAXY; i++ {
		row := lgetl()
		for j := 0; j < MAXX; j++ {
			it, mit, arg, marg := 0, 0, 0, 0
			ch := row[0]
			row = row[1:]
			switch ch {
			case '#':
				it = OWALL
			case 'D':
				it = OCLOSEDDOOR
				arg = rnd(30)
			case '~':
				if k != MAXLEVEL-1 {
					break
				}
				it = OLARNEYE
				mit = rund(8) + DEMONLORD
				marg = monster[mit].hitpoints
			case '!':
				if k != MAXLEVEL+MAXVLEVEL-1 {
					break
				}
				it = OPOTION
				arg = 21
				mit = DEMONLORD + 7
				marg = monster[mit].hitpoints
			case '.':
				if k < MAXLEVEL {
					break
				}
				mit = makemonst(k + 1)
				marg = monster[mit].hitpoints
			case '-':
				it = newobject(k+1, &arg)
			}
			item[j][i] = it
			iarg[j][i] = arg
			mitem[j][i] = mit
			hitp[j][i] = marg

			know[j][i] = false
			if WIZID && wizard {
				know[j][i] = true
			}
		}
	}
	lrclose()
	return true
}

/*
	function to make a treasure room on a level
	level 10's treasure room has the eye in it and demon lords
	level V3 has potion of cure dianthroritis and demon prince
*/
func treasureroom(lv int) {
	for tx := 1 + rnd(10); tx < MAXX-10; tx += 10 {
		if lv == MAXLEVEL-1 || lv == MAXLEVEL+MAXVLEVEL-1 || rnd(13) == 2 {
			xsize := rnd(6) + 3
			ysize := rnd(3) + 3
			ty := rnd(MAXY-9) + 1 /* upper left corner of room */
			if lv == MAXLEVEL-1 || lv == MAXLEVEL+MAXVLEVEL-1 {
				tx = tx + rnd(MAXX-24)
				troom(lv, xsize, ysize, tx, ty, rnd(3)+6)
			} else {
				troom(lv, xsize, ysize, tx, ty, rnd(9))
			}
		}
	}
}

/*
 *	subroutine to create a treasure room of any size at a given location
 *	room is filled with objects and monsters
 *	the coordinate given is that of the upper left corner of the room
 */
func troom(lv, xsize, ysize, tx, ty, glyph int) {
	for j := ty - 1; j <= ty+ysize; j++ {
		for i := tx - 1; i <= tx+xsize; i++ { /* clear out space for room */
			item[i][j] = 0
		}
	}
	for j := ty; j < ty+ysize; j++ {
		for i := tx; i < tx+xsize; i++ { /* now put in the walls */
			item[i][j] = OWALL
			mitem[i][j] = 0
		}
	}
	for j := ty + 1; j < ty+ysize-1; j++ {
		for i := tx + 1; i < tx+xsize-1; i++ { /* now clear out interior */
			item[i][j] = 0
		}
	}

	switch rnd(2) { /* locate the door on the treasure room */
	case 1:
		i := tx + rund(xsize)
		j := ty + (ysize-1)*rund(2)
		item[i][j] = OCLOSEDDOOR
		iarg[i][j] = glyph /* on horizontal walls */
	case 2:
		i := tx + (xsize-1)*rund(2)
		j := ty + rund(ysize)
		item[i][j] = OCLOSEDDOOR
		iarg[i][j] = glyph /* on vertical walls */
	}

	tp1 := playerx
	tp2 := playery
	playery = ty + (ysize >> 1)
	if c[HARDGAME] < 2 {
		for playerx = tx + 1; playerx <= tx+xsize-2; playerx += 2 {
			for i, j := 0, rnd(6); i <= j; i++ {
				something(lv + 2)
				createmonster(makemonst(lv + 1))
			}
		}
	} else {
		for playerx = tx + 1; playerx <= tx+xsize-2; playerx += 2 {
			for i, j := 0, rnd(4); i <= j; i++ {
				something(lv + 2)
				createmonster(makemonst(lv + 3))
			}
		}
	}

	playerx = tp1
	playery = tp2
}

/*
	***********
	MAKE_OBJECT
	***********
	subroutine to create the objects in the maze for the given level
*/
func makeobject(j int) {
	if j == 0 {
		fillroom(OENTRANCE, 0)  /* entrance to dungeon			 */
		fillroom(ODNDSTORE, 0)  /* the DND STORE				 */
		fillroom(OSCHOOL, 0)    /* college of Larn				 */
		fillroom(OBANK, 0)      /* 1st national bank of larn 	 */
		fillroom(OVOLDOWN, 0)   /* volcano shaft to temple 	 */
		fillroom(OHOME, 0)      /* the players home & family 	 */
		fillroom(OTRADEPOST, 0) /* the trading post			 */
		fillroom(OLRS, 0)       /* the larn revenue service 	 */
		return
	}
	if j == MAXLEVEL {
		fillroom(OVOLUP, 0) /* volcano shaft up from the temple */
	}

	/* make the fixed objects in the maze STAIRS	 */
	if j > 0 && j != MAXLEVEL-1 && j != MAXLEVEL+MAXVLEVEL-1 {
		fillroom(OSTAIRSDOWN, 0)
	}
	if j > 1 && j != MAXLEVEL {
		fillroom(OSTAIRSUP, 0)
	}

	/* make the random objects in the maze		 */

	fillmroom(rund(3), OBOOK, j)
	fillmroom(rund(3), OALTAR, 0)
	fillmroom(rund(3), OSTATUE, 0)
	fillmroom(rund(3), OPIT, 0)
	fillmroom(rund(3), OFOUNTAIN, 0)
	fillmroom(rnd(3)-2, OIVTELETRAP, 0)
	fillmroom(rund(2), OTHRONE, 0)
	fillmroom(rund(2), OMIRROR, 0)
	fillmroom(rund(2), OTRAPARROWIV, 0)
	fillmroom(rnd(3)-2, OIVDARTRAP, 0)
	fillmroom(rund(3), OCOOKIE, 0)
	if j == 1 {
		fillmroom(1, OCHEST, j)
	} else {
		fillmroom(rund(2), OCHEST, j)
	}
	if j != MAXLEVEL-1 && j != MAXLEVEL+MAXVLEVEL-1 {
		fillmroom(rund(2), OIVTRAPDOOR, 0)
	}
	if j <= 10 {
		fillmroom((rund(2)), ODIAMOND, rnd(10*j+1)+10)
		fillmroom(rund(2), ORUBY, rnd(6*j+1)+6)
		fillmroom(rund(2), OEMERALD, rnd(4*j+1)+4)
		fillmroom(rund(2), OSAPPHIRE, rnd(3*j+1)+2)
	}
	for i := 0; i < rnd(4)+3; i++ {
		fillroom(OPOTION, newpotion()) /* make a POTION	 */
	}
	for i := 0; i < rnd(5)+3; i++ {
		fillroom(OSCROLL, newscroll()) /* make a SCROLL	 */
	}
	for i := 0; i < rnd(12)+11; i++ {
		fillroom(OGOLDPILE, 12*rnd(j+1)+(j<<3)+10) /* make GOLD	 */
	}
	if j == 5 {
		fillroom(OBANK2, 0) /* branch office of the bank */
	}
	froom(2, ORING, 0)            /* a ring mail 			 */
	froom(1, OSTUDLEATHER, 0)     /* a studded leather	 */
	froom(3, OSPLINT, 0)          /* a splint mail		 */
	froom(5, OSHIELD, rund(3))    /* a shield				 */
	froom(2, OBATTLEAXE, rund(3)) /* a battle axe			 */
	froom(5, OLONGSWORD, rund(3)) /* a long sword			 */
	froom(5, OFLAIL, rund(3))     /* a flail				 */
	froom(4, OREGENRING, rund(3)) /* ring of regeneration */
	froom(1, OPROTRING, rund(3))  /* ring of protection	 */
	froom(2, OSTRRING, 4)         /* ring of strength + 4 */
	froom(7, OSPEAR, rnd(5))      /* a spear				 */
	froom(3, OORBOFDRAGON, 0)     /* orb of dragon slaying */
	froom(4, OSPIRITSCARAB, 0)    /* scarab of negate spirit */
	froom(4, OCUBEofUNDEAD, 0)    /* cube of undead control	 */
	froom(2, ORINGOFEXTRA, 0)     /* ring of extra regen		 */
	froom(3, ONOTHEFT, 0)         /* device of antitheft 		 */
	froom(2, OSWORDofSLASHING, 0) /* sword of slashing */
	if c[BESSMANN] == 0 {
		froom(4, OHAMMER, 0) /* Bessman's flailing hammer */
		c[BESSMANN] = 1
	}
	if c[HARDGAME] < 3 || rnd(4) == 3 {
		if j > 3 {
			froom(3, OSWORD, 3)       /* sunsword + 3  		 */
			froom(5, O2SWORD, rnd(4)) /* a two handed sword	 */
			froom(3, OBELT, 4)        /* belt of striking		 */
			froom(3, OENERGYRING, 3)  /* energy ring			 */
			froom(4, OPLATE, 5)       /* platemail + 5 		 */
		}
	}
}

/*
	subroutine to fill in a number of objects of the same kind
*/

func fillmroom(n, what, arg int) {
	for i := 0; i < n; i++ {
		fillroom(what, arg)
	}
}

func froom(n, theitem, arg int) {
	if rnd(151) < n {
		fillroom(theitem, arg)
	}
}

/*
	subroutine to put an object into an empty room
 *	uses a random walk
*/
func fillroom(what, arg int) {
	c[FILLROOM]++

	x := rnd(MAXX - 2)
	y := rnd(MAXY - 2)
	for item[x][y] != 0 {

		c[RANDOMWALK]++ /* count up these random walks */

		x += rnd(3) - 2
		y += rnd(3) - 2
		if x > MAXX-2 {
			x = 1
		}
		if x < 1 {
			x = MAXX - 2
		}
		if y > MAXY-2 {
			y = 1
		}
		if y < 1 {
			y = MAXY - 2
		}
	}
	item[x][y] = what
	iarg[x][y] = arg
}

/*
	subroutine to put monsters into an empty room without walls or other
	monsters
*/
func fillmonst(what int) int {
	for trys := 5; trys > 0; trys-- { /* max # of creation attempts */
		x := rnd(MAXX - 2)
		y := rnd(MAXY - 2)
		if item[x][y] == 0 && mitem[x][y] == 0 && playerx != x || playery != y {
			mitem[x][y] = what
			know[x][y] = false
			hitp[x][y] = monster[what].hitpoints
			return 0
		}
	}
	return -1 /* creation failure */
}

/*
	creates an entire set of monsters for a level
	must be done when entering a new level
	if sethp(1) then wipe out old monsters else leave them there
*/
func sethp(flg int) {
	if flg != 0 {
		for i := 0; i < MAXY; i++ {
			for j := 0; j < MAXX; j++ {
				stealth[j][i] = 0
			}
		}
	}
	if level == 0 {
		c[TELEFLAG] = 0
		return
	} /* if teleported and found level 1 then know level we are on */
	var j int
	if flg != 0 {
		j = rnd(12) + 2 + (level >> 1)
	} else {
		j = (level >> 1) + 1
	}
	for i := 0; i < j; i++ {
		fillmonst(makemonst(level))
	}
	positionplayer()
}

/*
 *	Function to destroy all genocided monsters on the present level
 */
func checkgen() {
	for y := 0; y < MAXY; y++ {
		for x := 0; x < MAXX; x++ {
			if monster[mitem[x][y]].genocided {
				mitem[x][y] = 0 /* no more monster */
			}
		}
	}
}
