package hex

import (

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
// TODO: define the formula that maps additive real to
// probability. Logistic, or something simpler?
//
// The QNet is a neural network that operates on a Hex board, and
// incrementally updates with each move to maintain state without
// recalculating the state of each neuron every time.
// The main neural network approximates V(s).
// Also, it maintains Q(s, a) - V(s) for all a in the deltaV
// array. This makes it fast to calculate ideal actions - you don't
// have to run the neural net for every possible action.

type QNet struct {
	startingPosition *TopoBoard

	deltaV [NumTopoSpots]float64
}

// Creates a new qnet that has no values on any features and thus just
// plays random playouts.
func NewQNet(board *TopoBoard) *QNet {
	qnet := &QNet{
		startingPosition: board,
	}
	return qnet
}

// Runs an epsilon-greedy playout.
func (qnet *QNet) Playout(epsilon float64) *QPlayout {
	panic("TODO: implement")
}

// Creates a new qnet that is trained on the provided list of playouts.
func TrainNewQNet(board *TopoBoard, playouts []*QPlayout) *QNet {
	panic("TODO: implement")
}
