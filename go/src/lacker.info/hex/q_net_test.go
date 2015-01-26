package hex

import (
	"log"
	"math"
	"math/rand"
	"testing"
)

func CheckApproxEq(x float64, y float64) {
	if math.Abs(x - y) > 0.0001 {
		log.Fatalf("%.5f too far from %.5f", x, y)
	}
}

func TestLogisticMath(t *testing.T) {
	CheckApproxEq(Logistic(0.0), 0.5)
	probs := []float64{0.0, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0}
	for _, prob := range probs {
		CheckApproxEq(Logistic(Logit(prob)), prob)
	}
}

func TestNeuronActivationWithReset(t *testing.T) {
	spot1 := MakeTopoSpot(1, 1)
	spot2 := MakeTopoSpot(2, 2)
	feature1 := MakeQFeature(Black, spot1)
	feature2 := MakeQFeature(White, spot2)

	rand.Seed(1)
	board := NewTopoBoard()
	net := NewQNet(board, Black)
	net.AddFeature(feature1)
	net.AddFeature(feature2)
	if net.GetNeuron(feature1, feature2).active != 2 {
		t.Fatal("expected 2")
	}
	net.Reset()
	if net.GetNeuron(feature1, feature2).active != 0 {
		t.Fatal("expected 0")
	}
	net.AddFeature(feature1)
}
