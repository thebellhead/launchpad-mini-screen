package launchpad_mini_screen

type Grid [8][8]uint8

const (
	LPD_MINI = "Launchpad Mini MIDI 1"

	R     uint8 = 3
	R_100 uint8 = 3
	R_66  uint8 = 2
	R_33  uint8 = 1

	G     uint8 = 56
	G_100 uint8 = 56
	G_66  uint8 = 40
	G_33  uint8 = 24

	YG     uint8 = 57
	YG_100 uint8 = 57
	YG_66  uint8 = 41
	YG_33  uint8 = 25

	O  uint8 = 19
	OY uint8 = 35
	Y  uint8 = 59
)

var EMPTY_GRID Grid
