package main

const MAXLEVEL = 11 /* max # levels in the dungeon			 */
const MAXVLEVEL = 3 /* max # of levels in the temple of the luran	 */
const MAXX = 67
const MAXY = 17

const SCORESIZE = 10    /* this is the number of people on a scoreboard max */
const MAXPLEVEL = 100   /* maximum player level allowed		 */
const MAXMONST = 56     /* maximum # monsters in the dungeon	 */
const SPNUM = 38        /* maximum number of spells in existence	 */
const MAXSCROLL = 28    /* maximum number of scrolls that are possible	 */
const MAXPOTION = 35    /* maximum number of potions that are possible	 */
const TIMELIMIT = 30000 /* the maximum number of moves before the game is called */
const TAXRATE = 1 / 20  /* the tax rate for the LRS */
const MAXOBJ = 93       /* the maximum number of objects   n < MAXOBJ */

/* this is the structure definition of the monster data	 */
type monst struct {
	name         string
	level        int
	armorclass   int
	damage       int
	attack       int
	defense      int8
	genocided    bool
	intelligence int /* monsters intelligence -- used to choose movement */
	gold         int
	hitpoints    int
	experience   int
}

/* this is the structure definition for the items in the dnd store */
type _itm struct {
	price int16
	obj   byte
	arg   byte
	qty   int8
}

/* this is the structure that holds the entire dungeon specifications	 */
type cel struct {
	hitp  int16 /* monster's hit points	 */
	mitem int8  /* the monster ID			 */
	item  int8  /* the object's ID			 */
	iarg  int16 /* the object's argument	 */
	know  int8  /* have we been here before */
}

/* this is the structure for maintaining & moving the spheres of annihilation */
type sphere struct {
	p         *sphere /* pointer to next structure */
	x, y, lev int8    /* location of the sphere */
	dir       int8    /* direction sphere is going in */
	lifetime  int16   /* duration of the sphere */
}

/* defines for the character attribute array	c[]	 */
const (
	STRENGTH       = 0 /* characters physical strength not due to objects */
	INTELLIGENCE   = 1
	WISDOM         = 2
	CONSTITUTION   = 3
	DEXTERITY      = 4
	CHARISMA       = 5
	HPMAX          = 6
	HP             = 7
	GOLD           = 8
	EXPERIENCE     = 9
	LEVEL          = 10
	REGEN          = 11
	WCLASS         = 12
	AC             = 13
	BANKACCOUNT    = 14
	SPELLMAX       = 15
	SPELLS         = 16
	ENERGY         = 17
	ECOUNTER       = 18
	MOREDEFENSES   = 19
	WEAR           = 20
	PROTECTIONTIME = 21
	WIELD          = 22
	AMULET         = 23
	REGENCOUNTER   = 24
	MOREDAM        = 25
	DEXCOUNT       = 26
	STRCOUNT       = 27
	BLINDCOUNT     = 28
	CAVELEVEL      = 29
	CONFUSE        = 30
	ALTPRO         = 31
	HERO           = 32
	CHARMCOUNT     = 33
	INVISIBILITY   = 34
	CANCELLATION   = 35
	HASTESELF      = 36
	EYEOFLARN      = 37
	AGGRAVATE      = 38
	GLOBE          = 39
	TELEFLAG       = 40
	SLAYING        = 41
	NEGATESPIRIT   = 42
	SCAREMONST     = 43
	AWARENESS      = 44
	HOLDMONST      = 45
	TIMESTOP       = 46
	HASTEMONST     = 47
	CUBEofUNDEAD   = 48
	GIANTSTR       = 49
	FIRERESISTANCE = 50
	BESSMANN       = 51
	NOTHEFT        = 52
	HARDGAME       = 53
	CPUTIME        = 54
	BYTESIN        = 55
	BYTESOUT       = 56
	MOVESMADE      = 57
	MONSTKILLED    = 58
	SPELLSCAST     = 59
	LANCEDEATH     = 60
	SPIRITPRO      = 61
	UNDEADPRO      = 62
	SHIELD         = 63
	STEALTH        = 64
	ITCHING        = 65
	LAUGHING       = 66
	DRAINSTRENGTH  = 67
	CLUMSINESS     = 68
	INFEEBLEMENT   = 69
	HALFDAM        = 70
	SEEINVISIBLE   = 71
	FILLROOM       = 72
	RANDOMWALK     = 73
	SPHCAST        = 74 /* nz if an active sphere of annihilation */
	WTW            = 75 /* walk through walls */
	STREXTRA       = 76 /* character strength due to objects or enchantments */
	TMP            = 77 /* misc scratch space */
	LIFEPROT       = 78 /* life protection counter */
)

/* defines for the objects in the game		 */

const (
	OALTAR        = 1
	OTHRONE       = 2
	OORB          = 3
	OPIT          = 4
	OSTAIRSUP     = 5
	OELEVATORUP   = 6
	OFOUNTAIN     = 7
	OSTATUE       = 8
	OTELEPORTER   = 9
	OSCHOOL       = 10
	OMIRROR       = 11
	ODNDSTORE     = 12
	OSTAIRSDOWN   = 13
	OELEVATORDOWN = 14
	OBANK2        = 15
	OBANK         = 16
	ODEADFOUNTAIN = 17
	OMAXGOLD      = 70
	OGOLDPILE     = 18
	OOPENDOOR     = 19
	OCLOSEDDOOR   = 20
	OWALL         = 21
	OTRAPARROW    = 66
	OTRAPARROWIV  = 67

	OLARNEYE = 22

	OPLATE       = 23
	OCHAIN       = 24
	OLEATHER     = 25
	ORING        = 60
	OSTUDLEATHER = 61
	OSPLINT      = 62
	OPLATEARMOR  = 63
	OSSPLATE     = 64
	OSHIELD      = 68
	OELVENCHAIN  = 92

	OSWORDofSLASHING = 26
	OHAMMER          = 27
	OSWORD           = 28
	O2SWORD          = 29
	OSPEAR           = 30
	ODAGGER          = 31
	OBATTLEAXE       = 57
	OLONGSWORD       = 58
	OFLAIL           = 59
	OLANCE           = 65
	OVORPAL          = 90
	OSLAYER          = 91

	ORINGOFEXTRA = 32
	OREGENRING   = 33
	OPROTRING    = 34
	OENERGYRING  = 35
	ODEXRING     = 36
	OSTRRING     = 37
	OCLEVERRING  = 38
	ODAMRING     = 39

	OBELT = 40

	OSCROLL = 41
	OPOTION = 42
	OBOOK   = 43
	OCHEST  = 44
	OAMULET = 45

	OORBOFDRAGON  = 46
	OSPIRITSCARAB = 47
	OCUBEofUNDEAD = 48
	ONOTHEFT      = 49

	ODIAMOND  = 50
	ORUBY     = 51
	OEMERALD  = 52
	OSAPPHIRE = 53

	OENTRANCE = 54
	OVOLDOWN  = 55
	OVOLUP    = 56
	OHOME     = 69

	OKGOLD        = 71
	ODGOLD        = 72
	OIVDARTRAP    = 73
	ODARTRAP      = 74
	OTRAPDOOR     = 75
	OIVTRAPDOOR   = 76
	OTRADEPOST    = 77
	OIVTELETRAP   = 78
	ODEADTHRONE   = 79
	OANNIHILATION = 80 /* sphere of annihilation */
	OTHRONE2      = 81
	OLRS          = 82 /* Larn Revenue Service */
	OCOOKIE       = 83
	OURN          = 84
	OBRASSLAMP    = 85
	OHANDofFEAR   = 86 /* hand of fear */
	OSPHTAILSMAN  = 87 /* tailsman of the sphere */
	OWWAND        = 88 /* wand of wonder */
	OPSTAFF       = 89 /* staff of power */
	/* used up to 92 */
)

/* defines for the monsters as objects		 */

const (
	BAT              = 1
	GNOME            = 2
	HOBGOBLIN        = 3
	JACKAL           = 4
	KOBOLD           = 5
	ORC              = 6
	SNAKE            = 7
	CENTIPEDE        = 8
	JACULI           = 9
	TROGLODYTE       = 10
	ANT              = 11
	EYE              = 12
	LEPRECHAUN       = 13
	NYMPH            = 14
	QUASIT           = 15
	RUSTMONSTER      = 16
	ZOMBIE           = 17
	ASSASSINBUG      = 18
	BUGBEAR          = 19
	HELLHOUND        = 20
	ICELIZARD        = 21
	CENTAUR          = 22
	TROLL            = 23
	YETI             = 24
	WHITEDRAGON      = 25
	ELF              = 26
	CUBE             = 27
	METAMORPH        = 28
	VORTEX           = 29
	ZILLER           = 30
	VIOLETFUNGI      = 31
	WRAITH           = 32
	FORVALAKA        = 33
	LAMANOBE         = 34
	OSEQUIP          = 35
	ROTHE            = 36
	XORN             = 37
	VAMPIRE          = 38
	INVISIBLESTALKER = 39
	POLTERGEIST      = 40
	DISENCHANTRESS   = 41
	SHAMBLINGMOUND   = 42
	YELLOWMOLD       = 43
	UMBERHULK        = 44
	GNOMEKING        = 45
	MIMIC            = 46
	WATERLORD        = 47
	BRONZEDRAGON     = 48
	GREENDRAGON      = 49
	PURPLEWORM       = 50
	XVART            = 51
	SPIRITNAGA       = 52
	SILVERDRAGON     = 53
	PLATINUMDRAGON   = 54
	GREENURCHIN      = 55
	REDDRAGON        = 56
	DEMONLORD        = 57
	DEMONPRINCE      = 64
)

const (
	BUFBIG      = 4096 /* size of the output buffer */
	MAXIBUF     = 4096 /* size of the input buffer */
	LOGNAMESIZE = 40   /* max size of the players name */
	PSNAMESIZE  = 40   /* max size of the process name */
)

/*
extern char     VERSION, SUBVERSION;
extern u_char   ckpflag;
extern const char *class[];
extern u_char   course[];
extern char     diagfile[], helpfile[], ckpfile[], larnlevels[],
		playerids[], optsfile[1024], psname[], savefilename[],
		scorefile[];
extern u_char   item[MAXX][MAXY], iven[], know[MAXX][MAXY];
extern const char *levelname[];
extern char     logfile[], lastmonst[];
extern u_char  *lpbuf, *lpend;
extern u_char  *lpnt, moved[MAXX][MAXY], mitem[MAXX][MAXY], monstlevel[];
extern char     monstnamelist[], objnamelist[];
extern u_char   nosignal;
extern const char *objectname[];
extern const char *potionhide[];
extern const char *spelcode[], *spelname[], *spelmes[];
extern char     spelweird[MAXMONST + 8][SPNUM];
extern u_char   predostuff, restorflag;
extern u_char   screen[MAXX][MAXY], sex;
extern const char *speldescript[];
extern const char *scrollhide[];
extern u_char   splev[], stealth[MAXX][MAXY];
extern short    diroffx[], diroffy[], hitp[MAXX][MAXY];
extern short    iarg[MAXX][MAXY], ivenarg[], lasthx, lasthy, lastnum, lastpx,
                lastpy;
extern short    nobeep, oldx, oldy, playerx, playery, level;
extern int      enable_scroll, srcount, userid, wisid;
extern gid_t    gid, egid;
extern long     outstanding_taxes, skill[], gltime, c[], cbak[];
extern time_t	initialtime;
extern unsigned long randx;
*/

/*
extern struct monst monster[];
extern struct sphere *spheres;
extern struct _itm itm[];
extern int      rmst, lasttime;
*/

func hardcond(a, b int) int {
	if c[HARDGAME] != 0 {
		return a
	}
	return b
}

/* macro to create scroll #'s with probability of occurrence */
func newscroll() int { return scprob[rund(81)] }

/* macro to return a potion # created with probability of occurrence */
func newpotion() int { return potprob[rund(41)] }

/* macro to return the + points on created leather armor */
func newleather() int { return nlpts[rund(hardcond(13, 15))] }

/* macro to return the + points on chain armor */
func newchain() int { return nch[rund(10)] }

/* macro to return + points on plate armor */
func newplate() int { return nplt[rund(hardcond(4, 12))] }

/* macro to return + points on new daggers */
func newdagger() int { return ndgg[rund(13)] }

/* macro to return + points on new swords */
func newsword() int { return nsw[rund(hardcond(6, 13))] }

/* macro to destroy object at present location */
func forget() { item[playerx][playery], know[playerx][playery] = 0, false }

/* macro to wipe out a monster at a location */
func disappear(x, y int) { mitem[x][y], know[x][y] = 0, false }

// TODO: fix term codes?

/* macro to turn on bold display for the terminal */
func setbold() {
	if boldon {
		lprcat("\033[1m")
	} else {
		lprcat("\033[7m")
	}
}

/* macro to turn off bold display for the terminal */
func resetbold() { lprcat("\033[m") }

/* macro to setup the scrolling region for the terminal */
func setscroll() { lprcat("\033[20;24r") }

/* macro to clear the scrolling region for the terminal */
func resetscroll() { lprcat("\033[;24r") }

/* macro to clear the screen and home the cursor */
func clear()    { lprcat("\033[2J\033[f"); cbak[SPELLS] = -50 }
func cltoeoln() { lprcat("\033[K") }

/* macro to output one byte to the output buffer */
// TODO: move this to io.go.
func lprc(ch byte) { lpbuf = append(lpbuf, ch) }

/* macro to seed the random number generator */
func seedrand(x uint32) { randx = x }

/* macros for miscellaneous data conversion */
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
