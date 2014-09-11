package hex

import (
	"time"
)

/*
The spot ranker algorithm is that you have a score for each spot. You then
do playouts ranking the spots in the given order, and you update the
scores according to which spots are winning so that spots that win get
higher scores.

The theory is that eventually this should converge to something more
intelligent than a shallow rave algorithm which doesn't do any move
after the first one intelligently.
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

// The player
type SpotRanker struct {
	Seconds float64
	Quiet bool
}

func (s SpotRanker) Play(b Board) (Spot, float64) {
	start := time.Now()
	
	// scores maps each spot to a ScoredSpot for it.
	// The scores start at zero. Spots that lose or aren't useful go
	// negative; spots that win go positive.
	scores := make(map[TopoSpot]*ScoredSpot)

	// ranked keeps the spots in sorted order.
	ranked := make([]*ScoredSpot, 0)

	// Populate
	moves := b.ToTopoBoard().PossibleTopoSpotMoves()
	for _, move := range moves {
		scoredSpot := &ScoredSpot{Spot: move, Score: 0.0}
		scores[move] = scoredSpot
		ranked = append(ranked, scoredSpot)
	}


	// Run playouts in a loop until we run out of time
	playouts := 0
	for {
		// First, sort the possible moves by score.
		// TODO

		// Check if we are out of time
		playouts++
		if SecondsSince(start) > s.Seconds {
			break
		}

		// Run the playout by moving in rank order.
		// TODO

		// Next, update the scores for all winning spots.
		// TODO

		// Finally, update the scores for all spots.
		// TODO
	}

	// Return the best move
	return ranked[0].Spot.ToSpot(), ranked[0].Score
}
