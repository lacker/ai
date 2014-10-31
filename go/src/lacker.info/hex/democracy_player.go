package hex

import (
	"log"
	"math"
)

// The DemocracyPlayer contains a bunch of LinearPlayers and they all
// vote on the best move.
// There is a fallback which just iterates through all possible spots,
// so that even the DemocracyPlayer with no LinearPlayers, or with
// LinearPlayers that have all given up on ideas, will be able to do
// something.

type DemocracyPlayer struct {
	startingPosition *TopoBoard
	color Color

	players []*LinearPlayer
	weights []float64
	fallbackSpot TopoSpot
}

func NewDemocracyPlayer(b *TopoBoard, c Color) *DemocracyPlayer {
	dp := &DemocracyPlayer{
		startingPosition: b,
		color: c,
		players: make([]*LinearPlayer, 0),
		fallbackSpot: TopLeftCorner,
	}
	return dp
}

func (demo *DemocracyPlayer) Color() Color {
	return demo.color
}

func (demo *DemocracyPlayer) StartingPosition() *TopoBoard {
	return demo.startingPosition
}

func (demo *DemocracyPlayer) Add(linear *LinearPlayer) {
	demo.AddWithWeight(linear, 1.0)
}

func (demo *DemocracyPlayer) AddWithWeight(linear *LinearPlayer,
	weight float64) {
	if demo.Color() != linear.Color() {
		log.Fatal("color mismatch")
	}

	if demo.StartingPosition() != linear.StartingPosition() {
		log.Fatal("position mismatch")
	}

	demo.players = append(demo.players, linear)
	demo.weights = append(demo.weights, weight)
}

func (demo *DemocracyPlayer) Debug() {
	log.Printf("%s democracy has size %d\n", demo.Color().Name(),
		len(demo.players))
	for i, player := range demo.players {
		log.Printf("Citizen %d has weight %.3f\n", i, demo.weights[i])
		player.Debug()
	}
}

// Makes the weights sum to 10,000 because that's a nice number.
// If there is no weight this is just a no-op.
func (demo *DemocracyPlayer) NormalizeWeights() {
	totalWeight := 0.0
	for _, w := range demo.weights {
		totalWeight += w
	}
	if totalWeight <= 0.0 {
		return
	}
	for i, w := range demo.weights {
		demo.weights[i] = w * 10000.0 / totalWeight
	}
}

// "linear" should be a linear player that makes the moves that lead
// to targetGame.
// The goal is to merge in linear, similar to what Add would do,
// except ensuring that the weight of linear is high enough so that we
// would play our side of the targetGame after merging.
func (demo *DemocracyPlayer) MergeForTheWin(
	linear *LinearPlayer, targetGame []TopoSpot, debug bool) {
	if demo.Color() != linear.Color() {
		log.Fatal("cannot merge wrong color")
	}
	if demo.StartingPosition() != linear.StartingPosition() {
		log.Fatal("cannot merge with different starting positions")
	}

	// Amount we want targetGame to win by
	epsilon := 0.1

	// Minimum amount we need to weigh linear in order to win by epsilon
	delta := 0.0
	
	// We are going to do a playout on a copy.
	board := linear.StartingPosition().ToTopoBoard()
	demo.Reset()

	// Play the playout
	for board.Winner == Empty {
		// Figure out what the next move should be
		nextMoveIndex := len(board.History)
		if nextMoveIndex >= len(targetGame) {
			log.Fatal("ran off the end of the target game")
		}
		nextMove := targetGame[nextMoveIndex]

		if board.GetToMove() == demo.Color() {
			// We need to train on this move.
			// See what we would do without linear
			bestMove, bestWeight, moveWeight, _ := demo.findBestMove(board)
			if bestMove != nextMove {
				// We'll need some extra weight on linear. How much?
				// Note that we just assume that linear would actually move
				// according to the target game. If that isn't the case our
				// output will get corrupted and meaningless.
				nextMoveWeight := moveWeight[nextMove]
				missingWeight := bestWeight - nextMoveWeight
				if missingWeight < 0.0 {
					log.Fatal("unclear why the best move was the best move")
				}
				delta = math.Max(delta, missingWeight + epsilon)
			}
		}
		board.MakeMove(nextMove)
	}

	// This only made sense before doing a DropLightestPlayer
	/*
	if delta == 0.0 {
		log.Fatal("merging for the win didn't even change anything")
	}
  */

	// Merge
	if debug {
		log.Printf("merging with weight %.1f\n", delta)
	}
	demo.AddWithWeight(linear, delta)
	demo.NormalizeWeights()
}

// Find the move that most of the players like.
// Returns the best move, the weight on it, the array of weights, and
// the total weight.
// If there is no weight on any move, this will return NotASpot and 0.
func (demo *DemocracyPlayer) findBestMove(
	board *TopoBoard) (TopoSpot, float64, []float64, float64) {
	bestMove := NotASpot
	bestWeight := 0.0
	moveWeight := make([]float64, NumTopoSpots)
	totalWeight := 0.0

	// Find the most-preferred move
	for i, player := range demo.players {
		move := player.BestMove(board)
		if move == NotASpot {
			continue
		}
		moveWeight[move] += demo.weights[i]
		totalWeight += demo.weights[i]
		
		if moveWeight[move] > bestWeight {
			bestWeight = moveWeight[move]
			bestMove = move
		}
	}

	return bestMove, bestWeight, moveWeight, totalWeight
}

// Make the move that most of the players make
func (demo *DemocracyPlayer) MakeMove(board *TopoBoard, debug bool) {
	if demo.Color() != board.GetToMove() {
		log.Fatal("not our turn to move")
	}

	bestMove, bestWeight, _, totalWeight := demo.findBestMove(board)

	// If we don't have any move, go to fallback
	if bestMove == NotASpot {
		for board.Get(demo.fallbackSpot) != Empty {
			demo.fallbackSpot++
		}
		bestMove = demo.fallbackSpot
		if debug {
			log.Printf("%s moves at the fallback: %s",
				demo.color.Name(), bestMove.String())
		}
	} else if debug {
		log.Printf("%s moves %s, which scored %d out of %d = %.1f%%",
			demo.color.Name(), bestMove.String(),
			bestWeight, len(demo.players),
			100.0 * bestWeight / totalWeight)
	}

	// Make the move
	board.MakeMove(bestMove)
}

// Drop the player with the least weight
func (demo *DemocracyPlayer) DropLightestPlayer() {
	if len(demo.weights) == 0 {
		log.Fatal("can't drop lightest player bc there are no players")
	}
	lightestIndex := 0
	lightestWeight := demo.weights[0]
	for i := 1; i < len(demo.weights); i++ {
		if demo.weights[i] < lightestWeight {
			lightestIndex = i
			lightestWeight = demo.weights[i]
		}
	}

	demo.players = append(
		demo.players[lightestIndex:],
		demo.players[:lightestIndex]...)
}

// Prepare for a new playout
func (demo *DemocracyPlayer) Reset() {
	for _, player := range demo.players {
		player.Reset()
	}
	demo.fallbackSpot = TopLeftCorner
}

// Limit to only a certain number of players by cutting the old ones
func (demo *DemocracyPlayer) Truncate(limit int) {
	numToCut := len(demo.players) - limit
	if numToCut <= 0 {
		return
	}

	demo.players = demo.players[numToCut:]
}
