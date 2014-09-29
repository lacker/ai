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

	// Make a move from the provided position.
	// Positions should be progressing through the game until Reset is
	// called, so a quick player can keep some state around.
	MakeMove(board *TopoBoard, debug bool)

	// Prints some debug information
	Debug()

	// Gets what color the player is
	Color() Color
}

// Plays out a game and returns the final board state.
func Playout(
	player1 QuickPlayer, player2 QuickPlayer, debug bool) *TopoBoard {
	
	if player1.Color() == player2.Color() {
		log.Fatal("both players are the same color")
	}

	if player1.StartingPosition() != player2.StartingPosition() {
		log.Fatal("starting positions don't match")
	}

	// Prepare for the game.
	// Run the playout on a copy so that we don't alter the original
	board := player1.StartingPosition().ToTopoBoard()
	player1.Reset()
	player2.Reset()

	// Play the playout
	for board.Winner == Empty {
		if player1.Color() == board.GetToMove() {
			player1.MakeMove(board, debug)
		} else {
			player2.MakeMove(board, debug)
		}
	}

	if debug {
		log.Printf("%s wins the playout", board.Winner.Name())
	}
	return board
}
