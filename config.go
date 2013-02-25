package main

import (
	"os"
	"path/filepath"
)

/*
 * All these strings will be appended to in main() to be complete filenames
 */

/* the game save filename */
var savefilename = filepath.Join(os.Getenv("HOME"), ".larn.save")

/* the logging file */
var logfile = "./llog12.5"

/* the score file */
var scorefile = filepath.Join(os.Getenv("HOME"), ".larn-score12.5")

/* the .larnopts filename */
var optsfile = filepath.Join(os.Getenv("HOME"), ".larnopts")

var diagfile = "Diagfile"    /* the diagnostic filename */
var ckpfile = "Larn12.5.ckp" /* the checkpoint filename */
const password = "pvnert(x)" /* the wizards password <=32 */

const WIZID = true // whether to allow wizard operations
