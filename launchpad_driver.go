package launchpad_mini_screen

import (
	"fmt"
	"gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/drivers"
	_ "gitlab.com/gomidi/midi/v2/drivers/portmididrv" // autoregisters driver
	"sync"
)

type LaunchpadDriver struct {
	Out     drivers.Out
	Send    func(msg midi.Message) error
	Channel uint8
	State   Grid
	Mu      *sync.Mutex
	RWMu    *sync.RWMutex
}

func NewLaunchpadDriver(deviceName string, channel uint8) (*LaunchpadDriver, error) {
	if deviceName == "" {
		deviceName = LPD_MINI
	}
	out, err := midi.FindOutPort(deviceName)
	if err != nil {
		return nil, fmt.Errorf("MIDI device %s not found", deviceName)
	}
	send, _ := midi.SendTo(out)
	ld := LaunchpadDriver{
		Out:     out,
		Send:    send,
		Channel: channel,
		State:   *new(Grid),
		Mu:      &sync.Mutex{},
		RWMu:    &sync.RWMutex{},
	}
	return &ld, nil
}

func (lpd *LaunchpadDriver) RenderGrid(g Grid) error {
	for x, row := range g {
		for y, cell := range row {
			note := CoordsToNote(x, y)
			if cell == 0 {
				err := lpd.Send(midi.NoteOff(lpd.Channel, note))
				if err != nil {
					return err
				}
			} else {
				err := lpd.Send(midi.NoteOn(lpd.Channel, note, cell))
				if err != nil {
					return err
				}
			}
		}
	}
	lpd.Mu.Lock()
	lpd.State = g
	lpd.Mu.Unlock()
	return nil
}

func (lpd *LaunchpadDriver) ClearRender() error {
	return lpd.RenderGrid(EMPTY_GRID)
}
