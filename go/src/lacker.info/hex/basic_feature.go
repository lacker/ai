package hex

import (
	"fmt"
)

// The most basic reasonable feature possible about a board: that a
// particular spot has a particular color.
type BasicFeature struct {
	Color Color
	Spot TopoSpot
}

func (bf BasicFeature) String() string {
	return fmt.Sprintf("%s%s", bf.Color.Name(), bf.Spot.String())
}
