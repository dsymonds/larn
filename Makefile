# EXTRA
#	Incorporates code to gather additional performance statistics
#
# TERMIO
#	Use sysv termio
# TERMIOS
#	Use posix termios
# HIDEBYLINK
#	If defined, the program attempts to hide from ps
# DOCHECKPOINTS
#	If not defined, checkpoint files are periodically written by the
#	larn process (no forking) if enabled in the .larnopts description
#	file.  Checkpointing is handy on an unreliable system, but takes
#	CPU. Inclusion of DOCHECKPOINTS will cause fork()ing to perform the
#	checkpoints (again if enabled in the .larnopts file).  This usually
#	avoids pauses in larn while the checkpointing is being done (on
#	large machines).
# VER
#	This is the version of the software, example:  12
# SUBVER
#	This is the revision of the software, example:  1
# FLUSHNO=#
#	Set the input queue excess flushing threshold (default 5)
# NOVARARGS
#	Define for systems that don't have varargs (a default varargs will
#	be used).
#  VT100
#	Compile for using vt100 family of terminals.  Omission of this
#	define will cause larn to use termcap, but it will be MUCH slower
#	due to an extra layer of output interpretation.  Also, only VT100
#	mode allows 2 different standout modes, inverse video, and bold video.
#	And only in VT100 mode is the scrolling region of the terminal used
#	(much nicer than insert/delete line sequences to simulate it, if
#	VT100 is omitted).
# NOLOG
#	Turn off logging.

.include <bsd.own.mk>

PROG=	larn
MAN=	larn.6
CPPFLAGS+=-DVER=12 -DSUBVER=0 -DTERMIOS
SRCS=	main.c object.c create.c tok.c display.c global.c data.c io.c \
	monster.c store.c diag.c help.c config.c nap.c bill.c scores.c \
	signal.c action.c moreobj.c movem.c regen.c fortune.c savelev.c
DPADD=	${LIBTERMINFO}
LDADD=	-lterminfo
HIDEGAME=hidegame
SETGIDGAME=yes

.if ${MKSHARE} != "no"
DAT=larnmaze larnopts larn.help
FILES=${DAT:S@^@${.CURDIR}/datfiles/@g}
FILESDIR=/usr/share/games/larn
.endif

COPTS.display.c += -Wno-format-nonliteral
COPTS.monster.c += -Wno-format-nonliteral

.include <bsd.prog.mk>
