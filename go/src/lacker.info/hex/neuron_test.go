package hex

import (
	"testing"
)

func TestBasicNeuron(t *testing.T) {
	n := Neuron{}

	prob := NeuronPredict(Black, n)
	if prob != 0.5 {
		t.Errorf("expected 0.5 black prob, got %f", prob)
	}
}
