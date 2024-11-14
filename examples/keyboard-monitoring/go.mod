module github.com/launchpad-mini-screen/examples/keyboard-monitoring

go 1.21.0

require (
	github.com/MarinX/keylogger v0.0.0-20210528193429-a54d7834cc1a
	github.com/Dormant512/launchpad-mini-screen v1.0.1
	gitlab.com/gomidi/midi/v2 v2.0.30
)

replace github.com/Dormant512/launchpad-mini-screen => ../..
