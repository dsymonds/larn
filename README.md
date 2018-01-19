Larn is a roguelike game written by Noah Morgan in 1986.
This is a Go port of Larn, by David Symonds (dsymonds@golang.org).
There is a web page at http://www.symonds.id.au/larn.

Status
------
I have transliterated each original C source file.
Some chunks have been stubbed out instead of ported (e.g. loading/saving),
but a substantial enough part works that you can play it.
I used Larn 12.0 as the basis, but have ported many of the fixes from Larn 12.3.
I am calling this "Larn 12.5".

Roadmap
-------
I will next be debugging, fixing and filling out the stubbed functions.

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
[BSD 3-Clause Licence](https://opensource.org/licenses/BSD-3-Clause).
