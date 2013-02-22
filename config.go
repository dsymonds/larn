package main

/*
 * All these strings will be appended to in main() to be complete filenames
 */

/* the game save filename */
var savefilename string

/* the logging file */
var logfile = "./llog12.0"

/* the help text file */
var helpfile = "datfiles/larn.help"

/* the score file */
var scorefile = "./lscore12.0"

/* the maze data file */
var larnlevels = "datfiles/larnmaze"

/* the .larnopts filename */
var optsfile = "/.larnopts"

var diagfile = "Diagfile"    /* the diagnostic filename */
var ckpfile = "Larn12.0.ckp" /* the checkpoint filename */
const password = "pvnert(x)" /* the wizards password <=32 */
var psname = "larn"          /* the process name */

const WIZID = true // whether to allow wizard operations
