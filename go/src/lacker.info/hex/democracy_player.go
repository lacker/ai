package hex

import (
	"log"
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

// TODO: check this works somehow
func (demo *DemocracyPlayer) Add(linear *LinearPlayer) {
	if demo.Color() != linear.Color() {
		log.Fatal("color mismatch")
	}

	if demo.StartingPosition() != linear.StartingPosition() {
		log.Fatal("position mismatch")
	}

	demo.players = append(demo.players, linear)
}

func (demo *DemocracyPlayer) Debug() {
	log.Printf("the demo consists of %d subplayers:\n",
		len(demo.players))
	for _, player := range demo.players {
		player.Debug()
		log.Printf("-----\n")
	}
}

// Make the move that most of the players make
func (demo *DemocracyPlayer) MakeMove(board *TopoBoard, debug bool) {
	if demo.Color() != board.GetToMove() {
		log.Fatal("not our turn to move")
	}

	moveCount := make([]int, NumTopoSpots)
	bestMove := NotASpot
	bestCount := 0

	// Find the most-preferred move
	for _, player := range demo.players {
		move := player.BestMove(board)
		if move == NotASpot {
			continue
		}
		moveCount[move]++
		if moveCount[move] > bestCount {
			bestCount = moveCount[move]
			bestMove = move
		}
	}

	// If we don't have any move, go to fallback
	if bestMove == NotASpot {
		for board.GetTopoSpot(demo.fallbackSpot) != Empty {
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
			bestCount, len(demo.players),
			100.0 * float64(bestCount) / float64(len(demo.players)))
	}

	// Make the move
	board.MakeMove(bestMove)
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
