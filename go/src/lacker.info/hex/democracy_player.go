package hex

import (
	"log"
	"math"
)

// The DemocracyPlayer contains a bunch of QuickPlayers and they all
// vote on the best move.
// There is a fallback which just iterates through all possible spots,
// so that even the DemocracyPlayer with no subplayers, or with
// subplayers that have all given up on ideas, will be able to do
// something.

type DemocracyPlayer struct {
	startingPosition *TopoBoard
	color Color

	players []QuickPlayer
	weights []float64
	fallbackSpot TopoSpot
}

func NewDemocracyPlayer(b *TopoBoard, c Color) *DemocracyPlayer {
	dp := &DemocracyPlayer{
		startingPosition: b,
		color: c,
		players: make([]QuickPlayer, 0),
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

func (demo *DemocracyPlayer) Add(quick QuickPlayer) {
	demo.AddWithWeight(quick, 1.0)
}

func (demo *DemocracyPlayer) AddWithWeight(quick QuickPlayer,
	weight float64) {
	if demo.Color() != quick.Color() {
		log.Fatal("color mismatch")
	}

	if demo.StartingPosition() != quick.StartingPosition() {
		log.Fatal("position mismatch")
	}

	demo.players = append(demo.players, quick)
	demo.weights = append(demo.weights, weight)
}

func (demo *DemocracyPlayer) Debug() {
	log.Printf("%s democracy has size %d\n", demo.Color().Name(),
		len(demo.players))
	for i, player := range demo.players {
		if i + 3 >= len(demo.players) {
			log.Printf("Citizen %d has weight %.3f\n", i, demo.weights[i])
			player.Debug()
		}
	}
}

// Returns a heuristic map of spots to how good they ever are to play.
// TODO: reverse this so that high = bad
func (demo *DemocracyPlayer) SpotScoreList() [NumTopoSpots]float64 {
	var scoreList [NumTopoSpots]float64
	for i, p := range demo.players {
		player := p.(*GhostPlayer)
		for _, spot := range player.ghostGame {
			scoreList[spot] += demo.weights[i]
		}
	}
	return scoreList
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

// "quick" should be a quick player that makes the moves that lead
// to targetGame.
// The goal is to merge in quick, similar to what Add would do,
// except ensuring that the weight of quick is high enough so that we
// would play our side of the targetGame after merging.
func (demo *DemocracyPlayer) MergeForTheWin(
	quick QuickPlayer, targetGame []TopoSpot, debug bool) {
	if demo.Color() != quick.Color() {
		log.Fatal("cannot merge wrong color")
	}
	if demo.StartingPosition() != quick.StartingPosition() {
		log.Fatal("cannot merge with different starting positions")
	}

	// Amount we want targetGame to win by
	epsilon := 1.0

	// Minimum amount we need to weigh quick in order to win by epsilon
	delta := 0.0
	
	// We are going to do a playout on a copy.
	board := quick.StartingPosition().ToTopoBoard()
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
			// See what we would do without quick
			bestMove, bestWeight, moveWeight, _ := demo.findBestMove(board)
			if bestMove != nextMove {
				// We'll need some extra weight on quick. How much?
				// Note that we just assume that quick would actually move
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

	if delta == 0.0 && debug {
		log.Printf("loop warning: merging for the win didn't change anything")
	}

	// Merge
	if debug {
		log.Printf("merging with weight %.1f\n", delta)
	}
	demo.AddWithWeight(quick, delta)
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
		move, subweight := player.BestMove(board, false)
		if move == NotASpot {
			continue
		}
		moveWeight[move] += demo.weights[i] * subweight
		totalWeight += demo.weights[i] * subweight
		
		if moveWeight[move] > bestWeight {
			bestWeight = moveWeight[move]
			bestMove = move
		}
	}

	return bestMove, bestWeight, moveWeight, totalWeight
}

// Prefers the move that's voted highest by the players
func (demo *DemocracyPlayer) BestMove(
	board *TopoBoard, debug bool) (TopoSpot, float64) {
	if demo.Color() != board.GetToMove() {
		log.Fatal("not our turn to move")
	}

	bestMove, bestWeight, _, totalWeight := demo.findBestMove(board)
	score := bestWeight / (totalWeight + 0.0001)

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

	return bestMove, score
}

// Drop the player with the least weight
func (demo *DemocracyPlayer) DropLightestPlayer(debug bool) {
	if len(demo.weights) == 0 {
		log.Fatal("can't drop lightest player bc there are no players")
	}
	if len(demo.players) != len(demo.weights) {
		log.Fatal("len players != len weights")
	}
	lightestIndex := 0
	lightestWeight := demo.weights[0]
	for i := 1; i < len(demo.weights); i++ {
		if demo.weights[i] < lightestWeight {
			lightestIndex = i
			lightestWeight = demo.weights[i]
		}
	}

	if debug {
		log.Printf("lightest player has weight %.2f:", demo.weights[lightestIndex])
		demo.players[lightestIndex].Debug()
	}

	demo.players = append(
		demo.players[:lightestIndex],
		demo.players[lightestIndex+1:]...)
	demo.weights = append(
		demo.weights[:lightestIndex],
		demo.weights[lightestIndex+1:]...)
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
