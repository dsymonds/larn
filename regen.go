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
	if c[TIMESTOP] != 0 {
		c[TIMESTOP]--
		if c[TIMESTOP] <= 0 {
			bottomline()
		}
		return
	} /* for stop time spell */
	flag := false

	if c[STRENGTH] < 3 {
		c[STRENGTH] = 3
		flag = true
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
	if c[HERO] != 0 {
		c[HERO]--
		if c[HERO] <= 0 {
			for i := 0; i < 6; i++ {
				c[i] -= 10
			}
			flag = true
		}
	}
	if c[ALTPRO] != 0 {
		c[ALTPRO]--
		if c[ALTPRO] <= 0 {
			c[MOREDEFENSES] -= 3
			flag = true
		}
	}
	if c[PROTECTIONTIME] != 0 {
		c[PROTECTIONTIME]--
		if c[PROTECTIONTIME] <= 0 {
			c[MOREDEFENSES] -= 2
			flag = true
		}
	}
	if c[DEXCOUNT] != 0 {
		c[DEXCOUNT]--
		if c[DEXCOUNT] <= 0 {
			c[DEXTERITY] -= 3
			flag = true
		}
	}
	if c[STRCOUNT] != 0 {
		c[STRCOUNT]--
		if c[STRCOUNT] <= 0 {
			c[STREXTRA] -= 3
			flag = true
		}
	}
	if c[BLINDCOUNT] != 0 {
		c[BLINDCOUNT]--
		if c[BLINDCOUNT] <= 0 {
			cursors()
			lprcat("\nThe blindness lifts  ")
			beep()
		}
	}
	if c[CONFUSE] != 0 {
		c[CONFUSE]--
		if c[CONFUSE] <= 0 {
			cursors()
			lprcat("\nYou regain your senses")
			beep()
		}
	}
	if c[GIANTSTR] != 0 {
		c[GIANTSTR]--
		if c[GIANTSTR] <= 0 {
			c[STREXTRA] -= 20
			flag = true
		}
	}
	if c[CHARMCOUNT] != 0 {
		c[CHARMCOUNT]--
		if c[CHARMCOUNT] <= 0 {
			flag = true
		}
	}
	if c[INVISIBILITY] != 0 {
		c[INVISIBILITY]--
		if c[INVISIBILITY] <= 0 {
			flag = true
		}
	}
	if c[CANCELLATION] != 0 {
		c[CANCELLATION]--
		if c[CANCELLATION] <= 0 {
			flag = true
		}
	}
	if c[WTW] != 0 {
		c[WTW]--
		if c[WTW] <= 0 {
			flag = true
		}
	}
	if c[HASTESELF] != 0 {
		c[HASTESELF]--
		if c[HASTESELF] <= 0 {
			flag = true
		}
	}
	if c[AGGRAVATE] != 0 {
		c[AGGRAVATE]--
	}
	if c[SCAREMONST] != 0 {
		c[SCAREMONST]--
		if c[SCAREMONST] <= 0 {
			flag = true
		}
	}
	if c[STEALTH] != 0 {
		c[STEALTH]--
		if c[STEALTH] <= 0 {
			flag = true
		}
	}
	if c[AWARENESS] != 0 {
		c[AWARENESS]--
	}
	if c[HOLDMONST] != 0 {
		c[HOLDMONST]--
		if c[HOLDMONST] <= 0 {
			flag = true
		}
	}
	if c[HASTEMONST] != 0 {
		c[HASTEMONST]--
	}
	if c[FIRERESISTANCE] != 0 {
		c[FIRERESISTANCE]--
		if c[FIRERESISTANCE] <= 0 {
			flag = true
		}
	}
	if c[GLOBE] != 0 {
		c[GLOBE]--
		if c[GLOBE] <= 0 {
			c[MOREDEFENSES] -= 10
			flag = true
		}
	}
	if c[SPIRITPRO] != 0 {
		c[SPIRITPRO]--
		if c[SPIRITPRO] <= 0 {
			flag = true
		}
	}
	if c[UNDEADPRO] != 0 {
		c[UNDEADPRO]--
		if c[UNDEADPRO] <= 0 {
			flag = true
		}
	}
	if c[HALFDAM] != 0 {
		c[HALFDAM]--
		if c[HALFDAM] <= 0 {
			cursors()
			lprcat("\nYou now feel better ")
			beep()
		}
	}
	if c[SEEINVISIBLE] != 0 {
		c[SEEINVISIBLE]--
		if c[SEEINVISIBLE] <= 0 {
			monstnamelist[INVISIBLESTALKER] = ' '
			cursors()
			lprcat("\nYou feel your vision return to normal")
			beep()
		}
	}
	if c[ITCHING] != 0 {
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
	if c[CLUMSINESS] != 0 {
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
