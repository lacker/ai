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

// Plays out a game and returns the final board state.
func Playout(
	player1 QuickPlayer, player2 QuickPlayer, debug bool) *TopoBoard {
	return PlayoutWithSnipList(player1, player2, nil, debug)
}

// Plays out a game, overriding the players whenever mandated to by
// the snip list. Returns the final board state.
func PlayoutWithSnipList(
	player1 QuickPlayer, player2 QuickPlayer,
	snipList []Snip, debug bool) *TopoBoard {

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
	snipListIndex := 0

	// Play the playout
	for board.Winner == Empty {
		if snipList != nil && len(snipList) > snipListIndex &&
			snipList[snipListIndex].ply == len(board.History) {
			// The snip list overrides the player
			board.MakeMove(snipList[snipListIndex].spot)
			snipListIndex++
		} else if player1.Color() == board.GetToMove() {
			MakeBestMove(player1, board, debug)
		} else {
			MakeBestMove(player2, board, debug)
		}
	}

	if debug {
		log.Printf("%s wins the playout", board.Winner.Name())
	}
	return board
}
