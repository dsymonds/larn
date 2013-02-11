package main

/*
	class[c[LEVEL]-1] gives the correct name of the players experience level
*/
const aa1 = " mighty evil master"
const aa2 = "apprentice demi-god"
const aa3 = "  minor demi-god   "
const aa4 = "  major demi-god   "
const aa5 = "    minor deity    "
const aa6 = "    major deity    "
const aa7 = "  novice guardian  "
const aa8 = "apprentice guardian"
const aa9 = "    The Creator    "

var class = [...]string{
	"  novice explorer  ", "apprentice explorer", " practiced explorer", /* -3 */
	"   expert explorer ", "  novice adventurer", "     adventurer    ", /* -6 */
	"apprentice conjurer", "     conjurer      ", "  master conjurer  ", /* -9 */
	"  apprentice mage  ", "        mage       ", "  experienced mage ", /* -12 */
	"     master mage   ", " apprentice warlord", "   novice warlord  ", /* -15 */
	"   expert warlord  ", "   master warlord  ", " apprentice gorgon ", /* -18 */
	"       gorgon      ", "  practiced gorgon ", "   master gorgon   ", /* -21 */
	"    demi-gorgon    ", "    evil master    ", " great evil master ", /* -24 */
	aa1, aa1, aa1, /* -27 */
	aa1, aa1, aa1, /* -30 */
	aa1, aa1, aa1, /* -33 */
	aa1, aa1, aa1, /* -36 */
	aa1, aa1, aa1, /* -39 */
	aa2, aa2, aa2, /* -42 */
	aa2, aa2, aa2, /* -45 */
	aa2, aa2, aa2, /* -48 */
	aa3, aa3, aa3, /* -51 */
	aa3, aa3, aa3, /* -54 */
	aa3, aa3, aa3, /* -57 */
	aa4, aa4, aa4, /* -60 */
	aa4, aa4, aa4, /* -63 */
	aa4, aa4, aa4, /* -66 */
	aa5, aa5, aa5, /* -69 */
	aa5, aa5, aa5, /* -72 */
	aa5, aa5, aa5, /* -75 */
	aa6, aa6, aa6, /* -78 */
	aa6, aa6, aa6, /* -81 */
	aa6, aa6, aa6, /* -84 */
	aa7, aa7, aa7, /* -87 */
	aa8, aa8, aa8, /* -90 */
	aa8, aa8, aa8, /* -93 */
	"  earth guardian   ", "   air guardian    ", "   fire guardian   ", /* -96 */
	"  water guardian   ", "  time guardian    ", " ethereal guardian ", /* -99 */
	aa9, aa9, aa9, /* -102 */
}

/*
	table of experience needed to be a certain level of player
	skill[c[LEVEL]] is the experience required to attain the next level
*/
var skill = [...]int{
	0, 10, 20, 40, 80, 160, 320, 640, 1280, 2560, 5120, /* 1-11 */
	10240, 20480, 40960, 100000, 200000, 400000, 700000, 1 * 1e6, /* 12-19 */
	2 * 1e6, 3 * 1e6, 4 * 1e6, 5 * 1e6, 6 * 1e6, 8 * 1e6, 10 * 1e6, /* 20-26 */
	12 * 1e6, 14 * 1e6, 16 * 1e6, 18 * 1e6, 20 * 1e6, 22 * 1e6, 24 * 1e6, 26 * 1e6, 28 * 1e6, /* 27-35 */
	30 * 1e6, 32 * 1e6, 34 * 1e6, 36 * 1e6, 38 * 1e6, 40 * 1e6, 42 * 1e6, 44 * 1e6, 46 * 1e6, /* 36-44 */
	48 * 1e6, 50 * 1e6, 52 * 1e6, 54 * 1e6, 56 * 1e6, 58 * 1e6, 60 * 1e6, 62 * 1e6, 64 * 1e6, /* 45-53 */
	66 * 1e6, 68 * 1e6, 70 * 1e6, 72 * 1e6, 74 * 1e6, 76 * 1e6, 78 * 1e6, 80 * 1e6, 82 * 1e6, /* 54-62 */
	84 * 1e6, 86 * 1e6, 88 * 1e6, 90 * 1e6, 92 * 1e6, 94 * 1e6, 96 * 1e6, 98 * 1e6, 100 * 1e6, /* 63-71 */
	105 * 1e6, 110 * 1e6, 115 * 1e6, 120 * 1e6, 125 * 1e6, 130 * 1e6, 135 * 1e6, 140 * 1e6, /* 72-79 */
	145 * 1e6, 150 * 1e6, 155 * 1e6, 160 * 1e6, 165 * 1e6, 170 * 1e6, 175 * 1e6, 180 * 1e6, /* 80-87 */
	185 * 1e6, 190 * 1e6, 195 * 1e6, 200 * 1e6, 210 * 1e6, 220 * 1e6, 230 * 1e6, 240 * 1e6, /* 88-95 */
	250 * 1e6, 260 * 1e6, 270 * 1e6, 280 * 1e6, 290 * 1e6, 300 * 1e6, /* 96-101 */
}

var lpbuf, lpnt, lpend *byte /* input/output pointers to the buffers */

var cell = make([]cel, (MAXLEVEL+MAXVLEVEL)*MAXX*MAXY) /* pointer to the dungeon storage	 */

var hitp [MAXX][MAXY]int                /* monster hp on level		 */
var iarg [MAXX][MAXY]int                /* arg for the item array	 */
var item [MAXX][MAXY]int                /* objects in maze if any	 */
var know [MAXX][MAXY]bool               /* whether here before	 */
var mitem [MAXX][MAXY]int               /* monster item array 		 */
var moved [MAXX][MAXY]byte              /* monster movement flags  */
var stealth [MAXX][MAXY]byte            /* 0=sleeping 1=awake monst */
var iven [26]int                        /* inventory for player			 */
var ivenarg [26]int                     /* inventory for player			 */
var lastmonst string                    /* this has the name of the current monster	 */
var beenhere [MAXLEVEL + MAXVLEVEL]bool /* true if have been on this level */

const VERSION = 12 /* this is the present version # of the program	 */
const SUBVERSION = 0

var nosignal byte /* set to 1 to disable the signals from doing anything */

/* 2 means that the trap handling routines
* must do a showplayer() after a trap.  0
* means don't showplayer() 0 - we are in
* create player screen 1 - we are in welcome
* screen 2 - we are in the normal game	 */
var predostuff byte

var loginname [20]int8        /* players login name */
var logname [LOGNAMESIZE]int8 /* players name storage for scoring				 */
var sex byte = 1              /* default is a man  0=woman						 */
var boldon = true             /* 1=bold objects  0=inverse objects				 */
var ckpflag byte = 0          /* 1 if want checkpointing of game, 0 otherwise	 */
var cheat byte = 0            /* 1 if the player has fudged save file			 */
var level = 0                 /* cavelevel player is on = c[CAVELEVEL]			 */
var wizard = false            /* the wizard mode flag							 */
var lastnum = 0               /* the number of the monster last hitting player 	 */
var hitflag int16 = 0         /* flag for if player has been hit when running 	 */
var hit2flag int16 = 0        /* flag for if player has been hit when running 	 */
var hit3flag int16 = 0        /* flag for if player has been hit flush input 	 */
var playerx, playery int      /* the room on the present level of the player		 */

var lastpx, lastpy int /* 0 --- MAXX-1  or  0 --- MAXY-1					 */
var oldx, oldy int
var lasthx, lasthy int = 0, 0 /* location of monster last hit by player		 */

var nobeep int16 = 0            /* true if program is not to beep  					 */
var randx uint32 = 33601        /* the random number seed						 */
var initialtime int32 = 0       /* time playing began 							 */
var gltime int32 = 0            /* the clock for the game						 */
var outstanding_taxes int32 = 0 /* present tax bill from score file 			 */
var c, cbak [100]int            /* the character description arrays			 */
var enable_scroll = 0           /* constant for enabled/disabled scrolling regn */

var spheres *sphere /* pointer to linked list for spheres of annihilation */
var levelname = [...]string{" H", " 1", " 2", " 3", " 4", " 5", " 6", " 7", " 8", " 9", "10", "V1", "V2", "V3"}

const objnamelist = " ATOP%^F&^+M=%^$$f*OD#~][[)))(((||||||||{?!BC}o:@.<<<<EVV))([[]]](^ [H*** ^^ S tsTLc............................................"

var monstnamelist = []byte(" BGHJKOScjtAELNQRZabhiCTYdegmvzFWflorXV pqsyUkMwDDPxnDDuD        ...............................................................")

// TODO: should this be a specific size? same above.
var objectname = [...]string{"",
	"a holy altar", "a handsome jewel encrusted throne", "the orb", "a pit",
	"a staircase leading upwards", "an elevator going up", "a bubbling fountain",
	"a great marble statue", "a teleport trap", "the college of Larn",
	"a mirror", "the DND store", "a staircase going down", "an elevator going down",
	"the bank of Larn", "the 5th branch of the Bank of Larn",
	"a dead fountain", "gold", "an open door", "a closed door",
	"a wall", "The Eye of Larn", "plate mail", "chain mail", "leather armor",
	"a sword of slashing", "Bessman's flailing hammer", "a sunsword",
	"a two handed sword", "a spear", "a dagger",
	"ring of extra regeneration", "a ring of regeneration", "a ring of protection",
	"an energy ring", "a ring of dexterity", "a ring of strength",
	"a ring of cleverness", "a ring of increase damage", "a belt of striking",
	"a magic scroll", "a magic potion", "a book", "a chest",
	"an amulet of invisibility", "an orb of dragon slaying",
	"a scarab of negate spirit", "a cube of undead control",
	"device of theft prevention", "a brilliant diamond", "a ruby",
	"an enchanting emerald", "a sparkling sapphire", "the dungeon entrance",
	"a volcanic shaft leaning downward", "the base of a volcanic shaft",
	"a battle axe", "a longsword", "a flail", "ring mail", "studded leather armor",
	"splint mail", "plate armor", "stainless plate armor", "a lance of death",
	"an arrow trap", "an arrow trap", "a shield", "your home",
	"gold", "gold", "gold", "a dart trap",
	"a dart trap", "a trapdoor", "a trapdoor", "the local trading post",
	"a teleport trap", "a massive throne",
	"a sphere of annihilation", "a handsome jewel encrusted throne",
	"the Larn Revenue Service", "a fortune cookie", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
}

/*
 *	for the monster data
 *
 *	array to do rnd() to create monsters <= a given level
 */
var monstlevel = []int{5, 11, 17, 22, 27, 33, 39, 42, 46, 50, 53, 56, 59}

var monster = [...]monst{
	/*
	 * NAME			LV	AC	DAM	ATT	DEF	GEN
	 * INT GOLD	HP	EXP
	 * -----------------------------------------------------------------
	 */
	{"", 0, 0, 0, 0, 0, false, 3, 0, 0, 0},
	{"bat", 1, 0, 1, 0, 0, false, 3, 0, 1, 1},
	{"gnome", 1, 10, 1, 0, 0, false, 8, 30, 2, 2},
	{"hobgoblin", 1, 14, 2, 0, 0, false, 5, 25, 3, 2},
	{"jackal", 1, 17, 1, 0, 0, false, 4, 0, 1, 1},
	{"kobold", 1, 20, 1, 0, 0, false, 7, 10, 1, 1},

	{"orc", 2, 12, 1, 0, 0, false, 9, 40, 4, 2},
	{"snake", 2, 15, 1, 0, 0, false, 3, 0, 3, 1},
	{"giant centipede", 2, 14, 0, 4, 0, false, 3, 0, 1, 2},
	{"jaculi", 2, 20, 1, 0, 0, false, 3, 0, 2, 1},
	{"troglodyte", 2, 10, 2, 0, 0, false, 5, 80, 4, 3},
	{"giant ant", 2, 8, 1, 4, 0, false, 4, 0, 5, 5},

	/*
	 * NAME			LV	AC	DAM	ATT	DEF	GEN
	 * INT GOLD	HP	EXP
	 * -----------------------------------------------------------------
	 */

	{"floating eye", 3, 8, 1, 0, 0, false, 3, 0, 5, 2},
	{"leprechaun", 3, 3, 0, 8, 0, false, 3, 1500, 13, 45},
	{"nymph", 3, 3, 0, 14, 0, false, 9, 0, 18, 45},
	{"quasit", 3, 5, 3, 0, 0, false, 3, 0, 10, 15},
	{"rust monster", 3, 4, 0, 1, 0, false, 3, 0, 18, 25},
	{"zombie", 3, 12, 2, 0, 0, false, 3, 0, 6, 7},

	{"assassin bug", 4, 9, 3, 0, 0, false, 3, 0, 20, 15},
	{"bugbear", 4, 5, 4, 15, 0, false, 5, 40, 20, 35},
	{"hell hound", 4, 5, 2, 2, 0, false, 6, 0, 16, 35},
	{"ice lizard", 4, 11, 2, 10, 0, false, 6, 50, 16, 25},
	{"centaur", 4, 6, 4, 0, 0, false, 10, 40, 24, 45},

	/*
	 * NAME			LV	AC	DAM	ATT	DEF	GEN
	 * INT GOLD	HP	EXP
	 * -----------------------------------------------------------------
	 */

	{"troll", 5, 4, 5, 0, 0, false, 9, 80, 50, 300},
	{"yeti", 5, 6, 4, 0, 0, false, 5, 50, 35, 100},
	{"white dragon", 5, 2, 4, 5, 0, false, 16, 500, 55, 1000},
	{"elf", 5, 8, 1, 0, 0, false, 15, 50, 22, 35},
	{"gelatinous cube", 5, 9, 1, 0, 0, false, 3, 0, 22, 45},

	{"metamorph", 6, 7, 3, 0, 0, false, 3, 0, 30, 40},
	{"vortex", 6, 4, 3, 0, 0, false, 3, 0, 30, 55},
	{"ziller", 6, 15, 3, 0, 0, false, 3, 0, 30, 35},
	{"violet fungi", 6, 12, 3, 0, 0, false, 3, 0, 38, 100},
	{"wraith", 6, 3, 1, 6, 0, false, 3, 0, 30, 325},
	{"forvalaka", 6, 2, 5, 0, 0, false, 7, 0, 50, 280},

	/*
	 * NAME			LV	AC	DAM	ATT	DEF	GEN
	 * INT GOLD	HP	EXP
	 * -----------------------------------------------------------------
	 */

	{"lama nobe", 7, 7, 3, 0, 0, false, 6, 0, 35, 80},
	{"osequip", 7, 4, 3, 16, 0, false, 4, 0, 35, 100},
	{"rothe", 7, 15, 5, 0, 0, false, 3, 100, 50, 250},
	{"xorn", 7, 0, 6, 0, 0, false, 13, 0, 60, 300},
	{"vampire", 7, 3, 4, 6, 0, false, 17, 0, 50, 1000},
	{"invisible stalker", 7, 3, 6, 0, 0, false, 5, 0, 50, 350},

	{"poltergeist", 8, 1, 4, 0, 0, false, 3, 0, 50, 450},
	{"disenchantress", 8, 3, 0, 9, 0, false, 3, 0, 50, 500},
	{"shambling mound", 8, 2, 5, 0, 0, false, 6, 0, 45, 400},
	{"yellow mold", 8, 12, 4, 0, 0, false, 3, 0, 35, 250},
	{"umber hulk", 8, 3, 7, 11, 0, false, 14, 0, 65, 600},

	/*
	 * NAME			LV	AC	DAM	ATT	DEF	GEN
	 * INT GOLD	HP	EXP
	 * -----------------------------------------------------------------
	 */

	{"gnome king", 9, -1, 10, 0, 0, false, 18, 2000, 100, 3000},
	{"mimic", 9, 5, 6, 0, 0, false, 8, 0, 55, 99},
	{"water lord", 9, -10, 15, 7, 0, false, 20, 0, 150, 15000},
	{"bronze dragon", 9, 2, 9, 3, 0, false, 16, 300, 80, 4000},
	{"green dragon", 9, 3, 8, 10, 0, false, 15, 200, 70, 2500},
	{"purple worm", 9, -1, 11, 0, 0, false, 3, 100, 120, 15000},
	{"xvart", 9, -2, 12, 0, 0, false, 13, 0, 90, 1000},

	{"spirit naga", 10, -20, 12, 12, 0, false, 23, 0, 95, 20000},
	{"silver dragon", 10, -1, 12, 3, 0, false, 20, 700, 100, 10000},
	{"platinum dragon", 10, -5, 15, 13, 0, false, 22, 1000, 130, 24000},
	{"green urchin", 10, -3, 12, 0, 0, false, 3, 0, 85, 5000},
	{"red dragon", 10, -2, 13, 3, 0, false, 19, 800, 110, 14000},

	{"type I demon lord", 12, -30, 18, 0, 0, false, 20, 0, 140, 50000},
	{"type II demon lord", 13, -30, 18, 0, 0, false, 21, 0, 160, 75000},
	{"type III demon lord", 14, -30, 18, 0, 0, false, 22, 0, 180, 100000},
	{"type IV demon lord", 15, -35, 20, 0, 0, false, 23, 0, 200, 125000},
	{"type V demon lord", 16, -40, 22, 0, 0, false, 24, 0, 220, 150000},
	{"type VI demon lord", 17, -45, 24, 0, 0, false, 25, 0, 240, 175000},
	{"type VII demon lord", 18, -70, 27, 6, 0, false, 26, 0, 260, 200000},
	{"demon prince", 25, -127, 30, 6, 0, false, 28, 0, 345, 300000},

	/*
	 * NAME				LV	AC	DAM	ATT	DEF
	 * GEN INT GOLD	HP	EXP
	 * -------------------------------------------------------------------
	 * --
	 */
}

/* name array for scrolls		 */

var scrollname = [...]string{"", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", ""}

var scrollhide = [...]string{
	" enchant armor",
	" enchant weapon",
	" enlightenment",
	" blank paper",
	" create monster",
	" create artifact",
	" aggravate monsters",
	" time warp",
	" teleportation",
	" expanded awareness",
	" haste monsters",
	" monster healing",
	" spirit protection",
	" undead protection",
	" stealth",
	" magic mapping",
	" hold monsters",
	" gem perfection",
	" spell extension",
	" identify",
	" remove curse",
	" annihilation",
	" pulverization",
	" life protection",
	"  ",
	"  ",
	"  ",
	"  ",
}

var potionname = [...]string{"", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", ""}

/* name array for magic potions	 */
var potionhide = [...]string{
	" sleep",
	" healing",
	" raise level",
	" increase ability",
	" wisdom",
	" strength",
	" raise charisma",
	" dizziness",
	" learning",
	" gold detection",
	" monster detection",
	" forgetfulness",
	" water",
	" blindness",
	" confusion",
	" heroism",
	" sturdiness",
	" giant strength",
	" fire resistance",
	" treasure finding",
	" instant healing",
	" cure dianthroritis",
	" poison",
	" see invisible",
	"  ",
	"  ",
	"  ",
	"  ",
	"  ",
	"  ",
	"  ",
	"  ",
	"  ",
	"  ",
	"  ",
}

/*
	spell data
*/
var spelknow [SPNUM]bool
var splev = [...]byte{1, 4, 9, 14, 18, 22, 26, 29, 32, 35, 37, 37, 37, 37, 37}

var spelcode = [...]string{
	"pro", "mle", "dex", "sle", "chm", "ssp",
	"web", "str", "enl", "hel", "cbl", "cre", "pha", "inv",
	"bal", "cld", "ply", "can", "has", "ckl", "vpr",
	"dry", "lit", "drl", "glo", "flo", "fgr",
	"sca", "hld", "stp", "tel", "mfi", /* 31 */
	"sph", "gen", "sum", "wtw", "alt", "per",
}

var spelname = [...]string{
	"protection", "magic missile", "dexterity",
	"sleep", "charm monster", "sonic spear",

	"web", "strength", "enlightenment",
	"healing", "cure blindness", "create monster",
	"phantasmal forces", "invisibility",

	"fireball", "cold", "polymorph",
	"cancellation", "haste self", "cloud kill",
	"vaporize rock",

	"dehydration", "lightning", "drain life",
	"invulnerability", "flood", "finger of death",

	"scare monster", "hold monster", "time stop",
	"teleport away", "magic fire",

	"sphere of annihilation", "genocide", "summon demon",
	"walk through walls", "alter reality", "permanence",
	"",
}

var speldescript = [...]string{
	/* 1 */
	"generates a +2 protection field",
	"creates and hurls a magic missile equivalent to a +1 magic arrow",
	"adds +2 to the caster's dexterity",
	"causes some monsters to go to sleep",
	"some monsters may be awed at your magnificence",
	"causes your hands to emit a screeching sound toward what they point",
	/* 7 */
	"causes strands of sticky thread to entangle an enemy",
	"adds +2 to the caster's strength for a short term",
	"the caster becomes aware of things in the vicinity",
	"restores some hp to the caster",
	"restores sight to one so unfortunate as to be blinded",
	"creates a monster near the caster appropriate for the location",
	"creates illusions, and if believed, monsters die",
	"the caster becomes invisible",
	/* 15 */
	"makes a ball of fire that burns on what it hits",
	"sends forth a cone of cold which freezes what it touches",
	"you can find out what this does for yourself",
	"negates the ability of a monster to use its special abilities",
	"speeds up the caster's movements",
	"creates a fog of poisonous gas which kills all that is within it",
	"this changes rock to air",
	/* 22 */
	"dries up water in the immediate vicinity",
	"your finger will emit a lightning bolt when this spell is cast",
	"subtracts hit points from both you and a monster",
	"this globe helps to protect the player from physical attack",
	"this creates an avalanche of H2O to flood the immediate chamber",
	"this is a holy spell and calls upon your god to back you up",
	/* 28 */
	"terrifies the monster so that hopefully it won't hit the magic user",
	"the monster is frozen in its tracks if this is successful",
	"all movement in the caverns ceases for a limited duration",
	"moves a particular monster around in the dungeon (hopefully away from you)",
	"this causes a curtain of fire to appear all around you",
	/* 33 */
	"anything caught in this sphere is instantly killed.  Warning -- dangerous",
	"eliminates a species of monster from the game -- use sparingly",
	"summons a demon who hopefully helps you out",
	"allows the player to walk through walls for a short period of time",
	"god only knows what this will do",
	"makes a character spell permanent, i. e. protection, strength, etc.",
	"",
}

var spelweird = [MAXMONST + 8][SPNUM]int8{
	/*
	 * p m d s c s    w s e h c c p i    b c p c h c v    d l d g f f
	 * s h s t m    s g s w a p
	 */
	/*
	 * r l e l h s    e t n e b r h n    a l l a a k p    r i r l l g
	 * c l t e f    p e u t l e
	 */
	/*
	 * o e x e m p    b r l l l e a v    l d y n s l r    y t l o o r
	 * a d p l i    h n m w t r
	 */

	/* bat */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* gnome */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* hobgoblin */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* jackal */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* kobold */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

	/* orc */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* snake */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* giant centipede */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* jaculi */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* troglodyte */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

	/* giant ant */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* floating eye */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* leprechaun */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* nymph */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* quasit */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

	/* rust monster */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* zombie */ {0, 0, 0, 8, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0, 4, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* assassin bug */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* bugbear */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* hell hound */ {0, 6, 0, 0, 0, 0, 12, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

	/* ice lizard */ {0, 0, 0, 0, 0, 0, 11, 0, 0, 0, 0, 0, 0, 0, 0, 15, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* centaur */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* troll */ {0, 7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* yeti */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 15, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* white dragon */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 14, 0, 0, 15, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

	/* elf */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 14, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* gelatinous cube */ {0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* metamorph */ {0, 13, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0, 4, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* vortex */ {0, 13, 0, 0, 0, 10, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0, 4, 0, 0, 0, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* ziller */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

	/* violet fungi */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* wraith */ {0, 0, 0, 8, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0, 4, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* forvalaka */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* lama nobe */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* osequip */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

	/* rothe */ {0, 7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* xorn */ {0, 7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* vampire */ {0, 0, 0, 8, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* invisible staker */ {0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* poltergeist */ {0, 13, 0, 8, 0, 4, 1, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 4, 0, 4, 0, 0, 0, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

	/* disenchantress */ {0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* shambling mound */ {0, 0, 0, 0, 0, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* yellow mold */ {0, 0, 0, 8, 0, 0, 1, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* umber hulk */ {0, 7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* gnome king */ {0, 7, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

	/* mimic */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* water lord */ {0, 13, 0, 8, 3, 4, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 9, 0, 0, 4, 0, 0, 0, 0, 0, 16, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* bronze dragon */ {0, 7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* green dragon */ {0, 7, 0, 0, 0, 0, 11, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* purple worm */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

	/* xvart */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* spirit naga */ {0, 13, 0, 8, 3, 4, 1, 0, 0, 0, 0, 0, 0, 5, 0, 4, 9, 0, 0, 4, 0, 4, 0, 0, 0, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* silver dragon */ {0, 6, 0, 9, 0, 0, 12, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* platinum dragon */ {0, 7, 0, 9, 0, 0, 11, 0, 0, 0, 0, 0, 14, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* green urchin */ {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	/* red dragon */ {0, 6, 0, 0, 0, 0, 12, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},

	/*
	 * p m d s c s    w s e h c c p i    b c p c h c v    d l d g f f
	 * s h s t m    s g s w a p
	 */
	/*
	 * r l e l h s    e t n e b r h n    a l l a a k p    r i r l l g
	 * c l t e f    p e u t l e
	 */
	/*
	 * o e x e m p    b r l l l e a v    l d y n s l r    y t l o o r
	 * a d p l i    h n m w t r
	 */

	/* demon lord */ {0, 7, 0, 4, 3, 0, 1, 0, 0, 0, 0, 0, 14, 5, 0, 0, 4, 0, 0, 4, 0, 4, 0, 0, 0, 4, 4, 0, 0, 0, 0, 0, 9, 0, 0, 0, 0, 0},
	/* demon lord */ {0, 7, 0, 4, 3, 0, 1, 0, 0, 0, 0, 0, 14, 5, 0, 0, 4, 0, 0, 4, 0, 4, 0, 0, 0, 4, 4, 0, 0, 0, 0, 0, 9, 0, 0, 0, 0, 0},
	/* demon lord */ {0, 7, 0, 4, 3, 0, 1, 0, 0, 0, 0, 0, 14, 5, 0, 0, 4, 0, 0, 4, 0, 4, 0, 0, 0, 4, 4, 0, 0, 0, 0, 0, 9, 0, 0, 0, 0, 0},
	/* demon lord */ {0, 7, 0, 4, 3, 0, 1, 0, 0, 0, 0, 0, 14, 5, 0, 0, 4, 0, 0, 4, 0, 4, 0, 0, 0, 4, 4, 0, 0, 0, 0, 0, 9, 0, 0, 0, 0, 0},
	/* demon lord */ {0, 7, 0, 4, 3, 0, 1, 0, 0, 0, 0, 0, 14, 5, 0, 0, 4, 0, 0, 4, 0, 4, 0, 0, 0, 4, 4, 0, 0, 0, 0, 0, 9, 0, 0, 0, 0, 0},
	/* demon lord */ {0, 7, 0, 4, 3, 0, 1, 0, 0, 0, 0, 0, 14, 5, 0, 0, 4, 0, 0, 4, 0, 4, 0, 0, 0, 4, 4, 0, 0, 0, 0, 0, 9, 0, 0, 0, 0, 0},
	/* demon lord */ {0, 7, 0, 4, 3, 0, 1, 0, 0, 0, 0, 0, 14, 5, 0, 0, 4, 0, 0, 4, 0, 4, 0, 0, 0, 4, 4, 0, 0, 0, 0, 0, 9, 0, 0, 0, 0, 0},
	/* demon prince */ {0, 7, 0, 4, 3, 9, 1, 0, 0, 0, 0, 0, 14, 5, 0, 0, 4, 0, 0, 4, 0, 4, 0, 0, 0, 4, 4, 4, 0, 0, 0, 4, 9, 0, 0, 0, 0, 0},
}

var spelmes = [...]string{"",
	/* 1 */ "the web had no effect on the %s",
	/* 2 */ "the %s changed shape to avoid the web",
	/* 3 */ "the %s isn't afraid of you",
	/* 4 */ "the %s isn't affected",
	/* 5 */ "the %s can see you with his infravision",
	/* 6 */ "the %s vaporizes your missile",
	/* 7 */ "your missile bounces off the %s",
	/* 8 */ "the %s doesn't sleep",
	/* 9 */ "the %s resists",
	/* 10 */ "the %s can't hear the noise",
	/* 11 */ "the %s's tail cuts it free of the web",
	/* 12 */ "the %s burns through the web",
	/* 13 */ "your missiles pass right through the %s",
	/* 14 */ "the %s sees through your illusions",
	/* 15 */ "the %s loves the cold!",
	/* 16 */ "the %s loves the water!",
}

/*
 *	function to create scroll numbers with appropriate probability of
 *	occurrence
 *
 *	0 - armor			1 - weapon		2 - enlightenment	3 - paper
 *	4 - create monster	5 - create item	6 - aggravate		7 - time warp
 *	8 - teleportation	9 - expanded awareness				10 - haste monst
 *	11 - heal monster	12 - spirit protection		13 - undead protection
 *	14 - stealth		15 - magic mapping			16 - hold monster
 *	17 - gem perfection 18 - spell extension		19 - identify
 *	20 - remove curse	21 - annihilation			22 - pulverization
 *  23 - life protection
 */
var scprob = [...]int{0, 0, 0, 0, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 3, 3,
	3, 3, 3, 4, 4, 4, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 7, 7, 7, 7, 8, 8, 8, 9, 9,
	9, 9, 10, 10, 10, 10, 11, 11, 11, 12, 12, 12, 13, 13, 13, 13, 14, 14,
	15, 15, 16, 16, 16, 17, 17, 18, 18, 19, 19, 19, 20, 20, 20, 20, 21, 22,
	22, 22, 23}

/*
 *	function to return a potion number created with appropriate probability
 *	of occurrence
 *
 *	0 - sleep				1 - healing					2 - raise level
 *	3 - increase ability	4 - gain wisdom				5 - gain strength
 *	6 - charismatic character	7 - dizziness			8 - learning
 *	9 - gold detection		10 - monster detection		11 - forgetfulness
 *	12 - water				13 - blindness				14 - confusion
 *	15 - heroism			16 - sturdiness				17 - giant strength
 *	18 - fire resistance	19 - treasure finding		20 - instant healing
 *	21 - cure dianthroritis	22 - poison					23 - see invisible
 */
var potprob = [...]int{0, 0, 1, 1, 1, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 9, 9, 9, 10, 10, 10, 11, 11, 12, 12, 13, 14, 15, 16, 17, 18, 19, 19, 20, 20, 22, 22, 23, 23}

var nlpts = [...]int{0, 0, 0, 0, 0, 1, 1, 2, 2, 3, 3, 4, 5, 6, 7}
var nch = [...]int{0, 0, 0, 1, 1, 1, 2, 2, 3, 4}
var nplt = [...]int{0, 0, 0, 0, 1, 1, 2, 2, 3, 4}
var ndgg = [...]int{0, 0, 0, 1, 1, 1, 1, 2, 2, 3, 3, 4, 5}
var nsw = [...]int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 2, 3}
