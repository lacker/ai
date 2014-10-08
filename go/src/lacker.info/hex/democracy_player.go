package hex

import (
	"log"
)

// The DemocracyPlayer contains a bunch of LinearPlayers and they all
// vote on the best move.

type DemocracyPlayer struct {
	startingPosition *TopoBoard
	color Color

	players []*LinearPlayer
}

func MakeDemocracyPlayer(b *TopoBoard, c Color) *DemocracyPlayer {
	dp := &DemocracyPlayer{
		startingPosition: b,
		color: c,
		players: make([]*LinearPlayer, 0),
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

// TODO: make the move that most of the players make
func (demo *DemocracyPlayer) MakeMove(board *TopoBoard, debug bool) {
	if len(demo.players) < 1 {
		log.Fatal("cannot make a move in a democracy with no players")
	}
	moves := make([]TopoSpot, len(demo.players))
	for i, player := range demo.players {
		moves[i] = player.BestMove(board)
	}

	panic("TODO: implement")
}
