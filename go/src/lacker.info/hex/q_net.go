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
//
// The main component of the QNet is the QNeuron, which represents a
// set of basic features that add a particular weight to V if all of
// them trigger.

type QNeuron struct {
	// When all of these features activate, weight is added to V
	features []QFeature

	weight float64

	// A bit mask for which of the features are active
	active uint8
}

func MakeQNeuron(features []QFeature, weight float64) QNeuron {
	if len(features) > 8 {
		panic("we can only handle 8 features because we use a bit mask")
	}
	return QNeuron{features:features, weight:weight}
}

type QNet struct {
	startingPosition *TopoBoard
	color Color
	
	// The extra output that would come from activated neurons if each
	// particular action were taken by this color
	deltaV [NumTopoSpots]float64

	// The output solely from the activated neurons
	baseV float64

	// The neurons that make up this net
	neurons []QNeuron

	// The empty spots in the starting position.
	// This is useful for iterating on the spots in random order, which
	// seeds more intelligently than lexicographical spot order.
	emptySpots []TopoSpot

	// The fraction of the time we intentionally go off-policy in order
	// to explore.
	epsilon float64
}

// Creates a new qnet that has no values on any features and thus just
// plays random playouts.
func NewQNet(board *TopoBoard, color Color) *QNet {
	qnet := &QNet{
		startingPosition: board,
		color: color,
		neurons: []QNeuron{},
		emptySpots: board.PossibleTopoSpotMoves(),
		epsilon: 0.05,
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

func (qnet *QNet) Reset() {
	panic("TODO")
}

// Updates the qnet to observe a new feature.
func (qnet *QNet) ObserveNewFeature(feature BasicFeature) {
	panic("TODO")
}

func (qnet *QNet) BestMove(board *TopoBoard, debug bool) (TopoSpot,
	float64) {
	panic("TODO")
}
