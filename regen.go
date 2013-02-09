package main

/*
	*******
	REGEN()
	*******
	regen()

	subroutine to regenerate player hp and spells
*/
func regen() {
	c[MOVESMADE]++
	if c[TIMESTOP] {
		c[TIMESTOP]--
		if c[TIMESTOP] <= 0 {
			bottomline()
		}
		return
	} /* for stop time spell */
	flag := 0

	if c[STRENGTH] < 3 {
		c[STRENGTH] = 3
		flag = 1
	}
	if c[HASTESELF] == 0 || c[HASTESELF]&1 == 0 {
		gltime++
	}

	if c[HP] != c[HPMAX] {
		c[REGENCOUNTER]--
		if c[REGENCOUNTER] < 0 { /* regenerate hit points	 */
			c[REGENCOUNTER] = 22 + (c[HARDGAME] << 1) - c[LEVEL]
			c[HP] += c[REGEN]
			if c[HP] > c[HPMAX] {
				c[HP] = c[HPMAX]
			}
			bottomhp()
		}
	}
	if c[SPELLS] < c[SPELLMAX] { /* regenerate spells	 */
		c[ECOUNTER]--
		if c[ECOUNTER] < 0 {
			c[ECOUNTER] = 100 + 4*(c[HARDGAME]-c[LEVEL]-c[ENERGY])
			c[SPELLS]++
			bottomspell()
		}
	}
	if c[HERO] {
		c[HERO]--
		if c[HERO] <= 0 {
			for i := 0; i < 6; i++ {
				c[i] -= 10
			}
			flag = 1
		}
	}
	if c[ALTPRO] {
		c[ALTPRO]--
		if c[ALTPRO] <= 0 {
			c[MOREDEFENSES] -= 3
			flag = 1
		}
	}
	if c[PROTECTIONTIME] {
		c[PROTECTIONTIME]--
		if c[PROTECTIONTIME] <= 0 {
			c[MOREDEFENSES] -= 2
			flag = 1
		}
	}
	if c[DEXCOUNT] {
		c[DEXCOUNT]--
		if c[DEXCOUNT] <= 0 {
			c[DEXTERITY] -= 3
			flag = 1
		}
	}
	if c[STRCOUNT] {
		c[STRCOUNT]--
		if c[STRCOUNT] <= 0 {
			c[STREXTRA] -= 3
			flag = 1
		}
	}
	if c[BLINDCOUNT] {
		c[BLINDCOUNT]--
		if c[BLINDCOUNT] <= 0 {
			cursors()
			lprcat("\nThe blindness lifts  ")
			beep()
		}
	}
	if c[CONFUSE] {
		c[CONFUSE]--
		if c[CONFUSE] <= 0 {
			cursors()
			lprcat("\nYou regain your senses")
			beep()
		}
	}
	if c[GIANTSTR] {
		c[GIANTSTR]--
		if c[GIANTSTR] <= 0 {
			c[STREXTRA] -= 20
			flag = 1
		}
	}
	if c[CHARMCOUNT] {
		c[CHARMCOUNT]--
		if c[CHARMCOUNT] <= 0 {
			flag = 1
		}
	}
	if c[INVISIBILITY] {
		c[INVISIBILITY]--
		if c[INVISIBILITY] <= 0 {
			flag = 1
		}
	}
	if c[CANCELLATION] {
		c[CANCELLATION]--
		if c[CANCELLATION] <= 0 {
			flag = 1
		}
	}
	if c[WTW] {
		c[WTW]--
		if c[WTW] <= 0 {
			flag = 1
		}
	}
	if c[HASTESELF] {
		c[HASTESELF]--
		if c[HASTESELF] <= 0 {
			flag = 1
		}
	}
	if c[AGGRAVATE] {
		c[AGGRAVATE]--
	}
	if c[SCAREMONST] {
		c[SCAREMONST]--
		if c[SCAREMONST] <= 0 {
			flag = 1
		}
	}
	if c[STEALTH] {
		c[STEALTH]--
		if c[STEALTH] <= 0 {
			flag = 1
		}
	}
	if c[AWARENESS] {
		c[AWARENESS]--
	}
	if c[HOLDMONST] {
		c[HOLDMONST]--
		if c[HOLDMONST] <= 0 {
			flag = 1
		}
	}
	if c[HASTEMONST] {
		c[HASTEMONST]--
	}
	if c[FIRERESISTANCE] {
		c[FIRERESISTANCE]--
		if c[FIRERESISTANCE] <= 0 {
			flag = 1
		}
	}
	if c[GLOBE] {
		c[GLOBE]--
		if c[GLOBE] <= 0 {
			c[MOREDEFENSES] -= 10
			flag = 1
		}
	}
	if c[SPIRITPRO] {
		c[SPIRITPRO]--
		if c[SPIRITPRO] <= 0 {
			flag = 1
		}
	}
	if c[UNDEADPRO] {
		c[UNDEADPRO]--
		if c[UNDEADPRO] <= 0 {
			flag = 1
		}
	}
	if c[HALFDAM] {
		c[HALFDAM]--
		if c[HALFDAM] <= 0 {
			cursors()
			lprcat("\nYou now feel better ")
			beep()
		}
	}
	if c[SEEINVISIBLE] {
		c[SEEINVISIBLE]--
		if c[SEEINVISIBLE] <= 0 {
			monstnamelist[INVISIBLESTALKER] = ' '
			cursors()
			lprcat("\nYou feel your vision return to normal")
			beep()
		}
	}
	if c[ITCHING] {
		if c[ITCHING] > 1 {
			if c[WEAR] != -1 || c[SHIELD] != -1 {
				if rnd(100) < 50 {
					c[WEAR], c[SHIELD] = -1, -1
					cursors()
					lprcat("\nThe hysteria of itching forces you to remove your armor!")
					beep()
					recalc()
					bottomline()
				}
			}
		}
		c[ITCHING]--
		if c[ITCHING] <= 0 {
			cursors()
			lprcat("\nYou now feel the irritation subside!")
			beep()
		}
	}
	if c[CLUMSINESS] {
		if c[WIELD] != -1 {
			if c[CLUMSINESS] > 1 {
				if item[playerx][playery] == 0 { /* only if nothing there */
					if rnd(100) < 33 { /* drop your weapon due to clumsiness */
						drop_object(c[WIELD])
					}
				}
			}
		}
		c[CLUMSINESS]--
		if c[CLUMSINESS] <= 0 {
			cursors()
			lprcat("\nYou now feel less awkward!")
			beep()
		}
	}
	if flag {
		bottomline()
	}
}
