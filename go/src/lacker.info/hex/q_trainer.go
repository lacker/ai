package hex

import (

)

// The QTrainer creates QNets for a particular position, and
// repeatedly trains them in order to approximate the best strategy
// with playouts and thus find the best move.
//
// QTrainer implements Player and in some ways is the next generation
// of MetaFarmer.

type QTrainer struct {
	Seconds float64
	Quiet bool

	// The nets we are training, one per player
	whiteNet *QNet
	blackNet *QNet
}

func (trainer QTrainer) Play(b Board) (NaiveSpot, float64) {
	panic("TODO")
}
