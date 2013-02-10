Larn is a roguelike game written by Noah Morgan in 1986.
This is a Go port of Larn, by David Symonds (dsymonds@golang.org).

Progress
--------
I have transliterated each original C source file.
I have done the minimal changes required to get it to be valid Go source
code as far as gofmt is concerned. Some chunks have been stubbed out
instead of ported (e.g. loading/saving).

Roadmap
-------
Next I will be making sufficient changes to get it to build,
linking against github.com/jabb/gocurse/curses, which is the only
Go ncurses library I could get to build and work.

After that, hopefully enough will work to be able to play; I shall then
start turning the code into good Go code, and filling out the previously
stubbed functions.

Next, I will abstract enough of the code to turn it into a web based
version, probably hosted on App Engine; you can be playing on one
web browser, sign in on another, and resume the same game. There will
be some kind of scoreboard too.

License
-------
The license of the original Larn by Noah Morgan is unclear. He posted
it to a public discussion group without stating the license; 1986 was
an era before licenses were "a thing". Given that, and given that I plan
to replace all the original C code, I am licensing this under the
[BSD 3-Clause Licence](http://www.opensource.org/licenses/bsd-license.php).
