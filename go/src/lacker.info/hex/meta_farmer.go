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

	whiteWinRate float64
	blackWinRate float64
}

func (mf *MetaFarmer) init(b *TopoBoard) {
	// Initialize the metafarmer
	mf.whitePlayer = MakeQuickPlayer(b, White)
	mf.blackPlayer = MakeQuickPlayer(b, Black)
	mf.whiteWinRate = 0.5
	mf.blackWinRate = 0.5
}

func (mf *MetaFarmer) updateWinRate(winner Color) {
	if winner == White {
		mf.whiteWinRate += 0.01
	} else {
		mf.blackWinRate += 0.01
	}
	mf.whiteWinRate /= 1.01
	mf.blackWinRate /= 1.01
}


func (mf *MetaFarmer) Play(b Board) (Spot, float64) {
	start := time.Now()
	mf.init(b.ToTopoBoard())

	for SecondsSince(start) < mf.Seconds {
		// Play a game
		// Have the loser learn
	}

	if !mf.Quiet {
		log.Printf("TODO: print something useful here\n")
	}

	panic("TODO")
}
