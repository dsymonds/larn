package main

/*
 *	routine to save the present level into storage
 */
func savelevel() {
	for x := 0; x < MAXX; x++ {
		for y := 0; y < MAXY; y++ {
			cell[level][x][y] = cel{
				mitem: mitem[x][y],
				hitp:  hitp[x][y],
				item:  item[x][y],
				know:  know[x][y],
				iarg:  iarg[x][y],
			}
		}
	}
}

/*
 *	routine to restore a level from storage
 */
func getlevel() {
	for x := 0; x < MAXX; x++ {
		for y := 0; y < MAXY; y++ {
			c := cell[level][x][y]
			mitem[x][y] = c.mitem
			hitp[x][y] = c.hitp
			item[x][y] = c.item
			know[x][y] = c.know
			iarg[x][y] = c.iarg
		}
	}
}
