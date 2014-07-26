package hex

import (
	"math"
)

/*
This Neuron is a simple predictor of a color: black or white.
This implementation is a logistic neuron.
This is basically an input neuron - it's assumed that the caller will
know which neurons are activating and use that.
*/
type Neuron struct {
	// A very high logit correlates to confidence in White.
	// A very negative logit correlates to confidence in Black.
	// A logit of zero indicates no information either way.
	Logit float64
}

/*
Predicts the odds of a particular color.
*/
func NeuronPredict(color Color, neurons ...Neuron) float64 {
	sumLogits := 0.0
	for _, neuron := range neurons {
		sumLogits += neuron.Logit
	}
	whiteOdds := 1.0 / (1.0 + math.Exp(-sumLogits))

	switch color {
	case White:
		return whiteOdds
	case Black:
		return 1.0 - whiteOdds
	case Empty:
		panic("can't get probability of empty")
	}
	panic("control should not reach here")
}

/*
Updates neurons used for a prediction.
Uses gradient ascent.
Typically this is called after NeuronPredict if we discover what the
right answer was, with the same args we used on NeuronPredict.
This update rule only really makes sense for single-layer neural
networks.
*/
func NeuronBackprop(color Color, neurons ...Neuron) {
	
}
