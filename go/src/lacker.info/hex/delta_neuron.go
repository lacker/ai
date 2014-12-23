package hex

import (
	"fmt"
	"log"
)

// A delta neuron keeps its state so that updating based on a single
// feature changing is relatively fast.
// It is either active, inactive, or deactivated. It starts
// inactive. If all of its input features are active, it becomes
// active. If there's a conflicting feature so that it can never
// become active, it deactivates.
// When it activates it bumps scores for a set of spots so that in the
// end, one particular spot will score the best.
// Thus, the DeltaNeuron constructs a sort of net that can predict
// which spot is the best spot to move, but cannot predict who is
// eventually likely to win the game.
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

	// The registry for moves.
	// Specific to one playout.
	registry *SpotRegistry
}

func NewDeltaNeuron(features []BasicFeature) *DeltaNeuron {
	return &DeltaNeuron{
		input: features,
		output: []ScoredSpot{},
	}
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
			dn.registry.Listen(feature.Spot, dn)
			return
		case -feature.Color:
			// Deactivate
			return
		default:
			log.Fatal("flow shouldn't get here")
		}
	}

	// Activate! This bumps scores in the spot picker.
	dn.active = true
	for _, scoredSpot := range dn.output {
		dn.spotPicker[scoredSpot.Spot] += scoredSpot.Score
	}
}

func (dn *DeltaNeuron) HandleNotification(spot TopoSpot) {
	dn.ContinueActivation()
}

// Get ready for a new playout on a new board.
// This board should always be a fresh clone of the same state.
func (dn *DeltaNeuron) ResetForBoard(board *TopoBoard,
	spotPicker *[NumTopoSpots]float64, registry *SpotRegistry) {
	dn.active = false
	dn.board = board
	dn.featureIndex = 0
	dn.spotPicker = spotPicker
	dn.registry = registry

	dn.ContinueActivation()
}

func (dn *DeltaNeuron) Bump(spot TopoSpot, score float64) {
	for i, scoredSpot := range dn.output {
		if scoredSpot.Spot == spot {
			dn.output[i].Score += score
			return
		}
	}
	dn.output = append(dn.output, ScoredSpot{Spot:spot, Score:score})
}

func (dn *DeltaNeuron) String() string {
	return fmt.Sprintf("[input:%v, output:%v]", dn.input, dn.output)
}
