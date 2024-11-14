package main

import (
	"context"
	lms "github.com/Dormant512/launchpad-mini-screen"
	"github.com/MarinX/keylogger"
	_ "gitlab.com/gomidi/midi/v2/drivers/portmididrv" // registers driver
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

const COL = lms.R

func parseFontFile(path string) (map[string]lms.Grid, error) {
	fontMap := make(map[string]lms.Grid)
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	stringSlice := strings.Split(string(bytes), "\n")
	curChar := ""
	curGrid := lms.Grid{}
	idx := 0
	for _, s := range stringSlice {
		if len(s) == 1 {
			curChar = s
		} else if len(s) == 0 {
			fontMap[curChar] = curGrid
			curGrid = lms.Grid{}
			idx = 0
		} else {
			for ch := 0; ch < len(curGrid[0]); ch++ {
				if s[ch] == '8' {
					curGrid[idx][ch] = COL
				}
			}
			idx++
		}
	}
	return fontMap, nil
}

func monitorRoutine(ctx context.Context, wg *sync.WaitGroup, k *keylogger.KeyLogger, lpd *lms.LaunchpadDriver) {
	defer wg.Done()
	fontMap, _ := parseFontFile("./letters")
	events := k.Read()
	for {
		select {
		case <-ctx.Done():
			err := k.Close()
			if err != nil {
				log.Fatal(err)
			}
			err = lpd.ClearRender()
			if err != nil {
				log.Fatal(err)
			}
			return
		case e := <-events:
			if e.Type == keylogger.EvKey && e.KeyPress() {
				err := lpd.RenderGrid(fontMap[e.KeyString()])
				if err != nil {
					log.Fatal(err)
				}
				continue
			}
			if e.Type == keylogger.EvKey && e.KeyRelease() {
				err := lpd.ClearRender()
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup

	lpd, err := lms.NewLaunchpadDriver(lms.LPD_MINI, 0)
	if err != nil {
		log.Fatal(err)
	}
	keyboard := keylogger.FindKeyboardDevice()
	if len(keyboard) == 0 {
		log.Fatal("can't find Keyboard")
		return
	}
	k, err := keylogger.New(keyboard)
	if err != nil {
		log.Fatal(err)
		return
	}

	wg.Add(1)
	go monitorRoutine(ctx, &wg, k, lpd)
	wg.Wait()
}
