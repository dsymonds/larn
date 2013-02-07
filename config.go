package main

/*
 * All these strings will be appended to in main() to be complete filenames
 */

/* the game save filename */
var savefilename string

/* the logging file */
var logfile = "/var/games/larn/llog12.0"

/* the help text file */
var helpfile = "/usr/share/games/larn/larn.help"

/* the score file */
var scorefile = "/var/games/larn/lscore12.0"

/* the maze data file */
var larnlevels = "/usr/share/games/larn/larnmaze"

/* the .larnopts filename */
var optsfile = "/.larnopts"

/* the player id datafile name */
var playerids = "/var/games/larn/playerids"

var diagfile = "Diagfile"    /* the diagnostic filename */
var ckpfile = "Larn12.0.ckp" /* the checkpoint filename */
const password = "pvnert(x)" /* the wizards password <=32 */
var psname = "larn"          /* the process name */

const WIZID = 1

var wisid = 0 /* the user id of the only person who can be wizard */
