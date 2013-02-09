# EXTRA
#	Incorporates code to gather additional performance statistics
#
# TERMIO
#	Use sysv termio
# TERMIOS
#	Use posix termios
# HIDEBYLINK
#	If defined, the program attempts to hide from ps
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
CPPFLAGS+=-DTERMIOS
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
