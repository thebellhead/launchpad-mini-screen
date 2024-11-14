package launchpad_mini_screen

func CoordsToNote(x, y int) uint8 {
	return uint8(x*16 + y)
}
