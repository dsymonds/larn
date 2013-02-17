package main

/* Keystrokes (roughly) between checkpoints */
const CHECKPOINT_INTERVAL = 400

var lastok = 0
var yrepcount = 0

const flushno = 5             /* input queue flushing threshold */
const MAXUM = 52              /* maximum number of user re-named monsters */
const MAXMNAME = 40           /* max length of a monster re-name */
var usermonster [MAXUM]string /* the user named monster name goes here */
var usermpoint byte = 0       /* the user monster pointer */

/*
	lexical analyzer for larn
*/
func yylex() int {
	if hit2flag {
		hit2flag = false
		yrepcount = 0
		return ' '
	}
	if yrepcount > 0 {
		yrepcount--
		return lastok
	} else {
		yrepcount = 0
	}
	if yrepcount == 0 {
		bottomdo()
		showplayer()
	} /* show where the player is	 */
	lflush()
	for {
		c[BYTESIN]++
		/* check for periodic checkpointing */
		if ckpflag {
			if (c[BYTESIN] % CHECKPOINT_INTERVAL) == 0 {
				savegame(ckpfile)
			}
		}
		var cc int
		// TODO
		//var ic int
		/*
			for {		// if keyboard input buffer is too big, flush some of it
				ioctl(0, FIONREAD, &ic);
				if (ic > flushno)
					read(0, &cc, 1);
				if ic <= flushno {
					break
				}
			}
		*/

		cc = ttgetch()

		if cc <= '9' && cc >= '0' {
			yrepcount = yrepcount*10 + cc - '0'
		} else {
			if yrepcount > 0 {
				yrepcount--
			}
			lastok = cc
			return lastok
		}
	}
	panic("unreachable")
}

/*
 *	flushall()		Function to flush all type-ahead in the input buffer
 */
func flushall() {
	// TODO
	/*
		char            cc;
		int             ic;
		for (;;) {		// if keyboard input buffer is too big, flush some of it
			ioctl(0, FIONREAD, &ic);
			if (ic <= 0)
				return;
			while (ic > 0) {
				read(0, &cc, 1);
				--ic;
			}		// gobble up the byte
		}
	*/
}

/*
	function to set the desired hardness
	enter with hard= -1 for default hardness, else any desired hardness
*/
func sethard(hard int) {
	j := c[HARDGAME]
	hashewon()
	if !restorflag { /* don't set c[HARDGAME] if restoring game */
		if hard >= 0 {
			c[HARDGAME] = hard
		}
	} else {
		c[HARDGAME] = j /* set c[HARDGAME] to proper value if restoring game */
	}

	k := c[HARDGAME]
	if k != 0 {
		for j = 0; j <= MAXMONST+8; j++ {
			mp := &monster[j]
			i := ((6+k)*mp.hitpoints + 1) / 6
			mp.hitpoints = icond(i < 0, 32767, i)
			i = ((6+k)*mp.damage + 1) / 5
			mp.damage = icond(i > 127, 127, i)
			i = (10 * mp.gold) / (10 + k)
			mp.gold = icond(i > 32767, 32767, i)
			i = mp.armorclass - k
			mp.armorclass = icond(i < -127, -127, i)
			i = (7*mp.experience)/(7+k) + 1
			mp.experience = icond(i <= 0, 1, i)
		}
	}
}

/*
	function to read and process the larn options file
*/
func readopts() {
	// TODO
	/*
		const char  *i;
		int    j, k;
		int             flag;

		flag = 1;		// set to 0 if a name is specified

		if !lopen(optsfile) {
			strcpy(logname, loginname);
			return;		// user name if no character name
		}
		i = " ";
		while (*i) {
			if ((i = lgetw()) == NULL)
				break;	// check for EOF
			while ((*i == ' ') || (*i == '\t'))
				i++;	// eat leading whitespace

			if (strcmp(i, "bold-objects") == 0)
				boldon = true
			else if (strcmp(i, "enable-checkpointing") == 0)
				ckpflag = true
			else if (strcmp(i, "inverse-objects") == 0)
				boldon = false
			else if (strcmp(i, "female") == 0)
				sex = 0;	// male or female
			else if (strcmp(i, "monster:") == 0) {	// name favorite monster
				if ((i = lgetw()) == 0)
					break;
				strlcpy(usermonster[usermpoint], i, MAXMNAME);
				if (usermpoint >= MAXUM)
					continue;	// defined all of em
				if (isalpha(j = usermonster[usermpoint][0])) {
					for (k = 1; k < MAXMONST + 8; k++)	// find monster
						if (monstnamelist[k] == j) {
							monster[k].name = &usermonster[usermpoint++][0];
							break;
						}
				}
			} else if (strcmp(i, "male") == 0)
				sex = 1;
			else if (strcmp(i, "name:") == 0) {	// defining players name
				if ((i = lgetw()) == 0)
					break;
				strlcpy(logname, i, LOGNAMESIZE);
				flag = 0;
			} else if (strcmp(i, "no-introduction") == 0)
				nowelcome = true
			else if (strcmp(i, "no-beep") == 0)
				nobeep = true
			else if (strcmp(i, "process-name:") == 0) {
				if ((i = lgetw()) == 0)
					break;
				strlcpy(psname, i, PSNAMESIZE);
			} else if (strcmp(i, "play-day-play") == 0) {
				// bypass time restrictions: ignored
			} else if (strcmp(i, "savefile:") == 0) {	// defining savefilename
				if ((i = lgetw()) == 0)
					break;
				strcpy(savefilename, i);
				flag = 0;
			}
		}
		if (flag)
			strcpy(logname, loginname);
	*/
}
