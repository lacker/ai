package hex

import (
	"log"
	"sort"
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
	// -10,000 is the worst possible score
	// 10,000 is the best possible score
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
func (player *QuickPlayer) MakeMove(board *TopoBoard, debug bool) {
	for player.index < len(player.ranking) {
		spot := player.ranking[player.index].Spot
		player.index++
		if board.GetTopoSpot(spot) == Empty {
			board.SetTopoSpot(spot, player.color)
			board.ToMove = -board.ToMove
			if debug {
				log.Printf("%s moves %s", player.color.Name(), spot.String())
			}
			return
		}
	}
	log.Fatal("ran out of ranking spots to play")
}

// Learns from a playouted game.
func (player *QuickPlayer) Learn(board *TopoBoard) {
	if board.Winner == Empty {
		log.Fatal("cannot learn from a board with no winner")
	}

	for _, scoredSpot := range player.ranking {
		// Count all spots played by the winner as a win.
		// Spots not played by either side would also have lost for the
		// loser, so they count as a loss.
		if board.GetTopoSpot(scoredSpot.Spot) == board.Winner {
			scoredSpot.Score += 1.0
		} else {
			scoredSpot.Score -= 1.0
		}
		scoredSpot.Score /= 1.0001
	}

	// Sort the possible moves by score
	sort.Stable(player.ranking)
}

// Searches the space of adjacent strategies near this player in order
// to find something that beats the opponent.
// This attempts to mutate the player into something that defeats the
// opponent. However, it might not be able to do so. It returns
// whether it succeeded.
func (player *QuickPlayer) AdaptToOpponent(opponent *QuickPlayer) bool {
	panic("TODO: implement")
}

// Plays out a game and returns the final board state.
func (player *QuickPlayer) Playout(
	opponent *QuickPlayer, debug bool) *TopoBoard {

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
			player.MakeMove(board, debug)
		} else {
			opponent.MakeMove(board, debug)
		}
	}

	if debug {
		log.Printf("%s wins the playout", board.Winner.Name())
	}
	return board
}

// Prints some debug information
func (player *QuickPlayer) Debug() {
	log.Printf("%s quickplayer prefers:\n", player.color.Name())
	for index, scoredSpot := range player.ranking {
		if index >= 10 {
			break
		}
		log.Printf("(%d, %d) scores %.1f\n",
			scoredSpot.Spot.ToSpot().Row, scoredSpot.Spot.ToSpot().Col,
			scoredSpot.Score)
	}
}
