package hex

import (
	"log"
	"sort"
	"time"
)

/*
The spot sorter algorithm is that you have a score for each spot. You then
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
type SpotSorter struct {
	Seconds float64
	Quiet bool
}

func (s SpotSorter) Play(b Board) (Spot, float64) {
	start := time.Now()
	
	// scores maps each spot to a ScoredSpot for it.
	// The scores start at zero. Spots that lose or aren't useful go
	// negative; spots that win go positive.
	scores := make(map[Spot]*ScoredSpot)

	// ranked keeps the spots in sorted order.
	ranked := make(ScoredSpotSlice, 0)

	// Populate
	moves := b.ToTopoBoard().PossibleTopoSpotMoves()
	for _, move := range moves {
		scoredSpot := &ScoredSpot{Spot: move, Score: 0.0}
		scores[move.ToSpot()] = scoredSpot
		ranked = append(ranked, scoredSpot)
	}


	// Run playouts in a loop until we run out of time
	wins := 0
	losses := 0
	playouts := 0
	for {
		// First, sort the possible moves by score.
		sort.Stable(ranked)

		// Check if we are out of time
		playouts++
		if SecondsSince(start) > s.Seconds {
			break
		}

		// Run the playout by moving in rank order.
		playout := b.ToTopoBoard()
		for _, move := range ranked {
			if !playout.MakeMove(move.Spot.ToSpot()) {
				log.Fatal("a playout played an invalid move")
			}
		}

		// Next, update the scores for all winning spots.
		winner := playout.Winner
		if winner == b.GetToMove() {
			wins++
		} else {
			losses++
		}
		for _, spot := range playout.GetWinningPathSpots() {
			scoredSpot, ok := scores[spot]
			if ok {
				scoredSpot.Score += 2.0
			}
		}

		// Finally, update the scores for all spots.
		for _, scoredSpot := range ranked {
			scoredSpot.Score -= 1.0
			scoredSpot.Score /= 1.01
		}
	}

	// Return the best move
	return ranked[0].Spot.ToSpot(), ranked[0].Score
}
