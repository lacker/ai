package hex

import (

)

// A delta neuron keeps its state so that updating based on a single
// feature changing is relatively fast.
// It is either active or inactive. It's active once its entire list
// of input features is active.
type DeltaNeuron struct {
	// Inputs that lead this neuron to be active.
	// This persists across playouts.
	input []BasicFeature

	// Weights on output. Positive encourages, negative discourages.
	// This persists across playouts.
	output []ScoredSpot

	// The board we are using for this specific playout.
	board *TopoBoard

	// The spot heap we are using for this specific playout.
	// spotHeap *SpotHeap
}

// Get ready for a new playout on a new board.
// This board should always be a fresh clone of the same state.
func (*DeltaNeuron) ResetForBoard(board *TopoBoard) {
	panic("TODO: some listener stuff. also use a spot heap")
}
