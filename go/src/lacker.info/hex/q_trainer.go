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

	// The playouts we have accumulated
	playouts []*QPlayout
}

func (trainer *QTrainer) init(b *TopoBoard) {
	trainer.whiteNet = NewQNet(b, White)
	trainer.blackNet = NewQNet(b, Black)
	trainer.playouts = []*QPlayout{}
}

// Plays one game and accumulates the playout
func (trainer *QTrainer) PlayOneGame(debug bool) {
	playout := NewQPlayout(trainer.whiteNet, trainer.blackNet)
	trainer.playouts = append(trainer.playouts, playout)
}

const DefaultBatchSize int = 100

// Plays a batch, til we have batchSize games.
func (trainer *QTrainer) PlayBatch(batchSize int) {
	for len(trainer.playouts) < DefaultBatchSize {
		trainer.PlayOneGame(false)
	}
}

func (trainer *QTrainer) Play(b Board) (NaiveSpot, float64) {
	trainer.init(b.ToTopoBoard())

	panic("TODO")
}
