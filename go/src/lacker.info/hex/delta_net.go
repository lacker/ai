package hex

import (
	"log"
)

// A delta net is a quickplayer that decides what to play by using a bunch
// of delta neurons.
type DeltaNet struct {
	startingPosition *TopoBoard
	color Color

	// This stores the neurons that just operate on a single basic feature
	neurons map[BasicFeature]*DeltaNeuron

	// This stores the default scores for spots.
	// This could be stored as a delta neuron with an empty input list,
	// but this seems simpler.
	defaultScores [NumTopoSpots]float64

	// Always move to this spot if it's available.
	// If this is NotASpot, ignore it.
	// This is useful just to override the first move so that we don't
	// overlearn it. It might help to expand this notion into a whole
	// tree.
	overrideSpot TopoSpot

	spotPicker [NumTopoSpots]float64
}

func NewDeltaNet(board *TopoBoard, color Color) *DeltaNet {
	return &DeltaNet{
		startingPosition: board,
		color: color,
		neurons: make(map[BasicFeature]*DeltaNeuron),
		overrideSpot: NotASpot,
	}
}

func (net *DeltaNet) Reset(game *QuickGame) {
	net.ResetWithBoardAndRegistry(game.board, game.Registry())
}

func (net *DeltaNet) ResetWithBoardAndRegistry(board *TopoBoard,
	registry *SpotRegistry) {
	for i := range net.spotPicker {
		net.spotPicker[i] = net.defaultScores[i]
	}

	for _, neuron := range net.neurons {
		neuron.ResetForBoard(board, &net.spotPicker, registry)
	}
}

func (net *DeltaNet) StartingPosition() *TopoBoard {
	return net.startingPosition
}

func (net *DeltaNet) Debug() {
	log.Printf("%d neurons:", len(net.neurons))
	for feature, neuron := range net.neurons {
		log.Printf("%s => %s", feature.String(), neuron.String())
	}
}

func (net *DeltaNet) Color() Color {
	return net.color
}

func (net *DeltaNet) GetNeuron(feature BasicFeature) *DeltaNeuron {
	neuron, ok := net.neurons[feature]
	if ok {
		return neuron
	}
	neuron = NewDeltaNeuron([]BasicFeature{feature})
	net.neurons[feature] = neuron
	return neuron
}

func (net *DeltaNet) BestMove(board *TopoBoard, debug bool) (TopoSpot,
	float64) {
	if net.overrideSpot != NotASpot && board.Get(net.overrideSpot) != Empty {
		return net.overrideSpot, 1337.0
	}

	bestSpot := NotASpot
	bestScore := -1000000.0
	for spot := TopLeftCorner; spot <= BottomRightCorner; spot++ {
		if net.spotPicker[spot] > bestScore && board.Get(spot) == Empty {
			bestSpot = spot
			bestScore = net.spotPicker[spot]
		}
	}
	
	if bestSpot == NotASpot {
		log.Fatal("best spot should not be NotASpot")
	}
	return bestSpot, bestScore
}

// The learning function
func (net *DeltaNet) EvolveToPlay(ending *TopoBoard, debug bool) {
	if debug {
		log.Printf("evolving %s DeltaNet", net.Color().Name())
	}

	// The range of moves we'll be scanning over
	begin := len(net.startingPosition.History)
	end := len(ending.History)

	// Improve default scores
	for i := begin; i < end; i++ {
		net.defaultScores[ending.History[i]] += 1.0
	}

	// Set the override spot
	if net.startingPosition.GetToMove() == net.color {
		newOverrideSpot := ending.History[begin]
		if net.overrideSpot != newOverrideSpot {
			if debug {
				log.Printf("changing override spot: %v -> %v",
					net.overrideSpot, newOverrideSpot)
			}
			net.overrideSpot = newOverrideSpot
		}
	} else {
		net.overrideSpot = NotASpot
	}

	// Do neuronal learning.
	// The strategy is that we iterate through the game, and every time
	// when we should do the right move, but we don't, we update some
	// features.

	board := net.startingPosition.ToTopoBoard()
	registry := NewSpotRegistry()

	net.ResetWithBoardAndRegistry(board, registry)

	for i := begin; i < end; i++ {
		nextMove := ending.History[i]

		if board.GetToMove() == net.color {
			// Check if we need to train.
			bestMove, bestScore := net.BestMove(board, debug)
			if bestMove != nextMove {
				// We do need to train.
				if debug {
					log.Printf("on move %d, %v should play %v instead of %v",
						i, net.color, nextMove, bestMove)
				}
				missingWeight := bestScore - net.spotPicker[nextMove]
				if missingWeight < 0 {
					log.Fatal("negative missing weight")
				}

				// Find the neurons that are learnable here
				learnable := []*DeltaNeuron{}
				for lookback := 1; lookback <= 2; lookback++ {
					index := i - lookback
					if index < begin {
						break
					}
					feature := ending.FeatureForHistoryIndex(index)
					learnable = append(learnable, net.GetNeuron(feature))
				}

				if len(learnable) == 0 {
					log.Fatal("no learnable neurons")
				}

				bumpSize := (1.0 + missingWeight) / float64(len(learnable))
				for _, neuron := range learnable {
					if debug {
						log.Printf("bumping %v => %v by %.1f",
							neuron, nextMove, bumpSize)
					}
					neuron.Bump(nextMove, bumpSize)
				}
			}
		}
		board.MakeMove(nextMove)
		registry.Notify(nextMove)
	}

	if board.Winner != net.color {
		log.Fatal("ended the game history but we didn't win")
	}
}
