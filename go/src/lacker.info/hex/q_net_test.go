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
}
