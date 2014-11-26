package hex

import (
)

// A Playout object keeps around all the details of a single playout.
// This is kind of like an options object; there are just a lot of
// different ways to handle a playout.

type XPlayout struct {
	player1 QuickPlayer
	player2 QuickPlayer
	debug bool
}

func NewPlayout(p1 QuickPlayer, p2 QuickPlayer, debug bool) *XPlayout {
	panic("TODO")
}
