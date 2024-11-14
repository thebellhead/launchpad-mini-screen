package main

import (
	lms "github.com/Dormant512/launchpad-mini-screen"
	"log"
	"time"
)

func main() {
	g := lms.Grid{}
	var i uint8 = 0
	for x := 0; x < len(g); x++ {
		for y := 0; y < len(g[0]); y++ {
			g[x][y] = i
			i++
		}
	}

	lpd, err := lms.NewLaunchpadDriver(lms.LPD_MINI, 0)
	if err != nil {
		log.Fatal(err)
	}

	err = lpd.RenderGrid(g)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(10 * time.Second)
	err = lpd.ClearRender()
	if err != nil {
		log.Fatal(err)
	}
}
