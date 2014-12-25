package hex

import (
	"testing"
)

func TestQFeatureConversion(t *testing.T) {
	for _, color := range []Color{Black, White} {
		for spot := TopLeftCorner; spot <= BottomRightCorner; spot++ {
			qf := MakeQFeature(color, spot)
			if qf.Color() != color || qf.Spot() != spot {
				t.Fatalf("bad conversion for %v %v -> %v", color, spot, qf)
			}
		}
	}
}
