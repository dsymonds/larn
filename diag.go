package main

import (
	"log"
)

/*
	***************************
	DIAG -- dungeon diagnostics
	***************************

	subroutine to print out data for debugging
*/
func diag() int {
	cursors()
	lwclose()
	if lcreat(diagfile) < 0 { /* open the diagnostic file	 */
		lcreat("")
		lprcat("\ndiagnostic failure\n")
		return -1
	}
	log.Print("Diagnosing . . .")
	lprcat("\n\nBeginning of DIAG diagnostics ----------\n")

	/* for the character attributes	 */

	lprintf("\n\nPlayer attributes:\n\nHit points: %2d(%2d)", c[HP], c[HPMAX])
	lprintf("\ngold: %d  Experience: %d  Character level: %d  Level in caverns: %d",
		c[GOLD], c[EXPERIENCE], c[LEVEL], level)
	lprintf("\nTotal types of monsters: %d", MAXMONST+8)

	lprcat("\f\nHere's the dungeon:\n\n")

	i := level
	for j := 0; j < MAXLEVEL+MAXVLEVEL; j++ {
		newcavelevel(j)
		lprintf("\nMaze for level %s:\n", levelname[level])
		diagdrawscreen()
	}
	newcavelevel(i)

	lprcat("\f\nNow for the monster data:\n\n")
	lprcat("   Monster Name      LEV  AC   DAM  ATT  DEF    GOLD   HP     EXP   \n")
	lprcat("--------------------------------------------------------------------------\n")
	for i := 0; i <= MAXMONST+8; i++ {
		lprintf("%19s  %2d  %3d ", monster[i].name, monster[i].level, monster[i].armorclass)
		lprintf(" %3d  %3d  %3d  ", monster[i].damage, monster[i].attack, monster[i].defense)
		lprintf("%6d  %3d   %6d\n", monster[i].gold, monster[i].hitpoints, monster[i].experience)
	}

	lprcat("\n\nHere's a Table for the to hit percentages\n")
	lprcat("\n     We will be assuming that players level = 2 * monster level")
	lprcat("\n     and that the players dexterity and strength are 16.")
	lprcat("\n    to hit: if (rnd(22) < (2[monst AC] + your level + dex + WC/8 -1)/2) then hit")
	lprcat("\n    damage = rund(8) + WC/2 + STR - c[HARDGAME] - 4")
	lprcat("\n    to hit:  if rnd(22) < to hit  then player hits\n")
	lprcat("\n    Each entry is as follows:  to hit / damage / number hits to kill\n")
	lprcat("\n          monster     WC = 4         WC = 20        WC = 40")
	lprcat("\n---------------------------------------------------------------")
	for i = 0; i <= MAXMONST+8; i++ {
		hit := 2*monster[i].armorclass + 2*monster[i].level + 16
		dam := 16 - c[HARDGAME]
		lprintf("\n%20s   %2d/%2d/%2d       %2d/%2d/%2d       %2d/%2d/%2d",
			monster[i].name,
			(hit / 2), max(0, dam+2), (monster[i].hitpoints/(dam+2) + 1),
			((hit + 2) / 2), max(0, dam+10), (monster[i].hitpoints/(dam+10) + 1),
			((hit + 5) / 2), max(0, dam+20), (monster[i].hitpoints/(dam+20) + 1))
	}

	lprcat("\n\nHere's the list of available potions:\n\n")
	for i = 0; i < MAXPOTION; i++ {
		lprintf("%20s\n", potionhide[i][1])
	}
	lprcat("\n\nHere's the list of available scrolls:\n\n")
	for i = 0; i < MAXSCROLL; i++ {
		lprintf("%20s\n", scrollhide[i][1])
	}
	lprcat("\n\nHere's the spell list:\n\n")
	lprcat("spell          name           description\n")
	lprcat("-------------------------------------------------------------------------------------------\n\n")
	for j := 0; j < SPNUM; j++ {
		lprc(' ')
		lprcat(spelcode[j])
		lprintf(" %21s  %s\n", spelname[j], speldescript[j])
	}

	lprcat("\n\nFor the c[] array:\n")
	for j := 0; j < 100; j += 10 {
		lprintf("\nc[%2d] = ", j)
		for i = 0; i < 9; i++ {
			lprintf("%5d ", c[i+j])
		}
	}

	lprcat("\n\nTest of random number generator ----------------")
	lprcat("\n    for 25,000 calls divided into 16 slots\n\n")

	var rndcount [16]int
	for i = 0; i < 16; i++ {
		rndcount[i] = 0
	}
	for i = 0; i < 25000; i++ {
		rndcount[rund(16)]++
	}
	for i = 0; i < 16; i++ {
		lprintf("  %5d", rndcount[i])
		if i == 7 {
			lprc('\n')
		}
	}

	lprcat("\n\n")
	lwclose()
	lcreat("")
	lprcat("Done Diagnosing . . .")
	return 0
}

/*
	subroutine to draw the whole screen as the player knows it
*/
func diagdrawscreen() {
	for i := 0; i < MAXY; i++ {
		/* for the east west walls of this line	 */
		for j := 0; j < MAXX; j++ {
			k := mitem[j][i]
			if k != 0 {
				lprc(monstnamelist[k])
			} else {
				lprc(objnamelist[item[j][i]])
			}
		}
		lprc('\n')
	}
}

/*
	to save the game in a file
*/
func savegame(fname string) int {
	//int    i, k;
	//struct sphere *sp;
	//struct stat     statbuf;

	// TODO
	return -1
	/*
		nosignal = 1
		lflush()
		savelevel()
		ointerest()
		if lcreat(fname) < 0 {
			lcreat("")
			lprintf("\nCan't open file <%s> to save game\n", fname)
			nosignal = 0
			return -1
		}
		set_score_output()
		lwrite(beenhere, MAXLEVEL+MAXVLEVEL)
		for k := 0; k < MAXLEVEL+MAXVLEVEL; k++ {
			if beenhere[k] {
				lwrite((char *) &cell[k * MAXX * MAXY], sizeof(struct cel) * MAXY * MAXX)
			}
		}
		struct tms cputime
		times(&cputime);	// get cpu time
		c[CPUTIME] += (cputime.tms_utime + cputime.tms_stime) / 60
		lwrite((char *) &c[0], 100 * sizeof(long))
		lprint(gltime)
		lprc(level)
		lprc(playerx)
		lprc(playery)
		lwrite(iven, 26)
		lwrite(ivenarg, 26*sizeof(short))
		for k := 0; k < MAXSCROLL; k++ {
			lprc(scrollname[k][0])
		}
		for k := 0; k < MAXPOTION; k++ {
			lprc(potionname[k][0])
		}
		lwrite(spelknow, SPNUM)
		lprc(wizard)
		lprc(rmst) // random monster generation counter
		for i := 0; i < 90; i++ {
			lprc(itm[i].qty)
		}
		lwrite(course, 25)
		lprc(cheat)
		lprc(VERSION)
		for i := 0; i < MAXMONST; i++ {
			lprc(monster[i].genocided) // genocide info
		}
		for (sp = spheres; sp; sp = sp->p)
			lwrite((char *) sp, sizeof(struct sphere));	// save spheres of annihilation
		time(&zzz);
		lprint((long) (zzz - initialtime));
		lwrite((char *) &zzz, sizeof(long));
		if (fstat(io_outfd, &statbuf) < 0)
			lprint(0L);
		else
			lprint((long) statbuf.st_ino);	// inode #
		lwclose()
		lastmonst = ""
		lcreat("")
		nosignal = 0
	*/
	return 0
}

func restoregame(fname string) {
	// TODO
	/*
		int    i, k;
		struct sphere *sp, *sp2;
		struct stat     filetimes;
		cursors();
		lprcat("\nRestoring . . .");
		lflush();
		if (lopen(fname) <= 0) {
			lcreat((char *) 0);
			lprintf("\nCan't open file <%s>to restore game\n", fname);
			nap(2000);
			c[GOLD] = c[BANKACCOUNT] = 0;
			died(-265);
			return;
		}
		lrfill((char *) beenhere, MAXLEVEL + MAXVLEVEL);
		for (k = 0; k < MAXLEVEL + MAXVLEVEL; k++)
			if (beenhere[k])
				lrfill((char *) &cell[k * MAXX * MAXY], sizeof(struct cel) * MAXY * MAXX);

		lrfill((char *) &c[0], 100 * sizeof(long));
		gltime = larn_lrint();
		level = c[CAVELEVEL] = lgetc();
		playerx = lgetc();
		playery = lgetc();
		lrfill((char *) iven, 26);
		lrfill((char *) ivenarg, 26 * sizeof(short));
		for (k = 0; k < MAXSCROLL; k++)
			scrollname[k] = lgetc() ? scrollhide[k] : "";
		for (k = 0; k < MAXPOTION; k++)
			potionname[k] = lgetc() ? potionhide[k] : "";
		lrfill((char *) spelknow, SPNUM);
		wizard = lgetc();
		rmst = lgetc();		// random monster creation flag

		for (i = 0; i < 90; i++)
			itm[i].qty = lgetc();
		lrfill((char *) course, 25);
		cheat = lgetc();
		if (VERSION != lgetc()) {	// version number
			cheat = 1;
			lprcat("Sorry, But your save file is for an older version of larn\n");
			nap(2000);
			c[GOLD] = c[BANKACCOUNT] = 0;
			died(-266);
			return;
		}
		for (i = 0; i < MAXMONST; i++)
			monster[i].genocided = lgetc();	// genocide info
		for (sp = 0, i = 0; i < c[SPHCAST]; i++) {
			sp2 = sp;
			sp = (struct sphere *) malloc(sizeof(struct sphere));
			if (sp == 0) {
				write(2, "Can't malloc() for sphere space\n", 32);
				break;
			}
			lrfill((char *) sp, sizeof(struct sphere));	// get spheres of annihilation
			sp->p = 0;	// null out pointer
			if (i == 0)
				spheres = sp;	// beginning of list
			else
				sp2->p = sp;
		}

		time(&zzz);
		initialtime = zzz - larn_lrint();
		// get the creation and modification time of file
		fstat(io_infd, &filetimes);
		lrfill((char *) &zzz, sizeof(long));
		zzz += 6;
		if (filetimes.st_ctime > zzz)
			fsorry();	// file create time
		else if (filetimes.st_mtime > zzz)
			fsorry();	// file modify time
		if (c[HP] < 0) {
			died(284);
			return;
		}			// died a post mortem death
		oldx = oldy = 0;
		// XXX the following will break on 64-bit inode numbers
		i = larn_lrint();		// inode #
		if (i && (filetimes.st_ino != (ino_t) i))
			fsorry();
		lrclose();
		if (strcmp(fname, ckpfile) == 0) {
			if (lappend(fname) < 0)
				fcheat();
			else {
				lprc(' ');
				lwclose();
			}
			lcreat((char *) 0);
		} else if (unlink(fname) < 0)
			fcheat();	// can't unlink save file
		// for the greedy cheater checker
		for (k = 0; k < 6; k++)
			if (c[k] > 99)
				greedy();
		if (c[HPMAX] > 999 || c[SPELLMAX] > 125)
			greedy();
		if (c[LEVEL] == 25 && c[EXPERIENCE] > skill[24]) {	// if patch up lev 25 player
			long            tmp;
			tmp = c[EXPERIENCE] - skill[24];	// amount to go up
			c[EXPERIENCE] = skill[24];
			raiseexperience((long) tmp);
		}
		getlevel();
		lasttime = gltime;
	*/
}

/*
	subroutine to not allow greedy cheaters
*/
func greedy() {
	if WIZID && wizard {
		return
	}

	lprcat("\n\nI am so sorry, but your character is a little TOO good!  Since this\n")
	lprcat("cannot normally happen from an honest game, I must assume that you cheated.\n")
	lprcat("In that you are GREEDY as well as a CHEATER, I cannot allow this game\n")
	lprcat("to continue.\n")
	nap(5000)
	c[GOLD], c[BANKACCOUNT] = 0, 0
	died(-267)
}

/*
	subroutine to not allow altered save files and terminate the attempted
	restart
*/
func fsorry() {
	lprcat("\nSorry, but your savefile has been altered.\n")
	lprcat("However, seeing as I am a good sport, I will let you play.\n")
	lprcat("Be advised though, you won't be placed on the normal scoreboard.")
	cheat = true
	nap(4000)
}

/*
	subroutine to not allow game if save file can't be deleted
*/
func fcheat() {
	if WIZID && wizard {
		return
	}

	lprcat("\nSorry, but your savefile can't be deleted.  This can only mean\n")
	lprcat("that you tried to CHEAT by protecting the directory the savefile\n")
	lprcat("is in.  Since this is unfair to the rest of the larn community, I\n")
	lprcat("cannot let you play this game.\n")
	nap(5000)
	c[GOLD], c[BANKACCOUNT] = 0, 0
	died(-268)
}
