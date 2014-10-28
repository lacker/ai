package hex

import (
)

// NaiveSpot and TopoSpot should both be implementations of Spot.
// NaiveSpot is, by test.
// TopoSpot is not yet. TODO
type Spot interface {
	GetRow() int
	GetCol() int
}
