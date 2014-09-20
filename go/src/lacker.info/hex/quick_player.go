package hex

import (
	"log"
)

/*
The QuickPlayer just plays one side of a position, using an algorithm
that just orders spots by preference in playing them. You can play out
a game or learn from it. Then other algorithms can use lots of
QuickPlayers to do more complicated things on top of that to become
a smarter player.
*/

type QuickPlayer struct {
	// The spots we prefer in sorted order
	// It would be nice if the scoring made some sense.
	ranking ScoredSpotSlice

	// QuickPlayers always go from the same starting position.
	// The starting position should never be mutated from the
	// QuickPlayer - that way lies only pain.
	startingPosition *TopoBoard

	// What color we play
	color Color

	// The index in the ranking that we're considering next.
	// index only makes sense mid-playout
	index int
}

func MakeQuickPlayer(b *TopoBoard, c Color) *QuickPlayer {
	qp := &QuickPlayer{
		ranking: make(ScoredSpotSlice, 0),
		startingPosition: b,
		color: c,
	}

	// Populate the ranking
	moves := qp.startingPosition.PossibleTopoSpotMoves()
	for _, move := range moves {
		scoredSpot := &ScoredSpot{Spot:move, Score:0.0}
		qp.ranking = append(qp.ranking, scoredSpot)
	}

	return qp
}

// Prepare for a new playout
func (player *QuickPlayer) Reset() {
	player.index = 0
}

// Make one move
func (player *QuickPlayer) MakeMove(board *TopoBoard) {
	for player.index < len(player.ranking) {
		spot := player.ranking[player.index].Spot
		player.index++
		if board.GetTopoSpot(spot) == Empty {
			board.SetTopoSpot(spot, player.color)
			return
		}
	}
	log.Fatal("ran out of ranking spots to play")
}

// Learns from a playouted game
func (player *QuickPlayer) Learn(board *TopoBoard) {
	log.Fatal("TODO")
}

// Plays out a game and returns the final board state.
func (player *QuickPlayer) Playout(opponent *QuickPlayer) *TopoBoard {
	if player.color == opponent.color {
		log.Fatal("both players are the same color")
	}

	if player.startingPosition != opponent.startingPosition {
		log.Fatal("starting positions don't match")
	}

	// Prepare for the game.
	// Run the playout on a copy so that we don't alter the original
	board := player.startingPosition.ToTopoBoard()
	player.Reset()
	opponent.Reset()

	// Play the playout
	for board.Winner == Empty {
		if player.color == board.GetToMove() {
			player.MakeMove(board)
		} else {
			opponent.MakeMove(board)
		}
	}

	return board
}
