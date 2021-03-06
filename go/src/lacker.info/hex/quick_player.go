package hex

import (
)

// QuickPlayer is an interface for a hex player that is designed to
// run playouts many times per move of a real game.

type QuickPlayer interface {
	// Resets to the starting board position for a new game
	Reset(game *QuickGame)

	// Gets what the starting position is
	// Typically all quick players instantiated at the same point should
	// actually be pointing to the same starting position, not just
	// positions that are the same.
	StartingPosition() *TopoBoard

	// Returns the best move and a score for it.
	// If this player has no idea it can return NotASpot.
	// Positions should be progressing through the game until Reset is
	// called, so a quick player can keep some state around.
	BestMove(board *TopoBoard, debug bool) (TopoSpot, float64)

	// Prints some debug information
	Debug()

	// Gets what color the player is
	Color() Color
}

// EvolvingPlayer is an interface for a QuickPlayer that the
// MetaFarmer can train two opposing players on.

type EvolvingPlayer interface {
	QuickPlayer

	EvolveToPlay(snipList []Snip, ending *TopoBoard, debug bool)

	// Finds a game that evolves from this one
	FindNewMainLine(opponent EvolvingPlayer, oldMainLine *TopoBoard,
		debug bool) ([]Snip, *TopoBoard)
}
