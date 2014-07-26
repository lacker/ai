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

func TestLearningTwoParameters(t *testing.T) {
	n1 := Neuron{}
	n2 := Neuron{}

	// When just n1 fires, the white-odds are 25%.
	// When n1 and n2 fire, the white-odds are 75%.
	for i := 0; i < 2000; i++ {
		var out Color
		if i % 4 == 0 {
			out = White
		} else {
			out = Black
		}

		NeuronBackprop(out, &n1)
		NeuronBackprop(-out, &n1, &n2)
	}

	p1 := NeuronPredict(White, n1)
	error := math.Abs(p1 - 0.25)
	if error > 0.02 {
		t.Errorf("p1 should be ~0.25 but was %f", p1)
	}

	p2 := NeuronPredict(White, n1, n2)
	error = math.Abs(p2 - 0.75)
	if error > 0.02 {
		t.Errorf("p2 should be ~0.75 but was %f", p2)
	}
}
