package hex

import (
	"log"
)

// QuickPlayer is an interface for a hex player that is designed to
// run playouts many times per move of a real game.

type QuickPlayer interface {
	// Resets to the starting board position
	Reset()

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

func MakeBestMove(player QuickPlayer, board *TopoBoard, debug bool) {
	if player.Color() != board.GetToMove() {
		log.Fatal("not the right player's turn")
	}
	spot, score := player.BestMove(board, debug)
	board.MakeMove(spot)
	if debug {
		log.Printf("%s moves %s with score %.2f",
			player.Color().Name(), spot.String(), score)
	}
}

