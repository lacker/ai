package hex

import (
)

// A delta net is a quickplayer that decides what to play by using a bunch
// of delta neurons.
type DeltaNet struct {
	startingPosition *TopoBoard
	color Color

	neurons []DeltaNeuron
	spotPicker [NumTopoSpots]float64
	registry *SpotRegistry
}

func NewDeltaNet(board *TopoBoard, color Color) *DeltaNet {
	return &DeltaNet{
		startingPosition: board,
		color: color,
		neurons: make([]DeltaNeuron, 0),
		registry: NewSpotRegistry(),
	}
}

func (net *DeltaNet) Reset() {
	for i, _ := range net.spotPicker {
		net.spotPicker[i] = 0.0
	}
	net.registry = NewSpotRegistry()

	for _, neuron := range net.neurons {
		neuron.ResetForBoard(net.startingPosition, &net.spotPicker, net.registry)
	}
}

func (net *DeltaNet) StartingPosition() *TopoBoard {
	return net.startingPosition
}

func (net *DeltaNet) Debug() {
}

func (net *DeltaNet) Color() Color {
	return net.color
}

func (net *DeltaNet) BestMove(board *TopoBoard, debug bool) (TopoSpot,
	float64) {
	bestSpot := NotASpot
	bestScore := -1000000.0
	for spot := TopLeftCorner; spot <= BottomRightCorner; spot++ {
		if net.spotPicker[spot] > bestScore {
			bestSpot = spot
			bestScore = net.spotPicker[spot]
		}
	}
	return bestSpot, bestScore
}
