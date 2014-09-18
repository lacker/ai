package hex

import (
	"log"
	"math/rand"
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

	// ranked keeps the spots in sorted order.
	// The scores start at zero. Spots that lose or aren't useful go
	// negative; spots that win go positive.
	ranked ScoredSpotSlice

	wins int
	losses int
}

// Initialize from a particular board position.
func (s SpotSorter) Init(b Board) {
	// Populate ranked
	s.ranked = make(ScoredSpotSlice, 0)

	moves := b.ToTopoBoard().PossibleTopoSpotMoves()
	for _, move := range moves {
		scoredSpot := &ScoredSpot{Spot: move, Score: 0.0}
		s.ranked = append(s.ranked, scoredSpot)
	}

	s.wins = 0
	s.losses = 0
}

func (s SpotSorter) Play(b Board) (Spot, float64) {
	start := time.Now()

	s.Init(b)
	
	// Run playouts in a loop until we run out of time
	for i := 0; true; i++ {
		// First, sort the possible moves by score.
		sort.Stable(s.ranked)

		// Check if we are out of time
		if SecondsSince(start) > s.Seconds {
			break
		}

		// Run the playout by moving in rank order.
		playout := b.ToTopoBoard()
		for _, move := range s.ranked {
			// On odd runs, sometimes pass, kind of just to introduce some
			// randomness and thus make our learning more robust.
			if i % 2 == 1 && rand.Float64() < 0.05 {
				playout.Pass()
			}

			if !playout.MakeMove(move.Spot.ToSpot()) {
				log.Fatal("a playout played an invalid move")
			}
			if playout.Winner != Empty {
				break
			}
		}

		// Next, update the overall win/loss score.
		// Only do this on even runs, so that we don't count the ones with
		// random fuzzing.
		winner := playout.Winner
		if i % 2 == 0 {
			if winner == b.GetToMove() {
				s.wins++
			} else {
				s.losses++
			}
		}

		// Finally, update the scores for all spots.
		for _, scoredSpot := range s.ranked {
			if playout.Get(scoredSpot.Spot.ToSpot()) == playout.Winner {
				// This counts all spots played by the winner as a win
				scoredSpot.Score += 1.0
			} else {
				scoredSpot.Score -= 1.0
			}
			scoredSpot.Score /= 1.0001
		}
	}

	winRate := float64(s.wins) / float64(s.wins + s.losses)

	if !s.Quiet {
		log.Printf("spot sorter ran %d playouts with win rate %.2f\n",
			s.wins + s.losses, winRate)
		for index, scoredSpot := range s.ranked {
			if index >= 25 {
				break
			}
			log.Printf("(%d, %d) scores %.1f\n",
				scoredSpot.Spot.ToSpot().Row, scoredSpot.Spot.ToSpot().Col,
				scoredSpot.Score)
		}
	}

	// Return the best move
	return s.ranked[0].Spot.ToSpot(), winRate
}
