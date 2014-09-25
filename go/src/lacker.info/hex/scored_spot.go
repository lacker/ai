package hex

import (
)

/*
The ScoredSpot is just a convenient way to sort spots according to some score.
*/

type ScoredSpot struct {
	Score float64
	Spot TopoSpot
}

// Make scored spot slices sortable
type ScoredSpotSlice []*ScoredSpot
func (slice ScoredSpotSlice) Len() int {
	return len(slice)
}

func (slice ScoredSpotSlice) Less(i, j int) bool {
	// Use > to implement < to get reverse sorting
	// This is because we want the highest scoring stuff first
	return slice[i].Score > slice[j].Score;
}

func (slice ScoredSpotSlice) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (ss ScoredSpot) Row() int {
	return ss.Spot.ToSpot().Row
}

func (ss ScoredSpot) Col() int {
	return ss.Spot.ToSpot().Col
}
