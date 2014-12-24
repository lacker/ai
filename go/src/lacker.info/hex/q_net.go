package hex

import (
	"log"
)

// In reinforcement learning, there are two common functions to learn.
//
// V(s) is the value of being in a state s.
// Q(s, a) is the value of taking action a from state s.
//
// In the case of playing Hex, we are defining "value" to be the
// probability of winning. The network outputs a real number that
// maps onto probabilities - negative for Black winning, positive for
// White winning.
//
// The QNet is a neural network that operates on a Hex board, and
// incrementally updates with each move to maintain state without
// recalculating the state of each neuron every time.
//
// Each QNet corresponds to a particular color. That means it's used
// to decide where that color should play. A QNet tracks two
// things for a particular state:
// A real value baseV
// An array of offsets for each possible action, deltaV[a]
// Q(s, a) is defined as baseV + deltaV(a).
// This makes it easy to choose the best action just by picking the a
// with the highest deltaV.
// Q(s, a) maps to the odds of winning the game.
//
// Another interpretation is that the neural net is calculating a
// function V(s), where it's the value of a state if it's the *other*
// player's turn to move. deltaV is then just tracking how the neural
// network would change with a particular move by this player. This
// explains how the neurons work - they don't add their output values
// directly to baseV; instead when they get one feature away from
// triggering they add their output values to deltaV.
//
// TODO: define how the Q(s, a) -> probability mapping
// works. logistic, or something simpler?


type QNet struct {
	startingPosition *TopoBoard
	color Color

	deltaV [NumTopoSpots]float64
}

// Creates a new qnet that has no values on any features and thus just
// plays random playouts.
func NewQNet(board *TopoBoard, color Color) *QNet {
	qnet := &QNet{
		startingPosition: board,
		color: color,
	}
	return qnet
}

func (qnet *QNet) StartingPosition() *TopoBoard {
	return qnet.startingPosition
}

func (qnet *QNet) Color() Color {
	return qnet.color
}

func (qnet *QNet) Debug() {
	log.Printf("TODO: real qnet debug info")
}

func (qnet *QNet) Reset(game *QuickGame) {
	panic("TODO")
}

func (qnet *QNet) BestMove(board *TopoBoard, debug bool) (TopoSpot,
	float64) {
	panic("TODO")
}
