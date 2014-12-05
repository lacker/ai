package hex

import (
	"log"
	"testing"
)

func TestDeltaNeuronBasicOperation(t *testing.T) {
	spot1 := MakeTopoSpot(1, 1)
	spot2 := MakeTopoSpot(2, 2)
	spot3 := MakeTopoSpot(3, 3)
	feature1 := BasicFeature{Spot:spot1, Color:Black}
	feature2 := BasicFeature{Spot:spot2, Color:White}
	
	neuron := NewDeltaNeuron([]BasicFeature{feature1, feature2})
	neuron.Bump(spot3, 1337.0)

	var spotPicker [NumTopoSpots]float64
	board := NewTopoBoard()
	registry := NewSpotRegistry()
	neuron.ResetForBoard(board, &spotPicker, registry)

	// Make the first move
	board.MakeMove(spot1)
	registry.Notify(spot1)
	if spotPicker[spot3] != 0.0 {
		log.Fatal("nothing should happen after first move")
	}

  // Make the second move
	board.MakeMove(spot2)
	registry.Notify(spot2)

	if spotPicker[spot3] != 1337.0 {
		log.Fatal("something should happen after second move")
	}
}
