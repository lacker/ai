package hex

import (
	"math"
	"testing"
)

func TestBasicNeuron(t *testing.T) {
	n := Neuron{}

	prob := NeuronPredict(Black, n)
	if prob != 0.5 {
		t.Errorf("expected 0.5 black prob, got %f", prob)
	}
}

func TestLearningOneParameter(t *testing.T) {
	n := Neuron{}

	// The neuron should trend towards a 1/4 rate
	for i := 0; i < 1000; i++ {
		var out Color
		if i % 4 == 0 {
			out = White
		} else {
			out = Black
		}
		NeuronBackprop(out, &n)
	}

	p := NeuronPredict(White, n)
	error := math.Abs(p - 0.25)
	if error > 0.02 {
		t.Errorf("p should be ~0.25 but was %f", p)
	}
}
