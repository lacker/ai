package hex

import (
)

// A QuickGame object keeps around all the details of a single playout.
// This is kind of like an options object; there are just a lot of
// different ways to handle a playout.

type QuickGame struct {
	player1 QuickPlayer
	player2 QuickPlayer
	debug bool

	// An optional override to control what the players do.
	SnipList []Snip

	// An optional registry to notify when moves are made.
	Registry *SpotRegistry
}

func NewQuickGame(p1 QuickPlayer, p2 QuickPlayer, debug bool) *QuickGame {
	return &QuickGame{
		player1: p1,
		player2: p2,
		debug: debug,
	}
}

// Plays out the game and returns the final board state.
// You are only supposed to call Playout once per QuickGame.
func (game *QuickGame) Playout() *TopoBoard {
	panic("TODO")
}
