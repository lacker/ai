package hex

import (
	"log"
)

/*
The GhostPlayer is like a ghost racer in Mario Kart.
It has a single game memorized, and it tries to play along with that
game.
If the color that plays on a particular spot in the ghost game
differs from the color that plays in that spot in the real game, the
ghost player becomes divergent and stops suggesting moves.
If the opponent played at a spot in the ghost game but it is empty in
the real game, that doesn't cause divergence; we just move on.
The ghost player always tries to make the next move that
its color made in the ghost game.
*/

type GhostPlayer struct {
	// Quickplayers always go from the same starting position
	startingPosition *TopoBoard

	// The sequence of moves in the ghost game
	// It is assumed that this alternates colors starting with the
	// player to move in startingPosition
	ghostGame []TopoSpot

	// What color we play
	color Color

	// The index in the ranking that is the next move that happened in
	// the ghost game
	index int

	// Whether the real game has diverged from the ghost game.
	// Once the ghost becomes divergent it will no longer suggest moves
	divergent bool
}

// Create a new ghost player from a finished game
func NewGhostPlayer(b *TopoBoard, c Color, ending *TopoBoard) *GhostPlayer {
	if len(b.History) >= len(ending.History) {
		log.Fatal("len b history >= len ending history")
	}

	gp := &GhostPlayer{
		ghostGame: ending.History[len(b.History):],
		startingPosition: b,
		color: c,
	}
	gp.Reset()
	return gp
}

func (player *GhostPlayer) Color() Color {
	return player.color
}

func (player *GhostPlayer) StartingPosition() *TopoBoard {
	return player.startingPosition
}

func (player *GhostPlayer) Reset() {
	player.divergent = false
	player.index = 0
}

// Returns the best move to make.
// If this player has nothing to suggest, returns NotASpot.
func (player *GhostPlayer) BestMove(board *TopoBoard) TopoSpot {
	if player.divergent {
		return NotASpot
	}

	for player.index < len(player.ghostGame) {
		// See what color the ghost was at this point
		var ghostColor Color
		if player.index % 2 == 0 {
			ghostColor = player.startingPosition.GetToMove()
		} else {
			ghostColor = -player.startingPosition.GetToMove()
		}

		spot := player.ghostGame[player.index]
		if ghostColor == player.color {
			// This is a move we made in the ghost game.
			switch board.Get(spot) {
			case player.color:
				// The real game already moved here. Keep looking
				player.index++
			case Empty:
				// Yahtzee
				return spot
			case -player.color:
				// Diverge
				player.divergent = true
				return NotASpot
			}
		} else {
			// This is a move our opponent made in the ghost game.
			switch board.Get(spot) {
			case player.color:
				// Diverge
				player.divergent = true
				return NotASpot
			case Empty:
				// Nobody has moved here. This isn't unusual enough to cause
				// divergence, so just continue.
				player.index++
			case -player.color:
				// The real game already moved here. Keep looking
				player.index++
			}
		}
	}

	panic("We shouldn't get here - we should either win or diverge.")
}
