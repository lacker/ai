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

func (demo *DemocracyPlayer) Add(linear *LinearPlayer) {
	if demo.Color() != linear.Color() {
		log.Fatal("color mismatch")
	}

	if demo.StartingPosition() != linear.StartingPosition() {
		log.Fatal("position mismatch")
	}

	panic("TODO: implement")
}
