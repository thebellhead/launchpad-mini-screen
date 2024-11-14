package main

import (
	"fmt"
	lms "github.com/Dormant512/launchpad-mini-screen"
	"log"
	"strings"
)

type checkerboard lms.Grid

func (c checkerboard) checkMove(move string) bool {
	if len(move) != 5 {
		return false
	}

	// check validity of move
	return true
}

func main() {
	board := lms.Grid{}
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[0]); y++ {
			if (x+y)%2 == 0 {
				switch {
				case y < 3:
					board[x][y] = lms.R_33
				case y >= 5:
					board[x][y] = lms.YG_33
				default:
					board[x][y] = lms.G_33
				}
			}
		}
	}

	lpd, err := lms.NewLaunchpadDriver(lms.LPD_MINI, 0)
	if err != nil {
		log.Fatal(err)
	}

	err = lpd.RenderGrid(board)
	if err != nil {
		log.Fatal(err)
	}

	msg := ""
	plrIsRed := true
	plr := ""

	for {
		if plrIsRed {
			plr = "red"
		} else {
			plr = "yellow"
		}
		fmt.Printf("Move (%s): ", plr)
		_, err = fmt.Scanln(&msg)
		if err != nil {
			log.Fatal(err)
		}

		plrIsRed = !plrIsRed

		if strings.ToLower(msg) == "exit" {
			fmt.Println("Exiting...")
			return
		}
	}
}
