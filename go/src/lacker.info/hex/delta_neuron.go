package hex

import (
	"log"
)

// A delta neuron keeps its state so that updating based on a single
// feature changing is relatively fast.
// It is either active, inactive, or deactivated. It starts
// inactive. If all of its input features are active, it becomes
// active. If there's a conflicting feature so that it can never
// become active, it deactivates.
type DeltaNeuron struct {
	// Inputs that lead this neuron to be active.
	// This persists across playouts.
	input []BasicFeature

	// Weights on output. Positive encourages, negative discourages.
	// This persists across playouts.
	output []ScoredSpot

	// The board we are using for this specific playout.
	board *TopoBoard

	// The map of spots to scores that will determine moves.
	// This is specific to one playout.
	spotPicker *[NumTopoSpots]float64

	// The next feature that we need to confirm to activate.
	// Just iterates through the features.
	// This is specific to one playout.
	featureIndex int

	// Whether this neuron has been activated.
	// Specific to one playout.
	active bool
}

// See if we should become active.
// If we are still inactive, this arranges for ContinueActivation to
// be called in the future as well.
func (dn *DeltaNeuron) ContinueActivation() {
	if dn.active {
		log.Fatal("shouldn't double-activate a neuron")
	}

	for dn.featureIndex < len(dn.input) {
		// Check this feature
		feature := dn.input[dn.featureIndex]
		switch dn.board.Get(feature.Spot) {
		case feature.Color:
			// This input is active, continue to the next one
			dn.featureIndex++
		case Empty:
			// We need to keep listening for changes here
			panic("TODO: how to listen")
		case -feature.Color:
			// Deactivate
			return
		default:
			log.Fatal("flow shouldn't get here")
		}
	}

	// Activate
	dn.active = true
	panic("TODO: do something when activating")
}

// Get ready for a new playout on a new board.
// This board should always be a fresh clone of the same state.
func (dn *DeltaNeuron) ResetForBoard(board *TopoBoard,
	spotPicker *[NumTopoSpots]float64) {
	panic("TODO: some listener stuff. also use a spot heap")
}
