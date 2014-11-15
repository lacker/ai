package hex

import (

)

// The most basic reasonable feature possible about a board: that a
// particular spot has a particular color.
type BasicFeature struct {
	Color Color
	Spot TopoSpot
}


