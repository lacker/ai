package hex

import (
	"log"
	"math"
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
