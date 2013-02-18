package main

import (
	"flag"
	"log"
	"os"
	"time"
)

const copyright = "\nLarn is copyrighted 1986 by Noah Morgan.\n"

var srcount = 0            /* line counter for showstr()	 */
var dropflag = 0           /* if 1 then don't lookforobject() next round */
var rmst = 80              /* random monster creation counter		 */
var userid int             /* the players login user id number */
var nowelcome, nomove bool /* if (nomove) then don't count next iteration as a move */
var viewflag int8
var restorflag = false /* whether restore has been done	 */

/*
 * if viewflag then we have done a 99 stay here and don't showcell in the
 * main loop
 */

var (
	replay = flag.String("replay", "", "if non-empty, a replay file to use")
	seed   = flag.Uint("seed", 0, "if non-zero, the random seed to use")
)

/*
const cmdhelp = `Cmd line format: larn [-slicnh] [-o<optsfile>] [-##] [++]
  -s   show the scoreboard
  -l   show the logfile (wizard id only)
  -i   show scoreboard with inventories of dead characters
  -c   create new scoreboard (wizard id only)
  -n   suppress welcome message on starting game
  -##  specify level of difficulty (example: -5)
  -h   print this help text
  ++   restore game from checkpoint file
  -o<optsfile>   specify .larnopts filename to be used instead of \"~/.larnopts\"
`
*/

//#ifdef VT100
//static char    *termtypes[] = {"vt100", "vt101", "vt102", "vt103", "vt125",
//	"vt131", "vt140", "vt180", "vt220", "vt240", "vt241", "vt320", "vt340", "vt341"};
//#endif	/* VT100 */
/*
	************
	MAIN PROGRAM
	************
*/
func main() {
	flag.Parse()

	// In case a panic occurs, be prepared to clean up the terminal.
	defer func() {
		debugf("caught panic!")
		if err := recover(); err != nil {
			clearvt100()
			panic(err) // re-panic
		}
	}()

	/*
	 *	first task is to identify the player
	 */
	//#ifndef VT100
	init_term() /* setup the terminal (find out what type)
	 * for termcap */
	//#endif	/* VT100 */
	/* try to get login name */
	// TODO: C version tried getlogin and getpwuid first
	ptr := os.Getenv("USER")
	if ptr == "" {
		ptr = os.Getenv("LOGNAME")
	}
	if ptr == "" {
		log.Fatal("Can't find your logname.  Who Are You?")
	}
	/*
	 *	second task is to prepare the pathnames the player will need
	 */
	loginname = ptr /* save loginname of the user for logging purposes */
	logname = ptr   /* this will be overwritten with the players name */
	ptr = os.Getenv("HOME")
	if ptr == "" {
		ptr = "."
	}
	savefilename = ptr
	savefilename = "/Larn.sav" /* save file name in home directory */
	optsfile = ptr + "/.larnopts"
	/* the .larnopts filename */

	lcreat("")
	seed := uint32(*seed)
	if seed == 0 {
		seed = uint32(time.Now().Unix())
	}
	newgame(seed) /* set the initial clock  */
	hard := -1

	//#ifdef VT100
	/*
	 *	check terminal type to avoid users who have not vt100 type terminals
	 */
	/*
		ttype = getenv("TERM");
		for (j = 1, i = 0; i < sizeof(termtypes) / sizeof(char *); i++)
			if (strcmp(ttype, termtypes[i]) == 0) {
				j = 0;
				break;
			}
		if (j) {
			lprcat("Sorry, Larn needs a VT100 family terminal for all its features.\n");
			lflush();
			exit(1);
		}
	*/
	//#endif	/* VT100 */

	/*
	 *	now make scoreboard if it is not there (don't clear)
	 */
	if !exists(scorefile) { /* not there */
		makeboard()
	}

	/*
	 *	now process the command line arguments
	 */
	// TODO: replace with Go's flag parser
	/*
		for i := 1; i < argc; i++ {
			if argv[i][0] == '-' {
				switch argv[i][1] {
				case 's':
					showscores();
					exit(0);	// show scoreboard

				case 'l':	// show log file
					diedlog();
					exit(0);

				case 'i':
					showallscores();
					exit(0);	// show all scoreboard

				case 'c':	// anyone with password can create scoreboard
					lprcat("Preparing to initialize the scoreboard.\n");
					if getpassword() {	// make new scoreboard
						makeboard();
						lprc('\n');
						showscores();
					}
					exit(0);

				case 'n':	// no welcome msg
					nowelcome = true
					argv[i][0] = 0;
					break;

				case '0':
				case '1':
				case '2':
				case '3':
				case '4':
				case '5':
				case '6':
				case '7':
				case '8':
				case '9':	// for hardness
					sscanf(&argv[i][1], "%d", &hard);
					break;

				case 'h':	// print out command line arguments
					write(1, cmdhelp, sizeof(cmdhelp));
					exit(0);

				case 'o':	// specify a .larnopts filename
					strncpy(optsfile, argv[i] + 2, 127);
					break;

				default:
					printf("Unknown option <%s>\n", argv[i]);
					exit(1);
				};
			}

			if (argv[i][0] == '+') {
				clear();
				restorflag = true
				if (argv[i][1] == '+') {
					hitflag = true
					restoregame(ckpfile);	// restore checkpointed game
				}
				i = argc;
			}
		}
	*/

	readopts() /* read the options file if there is one */

	userid = os.Geteuid() /* obtain the user's effective id number */
	if userid < 0 {
		log.Fatal("Can't obtain playerid")
	}

	if *replay != "" {
		loadReplay(*replay)
	}

	if exists(savefilename) { /* restore game if need to */
		clear()
		restorflag = true
		hitflag = true
		restoregame(savefilename) /* restore last game	 */
	}
	sigsetup()      /* trap all needed signals	 */
	sethard(hard)   /* set up the desired difficulty				 */
	setupvt100()    /* setup the terminal special mode				 */
	if c[HP] == 0 { /* create new game */
		makeplayer()    /* make the character that will play			 */
		newcavelevel(0) /* make the dungeon						 	 */
		predostuff = 1  /* tell signals that we are in the welcome screen */
		if !nowelcome {
			welcome() /* welcome the player to the game */
		}
	}
	drawscreen()   /* show the initial dungeon					 */
	predostuff = 2 /* tell the trap functions that they must do a showplayer() from here on */

	//nice(1);		/* games should be run niced */

	yrepcount, hit2flag = 0, false
	for {
		if dropflag == 0 {
			lookforobject() /* see if there is an object here	 */
		} else {
			dropflag = 0 /* don't show it just dropped an item */
		}
		if !hitflag {
			if c[HASTEMONST] != 0 {
				movemonst()
			}
			movemonst()
		} /* move the monsters		 */
		if viewflag == 0 {
			showcell(playerx, playery)
		} else {
			viewflag = 0 /* show stuff around player	 */
		}
		if hit3flag {
			flushall()
		}
		hitflag, hit3flag = false, false
		nomove = true
		bot_linex() /* update bottom line */
		for nomove {
			if hit3flag {
				flushall()
			}
			nomove = false
			parse()
		} /* get commands and make moves	 */
		regen() /* regenerate hp and spells			 */
		if c[TIMESTOP] == 0 {
			rmst--
			if rmst <= 0 {
				rmst = 120 - (level << 2)
				fillmonst(makemonst(level))
			}
		}
	}
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

/*
	showstr()

	show character's inventory
*/
func showstr() {
	number := 3
	for i := 0; i < 26; i++ {
		if iven[i] != 0 {
			number++ /* count items in inventory */
		}
	}
	t_setup(number)
	qshowstr()
	t_endup(number)
}

func qshowstr() {
	srcount = 0
	sigsav := nosignal
	nosignal = true /* don't allow ^c etc */
	if c[GOLD] != 0 {
		lprintf(".)   %d gold pieces", c[GOLD])
		srcount++
	}
	for k := len(iven) - 1; k >= 0; k-- {
		if iven[k] != 0 {
			for i := 22; i < 84; i++ {
				for j := 0; j <= k; j++ {
					if i == iven[j] {
						show3(j)
					}
				}
			}
			k = 0
		}
	}
	lprintf("\nElapsed time is %d.  You have %d mobuls left", (gltime+99)/100+1, (TIMELIMIT-gltime)/100)
	more()
	nosignal = sigsav
}

/*
 *	subroutine to clear screen depending on # lines to display
 */
func t_setup(count int) {
	if count < 20 { /* how do we clear the screen? */
		cl_up(79, count)
		cursor(1, 1)
	} else {
		resetscroll()
		clear()
	}
}

/*
 *	subroutine to restore normal display screen depending on t_setup()
 */
func t_endup(count int) {
	if count < 18 { /* how did we clear the screen? */
		draws(0, MAXX, 0, icond(count > MAXY, MAXY, count))
	} else {
		drawscreen()
		setscroll()
	}
}

/*
	function to show the things player is wearing only
*/
func showwear() {
	sigsav := nosignal
	nosignal = true /* don't allow ^c etc */
	srcount = 0

	count := 2
	for _, i := range iven { /* count number of items we will display */
		if i != 0 {
			switch i {
			case OLEATHER, OPLATE, OCHAIN, ORING, OSTUDLEATHER, OSPLINT, OPLATEARMOR, OSSPLATE, OSHIELD:
				count++
			}
		}
	}

	t_setup(count)

	for i := 22; i < 84; i++ {
		for j, it := range iven {
			if i == it {
				switch i {
				case OLEATHER, OPLATE, OCHAIN, ORING, OSTUDLEATHER, OSPLINT, OPLATEARMOR, OSSPLATE, OSHIELD:
					show3(j)
				}
			}
		}
	}
	more()
	nosignal = sigsav
	t_endup(count)
}

/*
	function to show the things player can wield only
*/
func showwield() {
	sigsav := nosignal
	nosignal = true /* don't allow ^c etc */
	srcount = 0

	count := 2
	for _, i := range iven { /* count how many items */
		if i != 0 {
			switch i {
			case ODIAMOND, ORUBY, OEMERALD, OSAPPHIRE, OBOOK, OCHEST, OLARNEYE, ONOTHEFT, OSPIRITSCARAB, OCUBEofUNDEAD, OPOTION, OSCROLL:
			default:
				count++
			}
		}
	}

	t_setup(count)

	for i := 22; i < 84; i++ {
		for j, it := range iven {
			if i == it {
				switch i {
				case ODIAMOND, ORUBY, OEMERALD, OSAPPHIRE, OBOOK, OCHEST, OLARNEYE, ONOTHEFT, OSPIRITSCARAB, OCUBEofUNDEAD, OPOTION, OSCROLL:
				default:
					show3(j)
				}
			}
		}
	}
	more()
	nosignal = sigsav
	t_endup(count)
}

/*
 *	function to show the things player can read only
 */
func showread() {
	sigsav := nosignal
	nosignal = true /* don't allow ^c etc */
	srcount = 0

	count := 2
	for _, it := range iven {
		switch it {
		case OBOOK, OSCROLL:
			count++
		}
	}
	t_setup(count)

	for i := 22; i < 84; i++ {
		for j, it := range iven {
			if i == it {
				switch i {
				case OBOOK, OSCROLL:
					show3(j)
				}
			}
		}
	}
	more()
	nosignal = sigsav
	t_endup(count)
}

/*
 *	function to show the things player can eat only
 */
func showeat() {
	sigsav := nosignal
	nosignal = true /* don't allow ^c etc */
	srcount = 0

	count := 2
	for _, it := range iven {
		switch it {
		case OCOOKIE:
			count++
		}
	}
	t_setup(count)

	for i := 22; i < 84; i++ {
		for j, it := range iven {
			if i == it {
				switch i {
				case OCOOKIE:
					show3(j)
				}
			}
		}
	}
	more()
	nosignal = sigsav
	t_endup(count)
}

/*
	function to show the things player can quaff only
*/
func showquaff() {
	sigsav := nosignal
	nosignal = true /* don't allow ^c etc */
	srcount = 0

	count := 2
	for _, it := range iven {
		switch it {
		case OPOTION:
			count++
		}
	}
	t_setup(count)

	for i := 22; i < 84; i++ {
		for j, it := range iven {
			if i == it {
				switch i {
				case OPOTION:
					show3(j)
				}
			}
		}
	}
	more()
	nosignal = sigsav
	t_endup(count)
}

func show1(idx int, str2 []string) {
	lprintf("\n%c)   %s", idx+'a', objectname[iven[idx]])
	if len(str2) != 0 && str2[ivenarg[idx]] != "" {
		lprintf(" of%s", str2[ivenarg[idx]])
	}
}

func show3(indx int) {
	switch iven[indx] {
	case OPOTION:
		show1(indx, potionname)
	case OSCROLL:
		show1(indx, scrollname)

	case OLARNEYE, OBOOK, OSPIRITSCARAB, ODIAMOND, ORUBY, OCUBEofUNDEAD, OEMERALD, OCHEST, OCOOKIE, OSAPPHIRE, ONOTHEFT:
		show1(indx, nil)

	default:
		lprintf("\n%c)   %s", indx+'a', objectname[iven[indx]])
		if ivenarg[indx] > 0 {
			lprintf(" + %d", ivenarg[indx])
		} else if ivenarg[indx] < 0 {
			lprintf(" %d", ivenarg[indx])
		}
	}
	if c[WIELD] == indx {
		lprcat(" (weapon in hand)")
	}
	if c[WEAR] == indx || c[SHIELD] == indx {
		lprcat(" (being worn)")
	}
	srcount++
	if srcount >= 22 {
		srcount = 0
		more()
		clear()
	}
}

/*
	subroutine to randomly create monsters if needed
*/
func randmonst() {
	if c[TIMESTOP] != 0 {
		return /* don't make monsters if time is stopped	 */
	}
	rmst--
	if rmst <= 0 {
		rmst = 120 - (level << 2)
		fillmonst(makemonst(level))
	}
}

/*
	parse()

	get and execute a command
*/
func parse() {
	for {
		k := yylex()
		switch k { /* get the token from the input and switch on it	 */
		case 'h':
			moveplayer(4)
			return /* west		 */
		case 'H':
			run(4)
			return /* west		 */
		case 'l':
			moveplayer(2)
			return /* east		 */
		case 'L':
			run(2)
			return /* east		 */
		case 'j':
			moveplayer(1)
			return /* south		 */
		case 'J':
			run(1)
			return /* south		 */
		case 'k':
			moveplayer(3)
			return /* north		 */
		case 'K':
			run(3)
			return /* north		 */
		case 'u':
			moveplayer(5)
			return /* northeast	 */
		case 'U':
			run(5)
			return /* northeast	 */
		case 'y':
			moveplayer(6)
			return /* northwest	 */
		case 'Y':
			run(6)
			return /* northwest	 */
		case 'n':
			moveplayer(7)
			return /* southeast	 */
		case 'N':
			run(7)
			return /* southeast	 */
		case 'b':
			moveplayer(8)
			return /* southwest	 */
		case 'B':
			run(8)
			return /* southwest	 */

		case '.':
			if yrepcount > 0 {
				viewflag = 1
			}
			return /* stay here		 */

		case 'w':
			yrepcount = 0
			wield()
			return /* wield a weapon */

		case 'W':
			yrepcount = 0
			wear()
			return /* wear armor	 */

		case 'r':
			yrepcount = 0
			if c[BLINDCOUNT] != 0 {
				cursors()
				lprcat("\nYou can't read anything when you're blind!")
			} else if c[TIMESTOP] == 0 {
				readscr()
			}
			return /* to read a scroll	 */

		case 'q':
			yrepcount = 0
			if c[TIMESTOP] == 0 {
				quaff()
			}
			return /* quaff a potion		 */

		case 'd':
			yrepcount = 0
			if c[TIMESTOP] == 0 {
				dropobj()
			}
			return /* to drop an object	 */

		case 'c':
			yrepcount = 0
			cast()
			return /* cast a spell	 */

		case 'i':
			yrepcount = 0
			nomove = true
			showstr()
			return /* status		 */

		case 'e':
			yrepcount = 0
			if c[TIMESTOP] == 0 {
				eatcookie()
			}
			return /* to eat a fortune cookie */

		case 'D':
			yrepcount = 0
			seemagic(0)
			nomove = true
			return /* list spells and scrolls */

		case '?':
			yrepcount = 0
			help()
			nomove = true
			return /* give the help screen */

		case 'S':
			clear()
			lprcat("Saving . . .")
			lflush()
			savegame(savefilename)
			wizard = true
			died(-257) /* save the game - doesn't return	 */

		case 'Z':
			yrepcount = 0
			if c[LEVEL] > 9 {
				oteleport(1)
				return
			}
			cursors()
			lprcat("\nAs yet, you don't have enough experience to use teleportation")
			return /* teleport yourself	 */

		case '^': /* identify traps */
			flag := 0
			yrepcount = 0
			cursors()
			lprc('\n')
			for j := playery - 1; j < playery+2; j++ {
				if j < 0 {
					j = 0
				}
				if j >= MAXY {
					break
				}
				for i := playerx - 1; i < playerx+2; i++ {
					if i < 0 {
						i = 0
					}
					if i >= MAXX {
						break
					}
					switch item[i][j] {
					case OTRAPDOOR, ODARTRAP, OTRAPARROW, OTELEPORTER:
						lprcat("\nIt's ")
						lprcat(objectname[item[i][j]])
						flag++
					}
				}
			}
			if flag == 0 {
				lprcat("\nNo traps are visible")
			}
			return

		case '_': /* this is the fudge player password for wizard mode */
			yrepcount = 0
			cursors()
			nomove = true
			if !WIZID || userid != wisid {
				lprcat("Sorry, you are not empowered to be a wizard.\n")
				scbr() /* system("stty -echo cbreak"); */
				lflush()
				return
			}
			if !getpassword() {
				scbr() /* system("stty -echo cbreak"); */
				return
			}
			wizard = true
			scbr() /* system("stty -echo cbreak"); */
			for i := 0; i < 6; i++ {
				c[i] = 70
			}
			iven[0], iven[1] = 0, 0
			take(OPROTRING, 50)
			take(OLANCE, 25)
			c[WIELD] = 1
			c[LANCEDEATH] = 1
			c[WEAR], c[SHIELD] = -1, -1
			raiseexperience(6000000)
			c[AWARENESS] += 25000
			{
				for i := 0; i < MAXY; i++ {
					for j := 0; j < MAXX; j++ {
						know[j][i] = true
					}
				}
				for i := 0; i < SPNUM; i++ {
					spelknow[i] = true
				}
				for i := 0; i < MAXSCROLL; i++ {
					scrollname[i] = scrollhide[i]
				}
				for i := 0; i < MAXPOTION; i++ {
					potionname[i] = potionhide[i]
				}
			}
			for i := 0; i < MAXSCROLL; i++ {
				if len(scrollname[i]) > 2 { /* no null items */
					item[i][0] = OSCROLL
					iarg[i][0] = i
				}
			}
			for i := MAXX - 1; i > MAXX-1-MAXPOTION; i-- {
				if len(potionname[i-MAXX+MAXPOTION]) > 2 { /* no null items */
					item[i][0] = OPOTION
					iarg[i][0] = i - MAXX + MAXPOTION
				}
			}
			for i := 1; i < MAXY; i++ {
				item[0][i] = i
				iarg[0][i] = 0
			}
			for i := MAXY; i < MAXY+MAXX; i++ {
				item[i-MAXY][MAXY-1] = i
				iarg[i-MAXY][MAXY-1] = 0
			}
			for i := MAXX + MAXY; i < MAXX+MAXY+MAXY; i++ {
				item[MAXX-1][i-MAXX-MAXY] = i
				iarg[MAXX-1][i-MAXX-MAXY] = 0
			}
			c[GOLD] += 25000
			drawscreen()
			return

		case 'T':
			yrepcount = 0
			cursors()
			if c[SHIELD] != -1 {
				c[SHIELD] = -1
				lprcat("\nYour shield is off")
				bottomline()
			} else if c[WEAR] != -1 {
				c[WEAR] = -1
				lprcat("\nYour armor is off")
				bottomline()
			} else {
				lprcat("\nYou aren't wearing anything")
			}
			return

		case 'g':
			cursors()
			lprintf("\nThe stuff you are carrying presently weighs %d pounds", packweight())
		case ' ':
			yrepcount = 0
			nomove = true
			return

		case 'v':
			yrepcount = 0
			cursors()
			lprintf("\nCaverns of Larn, Version %d.%d, Diff=%d",
				VERSION, SUBVERSION, c[HARDGAME])
			if wizard {
				lprcat(" Wizard")
			}
			nomove = true
			if cheat {
				lprcat(" Cheater")
			}
			lprcat(copyright)
			return

		case 'Q':
			yrepcount = 0
			quit()
			nomove = true
			return /* quit		 */

		case 'L' - 64:
			yrepcount = 0
			drawscreen()
			nomove = true
			return /* look		 */

		case 'A':
			yrepcount = 0
			nomove = true
			if WIZID && wizard {
				diag()
				return
			} /* create diagnostic file */
			return

		case 'P':
			cursors()
			if outstanding_taxes > 0 {
				lprintf("\nYou presently owe %d gp in taxes.",
					outstanding_taxes)
			} else {
				lprcat("\nYou do not owe any taxes.")
			}
			return
		}
	}
}

func parse2() {
	if c[HASTEMONST] != 0 {
		movemonst()
	}
	movemonst() /* move the monsters		 */
	randmonst()
	regen()
}

func run(dir int) {
	i := 1
	for i != 0 {
		i = moveplayer(dir)
		if i > 0 {
			if c[HASTEMONST] != 0 {
				movemonst()
			}
			movemonst()
			randmonst()
			regen()
		}
		if hitflag {
			i = 0
		}
		if i != 0 {
			showcell(playerx, playery)
		}
	}
}

/*
	function to wield a weapon
*/
func wield() {
	for {
		i := whatitem("wield")
		if i == '\033' {
			return
		}
		if i != '.' {
			if i == '*' {
				showwield()
			} else if iven[i-'a'] == 0 {
				ydhi(i)
				return
			} else if iven[i-'a'] == OPOTION {
				ycwi(i)
				return
			} else if iven[i-'a'] == OSCROLL {
				ycwi(i)
				return
			} else if c[SHIELD] != -1 && iven[i-'a'] == O2SWORD {
				lprcat("\nBut one arm is busy with your shield!")
				return
			} else {
				c[WIELD] = i - 'a'
				if iven[i-'a'] == OLANCE {
					c[LANCEDEATH] = 1
				} else {
					c[LANCEDEATH] = 0
				}
				bottomline()
				return
			}
		}
	}
}

/*
	common routine to say you don't have an item
*/
func ydhi(x int) {
	cursors()
	lprintf("\nYou don't have item %c!", x)
}
func ycwi(x int) {
	cursors()
	lprintf("\nYou can't wield item %c!", x)
}

/*
	function to wear armor
*/
func wear() {
	for {
		i := whatitem("wear")
		if i == '\033' {
			return
		}
		if i != '.' {
			if i == '*' {
				showwear()
			} else {
				switch iven[i-'a'] {
				case 0:
					ydhi(i)
					return
				case OLEATHER, OCHAIN, OPLATE, OSTUDLEATHER, ORING, OSPLINT, OPLATEARMOR, OSSPLATE:
					if c[WEAR] != -1 {
						lprcat("\nYou're already wearing some armor")
						return
					}
					c[WEAR] = i - 'a'
					bottomline()
					return
				case OSHIELD:
					if c[SHIELD] != -1 {
						lprcat("\nYou are already wearing a shield")
						return
					}
					if iven[c[WIELD]] == O2SWORD {
						lprcat("\nYour hands are busy with the two handed sword!")
						return
					}
					c[SHIELD] = i - 'a'
					bottomline()
					return
				default:
					lprcat("\nYou can't wear that!")
				}
			}
		}
	}
}

/*
	function to drop an object
*/
func dropobj() {
	p := &item[playerx][playery]
	for {
		i := whatitem("drop")
		if i == '\033' {
			return
		}
		if i == '*' {
			showstr()
		} else {
			if i == '.' { /* drop some gold */
				if *p != 0 {
					lprcat("\nThere's something here already!")
					return
				}
				lprcat("\n\n")
				cl_dn(1, 23)
				lprcat("How much gold do you drop? ")
				amt := readnum(c[GOLD])
				if amt == 0 {
					return
				}
				if amt > c[GOLD] {
					lprcat("\nYou don't have that much!")
					return
				}
				if amt <= 32767 {
					*p = OGOLDPILE
					i = amt
				} else if amt <= 327670 {
					*p = ODGOLD
					i = amt / 10
					amt = 10 * i
				} else if amt <= 3276700 {
					*p = OMAXGOLD
					i = amt / 100
					amt = 100 * i
				} else if amt <= 32767000 {
					*p = OKGOLD
					i = amt / 1000
					amt = 1000 * i
				} else {
					*p = OKGOLD
					i = 32767
					amt = 32767000
				}
				c[GOLD] -= amt
				lprintf("You drop %d gold pieces", amt)
				iarg[playerx][playery] = i
				bottomgold()
				know[playerx][playery] = false
				dropflag = 1
				return
			}
			drop_object(i - 'a')
			return
		}
	}
}

/*
 *	readscr()		Subroutine to read a scroll one is carrying
 */
func readscr() {
	for {
		i := whatitem("read")
		if i == '\033' {
			return
		}
		if i != '.' {
			if i == '*' {
				showread()
			} else {
				if iven[i-'a'] == OSCROLL {
					read_scroll(ivenarg[i-'a'])
					iven[i-'a'] = 0
					return
				}
				if iven[i-'a'] == OBOOK {
					readbook(ivenarg[i-'a'])
					iven[i-'a'] = 0
					return
				}
				if iven[i-'a'] == 0 {
					ydhi(i)
					return
				}
				lprcat("\nThere's nothing on it to read")
				return
			}
		}
	}
}

/*
 *	subroutine to eat a cookie one is carrying
 */
func eatcookie() {
	for {
		i := whatitem("eat")
		if i == '\033' {
			return
		}
		if i != '.' {
			if i == '*' {
				showeat()
			} else {
				if iven[i-'a'] == OCOOKIE {
					lprcat("\nThe cookie was delicious.")
					iven[i-'a'] = 0
					if c[BLINDCOUNT] == 0 {
						p := fortune()
						if p != "" {
							lprcat("  Inside you find a scrap of paper that says:\n")
							lprcat(p)
						}
					}
					return
				}
				if iven[i-'a'] == 0 {
					ydhi(i)
					return
				}
				lprcat("\nYou can't eat that!")
				return
			}
		}
	}
}

/*
 *	subroutine to quaff a potion one is carrying
 */
func quaff() {
	for {
		i := whatitem("quaff")
		if i == '\033' {
			return
		}
		if i != '.' {
			if i == '*' {
				showquaff()
			} else {
				if iven[i-'a'] == OPOTION {
					quaffpotion(ivenarg[i-'a'])
					iven[i-'a'] = 0
					return
				}
				if iven[i-'a'] == 0 {
					ydhi(i)
					return
				}
				lprcat("\nYou wouldn't want to quaff that, would you? ")
				return
			}
		}
	}
}

/*
	function to ask what player wants to do
*/
func whatitem(str string) int {
	cursors()
	lprintf("\nWhat do you want to %s [* for all] ? ", str)
	i := 0
	for i > 'z' || (i < 'a' && i != '*' && i != '\033' && i != '.') {
		i = ttgetch()
	}
	if i == '\033' {
		lprcat(" aborted")
	}
	return i
}

/*
	subroutine to get a number from the player
	and allow * to mean return amt, else return the number entered
*/
func readnum(mx int) int {
	var amt int
	sncbr()
	i := ttgetch()
	if i == '*' {
		amt = mx /* allow him to say * for all gold */
	} else {
		for i != '\n' {
			if i == '\033' {
				scbr()
				lprcat(" aborted")
				return 0
			}
			if i <= '9' && i >= '0' && amt < 99999999 {
				amt = amt*10 + i - '0'
			}
			i = ttgetch()
		}
	}
	scbr()
	return amt
}
