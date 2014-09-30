package hex

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
	"strings"
)

/*
The LinearPlayer just plays one side of a position, using an algorithm
that just orders spots by preference in playing them. You can play out
a game or learn from it.
*/

const MaxScore float64 = 10000.0

type LinearPlayer struct {
	// The spots we prefer in sorted order
	// -10,000 is the worst possible score
	// 10,000 is the best possible score
	ranking ScoredSpotSlice

	// LinearPlayers always go from the same starting position.
	// The starting position should never be mutated from the
	// LinearPlayer - that way lies only pain.
	startingPosition *TopoBoard

	// What color we play
	color Color

	// The index in the ranking that we're considering next.
	// index only makes sense mid-playout
	index int
}

func MakeLinearPlayer(b *TopoBoard, c Color) *LinearPlayer {
	qp := &LinearPlayer{
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

func (player *LinearPlayer) Color() Color {
	return player.color
}

func (player *LinearPlayer) StartingPosition() *TopoBoard {
	return player.startingPosition
}

// Prepare for a new playout
func (player *LinearPlayer) Reset() {
	player.index = 0
}

// Make one move
func (player *LinearPlayer) MakeMove(board *TopoBoard, debug bool) {
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

// Encodes the spot-ranking as a string
func (player *LinearPlayer) RankingString() string {
	parts := make([]string, len(player.ranking))
	for i, scoredSpot := range player.ranking {
		parts[i] = fmt.Sprintf("%d,%d", scoredSpot.Row(), scoredSpot.Col())
	}
	return strings.Join(parts, ">")
}

// Updates scores based on a game the player lost.
// Heat indicates how much to fluctuate the scores.
// A heat of 1.0 is typical. Over 10k there is diminishing returns
// because basically all ordering will be determined by this
// particular board.
func (player *LinearPlayer) updateScores(board *TopoBoard, heat float64) {
	for _, scoredSpot := range player.ranking {
		// Count all spots played by the winner as a win.
		// Spots not played by either side would also have lost for the
		// loser, so they count as a loss.
		if board.GetTopoSpot(scoredSpot.Spot) == board.Winner {
			scoredSpot.Score += heat * (1.1 - 0.2 * rand.Float64())
		} else {
			scoredSpot.Score -= heat * (1.1 - 0.2 * rand.Float64())
		}
		scoredSpot.Score /= (1.0 + heat / MaxScore)
	}
}

// Randomizes scores and sorts moves in random order
func (player *LinearPlayer) randomize() {
	// Randomize preferences
	for _, scoredSpot := range player.ranking {
		scoredSpot.Score = MaxScore * (rand.Float64() * 2.0 - 1.0)
	}
	sort.Stable(player.ranking)
}

// Learns from a playouted game.
func (player *LinearPlayer) LearnFromWin(board *TopoBoard, debug bool) {
}

// Learns from a playouted game.
func (player *LinearPlayer) LearnFromLoss(board *TopoBoard, debug bool) {
	if board.Winner == Empty {
		log.Fatal("cannot learn from a board with no winner")
	}

	for heat := 1.0; true; heat *= 2.0 {
		player.updateScores(board, heat)
		if !sort.IsSorted(player.ranking) {
			// We learned something. Update our move ordering
			sort.Stable(player.ranking)
			if debug {
				log.Printf("%s has a new mind to try.", player.color.Name())
			}
			return
		}
		if heat > MaxScore {
			// It's impossible to learn anything from this game.
			// Randomize.
			if debug {
				log.Printf("%s cannot learn anything from this game. randomize.",
					player.color.Name())
			}
			player.randomize()
			return
		}
	}
}

// Plays out a game and returns the final board state.
// TODO: delete this
func (player *LinearPlayer) PlayoutDeprecated(
	opponent *LinearPlayer, debug bool) *TopoBoard {

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
func (player *LinearPlayer) Debug() {
	log.Printf("%s linear player prefers:\n", player.color.Name())
	for index, scoredSpot := range player.ranking {
		if index >= 10 {
			break
		}
		log.Printf("(%d, %d) scores %.1f\n",
			scoredSpot.Spot.Row(), scoredSpot.Spot.Col(), scoredSpot.Score)
	}
}