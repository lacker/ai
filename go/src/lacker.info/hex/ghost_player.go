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
conformity goes down, so the strength of moves reported goes down.
*/

// For each move that differs from the ghost game, conformity goes
// down this much
const DivergencePenalty = 0.99

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

	// This is 1.0 for games that exactly adhere to our ghost game, and
	// goes down whenever observed moves differ.
	conformity float64
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
	player.conformity = 1.0
	player.index = 0
}

func (player *GhostPlayer) ghostColorAtIndex(i int) Color {
	if i % 2 == 0 {
		return player.startingPosition.GetToMove()
	} else {
		return -player.startingPosition.GetToMove()
	}
}

// Returns the best move to make.
// If this player has nothing to suggest, returns NotASpot.
func (player *GhostPlayer) BestMove(
	board *TopoBoard, debug bool) (TopoSpot, float64) {
	for player.index < len(player.ghostGame) {
		ghostColor := player.ghostColorAtIndex(player.index)
		spot := player.ghostGame[player.index]

		switch board.Get(spot) {
		case ghostColor:
			// Continue with no penalty
			player.index++
		case Empty:
			// Yahtzee
			return spot, player.conformity
		case -ghostColor:
			// Continue with a divergence penalty
			player.conformity *= DivergencePenalty
			player.index++
		default:
			panic("control should not get here")
		}
	}

	return NotASpot, 0.0
}

func (player *GhostPlayer) Debug() {
	log.Printf("%s ghost player prefers:\n", player.color.Name())
	for index, spot := range player.ghostGame {
		if index >= 10 {
			break
		}
		log.Printf("%s: (%d, %d)", player.ghostColorAtIndex(index).Name(),
			spot.Row(), spot.Col())
	}
}
