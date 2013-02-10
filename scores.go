package main

import (
	"os"
)

/*
 * Functions in this file are:
 *
 * readboard() 	Function to read in the scoreboard into a static buffer
 * writeboard()	Function to write the scoreboard from readboard()'s buffer
 * makeboard() 	Function to create a new scoreboard (wipe out old one)
 * hashewon()	 Function to return 1 if player has won a game before, else 0
 * long paytaxes(x)	 Function to pay taxes if any are due winshou()
 * ubroutine to print out the winning scoreboard shou(x)
 * ubroutine to print out the non-winners scoreboard showscores()
 * unction to show the scoreboard on the terminal showallscores()
 * Function to show scores and the iven lists that go with them sortboard()
 * unction to sort the scoreboard newscore(score, whoo, whyded, winner)
 * Function to add entry to scoreboard new1sub(score,i,whoo,taxes)
 * Subroutine to put player into a new2sub(score,i,whoo,whyded)
 * Subroutine to put player into a died(x) 	Subroutine to record who
 * played larn, and what the score was diedsub(x) Subroutine to print out a
 * line showing player when he is killed diedlog() 	Subroutine to read a
 * log file and print it out in ascii format getplid(name)
 * on to get players id # from id file
 *
 */

/* This is the structure for the scoreboard 		 */
type scofmt struct {
	score   int32       /* the score of the player 							 */
	suid    int32       /* the user id number of the player 				 */
	what    int16       /* the number of the monster that killed player 	 */
	level   int16       /* the level player was on when he died 			 */
	hardlev int16       /* the level of difficulty player played at 		 */
	order   int16       /* the relative ordering place of this entry 		 */
	who     [40]int8    /* the name of the character 						 */
	sciv    [26][2]int8 /* this is the inventory list of the character 		 */
}

/* This is the structure for the winning scoreboard */
type wscofmt struct {
	score    int32    /* the score of the player 							 */
	timeused int32    /* the time used in mobuls to win the game 			 */
	taxes    int32    /* taxes he owes to LRS 							 */
	suid     int32    /* the user id number of the player 				 */
	hardlev  int16    /* the level of difficulty player played at 		 */
	order    int16    /* the relative ordering place of this entry 		 */
	who      [40]int8 /* the name of the character 						 */
}

/* 102 bytes struct for the log file 				 */
type log_fmt struct {
	score          int32 /* the players score 								 */
	diedtime       int32 /* time when game was over 							 */
	cavelev        int16 /* level in caves 									 */
	diff           int16 /* difficulty player played at 						 */
	elapsedtime    int32 /* real time of game in seconds 					 */
	bytout         int32 /* bytes input and output 							 */
	bytin          int32
	moves          int32    /* number of moves made by player 					 */
	ac             int16    /* armor class of player 							 */
	hp, hpmax      int16    /* players hitpoints 								 */
	cputime        int16    /* CPU time needed in seconds 						 */
	killed, spused int16    /* monsters killed and spells cast 					 */
	usage          int16    /* usage of the CPU in % 							 */
	lev            int16    /* player level 									 */
	who            [12]int8 /* player name 										 */
	what           [46]int8 /* what happened to player 							 */
}

var sco [SCORESIZE]scofmt   /* the structure for the scoreboard  */
var winr [SCORESIZE]wscofmt /* struct for the winning scoreboard */
var logg log_fmt            /* structure for the log file 		 */
var whydead = [...]string{
	"quit", "suspended", "self - annihilated", "shot by an arrow",
	"hit by a dart", "fell into a pit", "fell into a bottomless pit",
	"a winner", "trapped in solid rock", "killed by a missing save file",
	"killed by an old save file", "caught by the greedy cheater checker trap",
	"killed by a protected save file", "killed his family and committed suicide",
	"erased by a wayward finger", "fell through a bottomless trap door",
	"fell through a trap door", "drank some poisonous water",
	"fried by an electric shock", "slipped on a volcano shaft",
	"killed by a stupid act of frustration", "attacked by a revolting demon",
	"hit by his own magic", "demolished by an unseen attacker",
	"fell into the dreadful sleep", "killed by an exploding chest",
	/* 26 */ "killed by a missing maze data file", "annihilated in a sphere",
	"died a post mortem death", "wasted by a malloc() failure",
}

/*
 * readboard() 	Function to read in the scoreboard into a static buffer
 *
 * returns -1 if unable to read in the scoreboard, returns 0 if all is OK
 */
func readboard() int {
	if gid != egid {
		setegid(egid)
	}
	i := lopen(scorefile)
	if gid != egid {
		setegid(gid)
	}
	if i < 0 {
		lprcat("Can't read scoreboard\n")
		lflush()
		return -1
	}
	// TODO
	return -1
	/*
		lrfill((char *) sco, sizeof(sco))
		lrfill((char *) winr, sizeof(winr))
		lrclose()
		lcreat((char *) 0)
		return (0)
	*/
}

/*
 * writeboard()	Function to write the scoreboard from readboard()'s buffer
 *
 * returns -1 if unable to write the scoreboard, returns 0 if all is OK
 */
func writeboard() int {
	set_score_output()
	if gid != egid {
		setegid(egid)
	}
	i := lcreat(scorefile)
	if gid != egid {
		setegid(gid)
	}
	if i < 0 {
		lprcat("Can't write scoreboard\n")
		lflush()
		return (-1)
	}
	// TODO
	return -1
	/*
		lwrite((char *) sco, sizeof(sco))
		lwrite((char *) winr, sizeof(winr))
		lwclose()
		lcreat((char *) 0)
		return (0)
	*/
}

/*
 * makeboard() 		Function to create a new scoreboard (wipe out old one)
 *
 * returns -1 if unable to write the scoreboard, returns 0 if all is OK
 */
func makeboard() int {
	set_score_output()
	for i := 0; i < SCORESIZE; i++ {
		winr[i].taxes, winr[i].score, sco[i].score = 0, 0, 0
		winr[i].order, sco[i].order = i, i
	}
	if writeboard() {
		return -1
	}
	if gid != egid {
		setegid(egid)
	}
	chmod(scorefile, 0660)
	if gid != egid {
		setegid(gid)
	}
	return 0
}

/*
 * hashewon()	 Function to return 1 if player has won a game before, else 0
 *
 * This function also sets c[HARDGAME] to appropriate value -- 0 if not a
 * winner, otherwise the next level of difficulty listed in the winners
 * scoreboard.  This function also sets outstanding_taxes to the value in
 * the winners scoreboard.
 */
func hashewon() int {
	c[HARDGAME] = 0
	if readboard() < 0 {
		return 0 /* can't find scoreboard */
	}
	for i = 0; i < SCORESIZE; i++ { /* search through winners scoreboard */
		if winr[i].suid == userid {
			if winr[i].score > 0 {
				c[HARDGAME] = winr[i].hardlev + 1
				outstanding_taxes = winr[i].taxes
				return 1
			}
		}
	}
	return 0
}

/*
 * long paytaxes(x)		 Function to pay taxes if any are due
 *
 * Enter with the amount (in gp) to pay on the taxes.
 * Returns amount actually paid.
 */
func paytaxes(int32 x) int32 {
	if x < 0 {
		return 0
	}
	if readboard() < 0 {
		return 0
	}
	for i := 0; i < SCORESIZE; i++ {
		if winr[i].suid == userid { /* look for players winning entry */
			if winr[i].score > 0 { /* search for a winning entry for the player */
				amt := winr[i].taxes
				if x < amt {
					amt = x /* don't overpay taxes (Ughhhhh) */
				}
				winr[i].taxes -= amt
				outstanding_taxes -= amt
				set_score_output()
				if writeboard() < 0 {
					return 0
				}
				return amt
			}
		}
	}
	return 0 /* couldn't find user on winning scoreboard */
}

/*
 * winshou()		Subroutine to print out the winning scoreboard
 *
 * Returns the number of players on scoreboard that were shown
 */
func winshou() int {
	j, count := 0, 0
	for i := 0; i < SCORESIZE; i++ { /* is there anyone on * the scoreboard? */
		if winr[i].score != 0 {
			j++
			break
		}
	}
	if j != 0 {
		lprcat("\n  Score    Difficulty   Time Needed   Larn Winners List\n")

		for i := 0; i < SCORESIZE; i++ { /* this loop is needed to print out the */
			for j = 0; j < SCORESIZE; j++ { /* winners in order */
				p := &winr[j] /* pointer to the scoreboard entry */
				if p.order == i {
					if p.score {
						count++
						lprintf("%10d     %2d      %5d Mobuls   %s \n",
							p.score, p.hardlev, p.timeused, p.who)
					}
					break
				}
			}
		}
	}
	return count /* return number of people on scoreboard */
}

/*
 * shou(x)			Subroutine to print out the non-winners scoreboard
 * 	int x;
 *
 * Enter with 0 to list the scores, enter with 1 to list inventories too
 * Returns the number of players on scoreboard that were shown
 */
func shou(int x) int {
	j, count := 0, 0
	for i := 0; i < SCORESIZE; i++ { /* is the scoreboard empty? */
		if sco[i].score != 0 {
			j++
			break
		}
	}
	if j != 0 {
		lprcat("\n   Score   Difficulty   Larn Visitor Log\n")
		for i := 0; i < SCORESIZE; i++ { /* be sure to print them out in order */
			for j = 0; j < SCORESIZE; j++ {
				if sco[j].order == i {
					if sco[j].score {
						count++
						lprintf("%10d     %2d       %s ",
							sco[j].score, sco[j].hardlev, sco[j].who)
						if sco[j].what < 256 {
							lprintf("killed by a %s", monster[sco[j].what].name)
						} else {
							lprintf("%s", whydead[sco[j].what-256])
						}
						if x != 263 {
							lprintf(" on %s", levelname[sco[j].level])
						}
						if x {
							for n := 0; n < 26; n++ {
								iven[n] = sco[j].sciv[n][0]
								ivenarg[n] = sco[j].sciv[n][1]
							}
							for k := 1; k < 99; k++ {
								for n = 0; n < 26; n++ {
									if k == iven[n] {
										srcount = 0
										show3(n)
									}
								}
							}
							lprcat("\n\n")
						} else {
							lprc('\n')
						}
					}
					j = SCORESIZE
				}
			}
		}
	}
	return count /* return the number of players just shown */
}

/*
 * showscores()		Function to show the scoreboard on the terminal
 *
 * Returns nothing of value
 */
const esb = "The scoreboard is empty.\n"

func showscores() {
	lflush()
	lcreat("")
	if readboard() < 0 {
		return
	}
	i := winshou()
	j := shou(0)
	if i+j == 0 {
		lprcat(esb)
	} else {
		lprc('\n')
	}
	lflush()
}

/*
 * showallscores()	Function to show scores and the iven lists that go with them
 *
 * Returns nothing of value
 */
func showallscores() {
	lflush()
	lcreat("")
	if readboard() < 0 {
		return
	}
	c[WEAR], c[WIELD], c[SHIELD] = -1, -1, -1 /* not wielding or wearing anything */
	for i := 0; i < MAXPOTION; i++ {
		potionname[i] = potionhide[i]
	}
	for i := 0; i < MAXSCROLL; i++ {
		scrollname[i] = scrollhide[i]
	}
	i = winshou()
	j = shou(1)
	if i+j == 0 {
		lprcat(esb)
	} else {
		lprc('\n')
	}
	lflush()
}

/*
 * sortboard()		Function to sort the scoreboard
 *
 * Returns 0 if no sorting done, else returns 1
 */
func sortboard() int {
	for i = 0; i < SCORESIZE; i++ {
		sco[i].order, winr[i].order = -1, -1
	}
	j, pos = 0, 0
	for pos < SCORESIZE {
		jdat := 0
		for i := 0; i < SCORESIZE; i++ {
			if sco[i].order < 0 && sco[i].score >= jdat {
				j = i
				jdat = sco[i].score
			}
		}
		sco[j].order = pos
		pos++
	}
	pos = 0
	for pos < SCORESIZE {
		jdat := 0
		for i := 0; i < SCORESIZE; i++ {
			if winr[i].order < 0 && winr[i].score >= jdat {
				j = i
				jdat = winr[i].score
			}
		}
		winr[j].order = pos
		pos++
	}
	return 1
}

/*
 * newscore(score, whoo, whyded, winner) 	Function to add entry to scoreboard
 * 	int score, winner, whyded;
 * 	char *whoo;
 *
 * Enter with the total score in gp in score,  players name in whoo,
 * 	died() reason # in whyded, and TRUE/FALSE in winner if a winner
 * ex.		newscore(1000, "player 1", 32, 0);
 */
func newscore(score int32, whoo string, whyded, winner int) {
	if readboard() < 0 {
		return /* do the scoreboard	 */
	}
	/* if a winner then delete all non-winning scores */
	if cheat {
		winner = 0 /* if he cheated, don't let him win */
	}
	if winner {
		for i := 0; i < SCORESIZE; i++ {
			if sco[i].suid == userid {
				sco[i].score = 0
			}
		}
		taxes := score * TAXRATE
		score += 100000 * c[HARDGAME] /* bonus for winning */
		/*
		 * if he has a slot on the winning scoreboard update it if
		 * greater score
		 */
		for i := 0; i < SCORESIZE; i++ {
			if winr[i].suid == userid {
				new1sub(score, i, whoo, taxes)
				return
			}
		}
		/*
		 * he had no entry. look for last entry and see if he has a
		 * greater score
		 */
		for i := 0; i < SCORESIZE; i++ {
			if winr[i].order == SCORESIZE-1 {
				new1sub(score, i, whoo, taxes)
				return
			}
		}
	} else if !cheat { /* for not winning scoreboard */
		/*
		 * if he has a slot on the scoreboard update it if greater
		 * score
		 */
		for i := 0; i < SCORESIZE; i++ {
			if sco[i].suid == userid {
				new2sub(score, i, whoo, whyded)
				return
			}
		}
		/*
		 * he had no entry. look for last entry and see if he has a
		 * greater score
		 */
		for i := 0; i < SCORESIZE; i++ {
			if sco[i].order == SCORESIZE-1 {
				new2sub(score, i, whoo, whyded)
				return
			}
		}
	}
}

/*
 * new1sub(score,i,whoo,taxes) 	  Subroutine to put player into a
 * 	int score,i,whyded,taxes;		  winning scoreboard entry if his score
 * 	char *whoo; 					  is high enough
 *
 * Enter with the total score in gp in score,  players name in whoo,
 * 	died() reason # in whyded, and TRUE/FALSE in winner if a winner
 * 	slot in scoreboard in i, and the tax bill in taxes.
 * Returns nothing of value
 */
func new1sub(score int32, i int, whoo string, taxes int32) {
	p := &winr[i]
	p.taxes += taxes
	if score >= p.score || c[HARDGAME] > p.hardlev {
		p.who = whoo
		p.score = score
		p.hardlev = c[HARDGAME]
		p.suid = userid
		p.timeused = gltime / 100
	}
}

/*
 * new2sub(score,i,whoo,whyded)	 	  Subroutine to put player into a
 * 	int score,i,whyded,taxes;		  non-winning scoreboard entry if his
 * 	char *whoo; 					  score is high enough
 *
 * Enter with the total score in gp in score,  players name in whoo,
 * 	died() reason # in whyded, and slot in scoreboard in i.
 * Returns nothing of value
 */
func new2sub(score int32, i int, whoo string, whyded int) {
	p := &sco[i]
	if score >= p.score || c[HARDGAME] > p.hardlev {
		p.who = whoo
		p.score = score
		p.what = whyded
		p.hardlev = c[HARDGAME]
		p.suid = userid
		p.level = level
		for j := 0; j < 26; j++ {
			p.sciv[j][0] = iven[j]
			p.sciv[j][1] = ivenarg[j]
		}
	}
}

/*
 * died(x) 	Subroutine to record who played larn, and what the score was
 * 	int x;
 *
 * if x < 0 then don't show scores
 * died() never returns! (unless c[LIFEPROT] and a reincarnatable death!)
 *
 * 	< 256	killed by the monster number
 * 	256		quit
 * 	257		suspended
 * 	258		self - annihilated
 * 	259		shot by an arrow
 * 	260		hit by a dart
 * 	261		fell into a pit
 * 	262		fell into a bottomless pit
 * 	263		a winner
 * 	264		trapped in solid rock
 * 	265		killed by a missing save file
 * 	266		killed by an old save file
 * 	267		caught by the greedy cheater checker trap
 * 	268		killed by a protected save file
 * 	269		killed his family and killed himself
 * 	270		erased by a wayward finger
 * 	271		fell through a bottomless trap door
 * 	272		fell through a trap door
 * 	273		drank some poisonous water
 * 	274		fried by an electric shock
 * 	275		slipped on a volcano shaft
 * 	276		killed by a stupid act of frustration
 * 	277		attacked by a revolting demon
 * 	278		hit by his own magic
 * 	279		demolished by an unseen attacker
 * 	280		fell into the dreadful sleep
 * 	281		killed by an exploding chest
 * 	282		killed by a missing maze data file
 * 	283		killed by a sphere of annihilation
 * 	284		died a post mortem death
 * 	285		malloc() failure
 * 	300		quick quit -- don't put on scoreboard
 */

var scorerror int

func died(x int) {
	if c[LIFEPROT] > 0 { /* if life protection */
		q := x
		if q < 0 {
			q = -q
		}
		switch q {
		case 256, 257, 262, 263, 265, 266, 267, 268, 269, 271, 282, 284, 285, 300:
			goto invalid /* can't be saved */
		}
		c[LIFEPROT]--
		c[HP] = 1
		c[CONSTITUTION]--
		cursors()
		lprcat("\nYou feel wiiieeeeerrrrrd all over! ")
		beep()
		lflush()
		sleep(4)
		return /* only case where died() returns */
	}
invalid:
	clearvt100()
	lflush()
	f := 0
	if ckpflag {
		unlink(ckpfile) /* remove checkpoint file if used */
	}
	if x < 0 {
		f++
		x = -x
	} /* if we are not to display the scores */
	if x == 300 || x == 257 {
		os.Exit(0) /* for quick exit or saved game */
	}
	var win int
	if x == 263 {
		win = 1
	} else {
		win = 0
	}
	c[GOLD] += c[BANKACCOUNT]
	c[BANKACCOUNT] = 0
	/* now enter the player at the end of the scoreboard */
	newscore(c[GOLD], logname, x, win)
	diedsub(x) /* print out the score line */
	lflush()

	set_score_output()
	if wizard == 0 && c[GOLD] > 0 { /* wizards can't score		 */
		if gid != egid {
			setegid(egid)
		}
		if lappend(logfile) < 0 { /* append to file */
			if lcreat(logfile) < 0 { /* and can't create new log file */
				lcreat("")
				lprcat("\nCan't open record file:  I can't post your score.\n")
				sncbr()
				resetscroll()
				lflush()
				exit(0)
			}
			if gid != egid {
				setegid(egid)
			}
			chmod(logfile, 0660)
			if gid != egid {
				setegid(gid)
			}
		}
		if gid != egid {
			setegid(gid)
		}
		logg.who = loginname
		logg.score = c[GOLD]
		logg.diff = c[HARDGAME]
		if x < 256 {
			ch := monster[x].name[0]
			var mod string
			if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
				mod = "an"
			} else {
				mod = "a"
			}
			snprintf(logg.what, sizeof(logg.what),
				"killed by %s %s", mod, monster[x].name)
		} else {
			snprintf(logg.what, sizeof(logg.what),
				"%s", whydead[x-256])
		}
		logg.cavelev = level
		time(&zzz) /* get CPU time -- write out score info */
		logg.diedtime = zzz

		// TODO
		// times(&cputime)/* get CPU time -- write out score info */
		/*
			logg.cputime = i = (cputime.tms_utime + cputime.tms_stime) / 60 + c[CPUTIME]
		*/
		logg.lev = c[LEVEL]
		logg.ac = c[AC]
		logg.hpmax = c[HPMAX]
		logg.hp = c[HP]
		logg.elapsedtime = (zzz - initialtime + 59) / 60
		logg.usage = (10000 * i) / (zzz - initialtime)
		logg.bytin = c[BYTESIN]
		logg.bytout = c[BYTESOUT]
		logg.moves = c[MOVESMADE]
		logg.spused = c[SPELLSCAST]
		logg.killed = c[MONSTKILLED]

		// TODO
		//lwrite((char *) &logg, sizeof(struct log_fmt))
		lwclose()

		/*
		 * now for the scoreboard maintenance -- not for a suspended
		 * game
		 */
		if x != 257 {
			if sortboard() {
				set_score_output()
				scorerror = writeboard()
			}
		}
	}
	if x == 256 || x == 257 || f != 0 {
		os.Exit(0)
	}
	if scorerror == 0 {
		showscores() /* if we updated the scoreboard */
	}
	if x == 263 {
		mailbill()
	}
	os.Exit(0)
}

/*
 * diedsub(x) Subroutine to print out the line showing the player when he is killed
 * 	int x;
 */
func diedsub(x int) {
	lprintf("Score: %d, Diff: %d,  %s ", c[GOLD], c[HARDGAME], logname)
	if x < 256 {
		ch := monster[x].name[0]
		var mod string
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			mod = "an"
		} else {
			mod = "a"
		}
		lprintf("killed by %s %s", mod, monster[x].name)
	} else {
		lprintf("%s", whydead[x-256])
	}
	if x != 263 {
		lprintf(" on %s\n", levelname[level])
	} else {
		lprc('\n')
	}
}

/*
 * diedlog() 	Subroutine to read a log file and print it out in ascii format
 */
func diedlog() {
	// TODO
	/*
		char  *p
		static char  q[] = "?"
		struct stat     stbuf
		time_t t

		lcreat((char *) 0)
		if (lopen(logfile) < 0) {
			lprintf("Can't locate log file <%s>\n", logfile)
			return
		}
		if (fstat(io_infd, &stbuf) < 0) {
			lprintf("Can't  stat log file <%s>\n", logfile)
			return
		}
		for n := stbuf.st_size / sizeof(struct log_fmt); n > 0; --n {
			lrfill((char *) &logg, sizeof(struct log_fmt))
			t = logg.diedtime
			if (p = ctime(&t)) == NULL {
				p = q
			} else {
				p[16] = '\n'
				p[17] = 0
			}
			lprintf("Score: %d, Diff: %d,  %s %s on %d at %s", (logg.score), (logg.diff), logg.who, logg.what, (logg.cavelev), p + 4)
	#ifdef EXTRA
			if logg.moves <= 0 {
				logg.moves = 1
			}
			lprintf("  Experience Level: %d,  AC: %d,  HP: %d/%d,  Elapsed Time: %d minutes\n", (logg.lev), (logg.ac), (logg.hp), (logg.hpmax), (logg.elapsedtime))
			lprintf("  CPU time used: %d seconds,  Machine usage: %d.%02d%%\n", (logg.cputime), (logg.usage / 100), (logg.usage % 100))
			lprintf("  BYTES in: %d, out: %d, moves: %d, deaths: %d, spells cast: %d\n", (logg.bytin), (logg.bytout), (logg.moves), (logg.killed), (logg.spused))
			lprintf("  out bytes per move: %d,  time per move: %d ms\n", (logg.bytout / logg.moves), ((logg.cputime * 1000) / logg.moves))
	#endif
		}
		lflush()
		lrclose()
	*/
}
