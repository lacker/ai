package hex

import (
	"log"
	"time"
)

/*
The meta farmer keeps a lot of quick players and does what the best
ones of those do.
*/

type MetaFarmer struct {
	Seconds float64
	Quiet bool

	// The players we are farming
	whitePlayer *QuickPlayer
	blackPlayer *QuickPlayer
}

func (mf MetaFarmer) Play(b Board) (Spot, float64) {
	start := time.Now()

	for SecondsSince(start) < mf.Seconds {
		// ?
	}

	if !mf.Quiet {
		log.Printf("TODO: print something useful here\n")
	}

	panic("TODO")
}
