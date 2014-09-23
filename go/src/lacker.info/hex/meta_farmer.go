package hex

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	gamesPlayed int
}

func (mf *MetaFarmer) init(b *TopoBoard) {
	mf.whitePlayer = MakeQuickPlayer(b, White)
	mf.blackPlayer = MakeQuickPlayer(b, Black)
	mf.whiteWinRate = 0.5
	mf.blackWinRate = 0.5
	mf.gamesPlayed = 0
}

func (mf *MetaFarmer) updateWinRate(winner Color) {
	mf.gamesPlayed++
	if winner == White {
		mf.whiteWinRate += 0.001
	} else {
		mf.blackWinRate += 0.001
	}
	mf.whiteWinRate /= 1.001
	mf.blackWinRate /= 1.001
}


func (mf MetaFarmer) Play(b Board) (Spot, float64) {
	start := time.Now()
	mf.init(b.ToTopoBoard())

	for i := 0; true; i++ {
		if Debug {
			// Read a debugger command
			fmt.Printf("> ")
			bio := bufio.NewReader(os.Stdin)
			line, _, _ := bio.ReadLine()
			command := string(line)
			log.Printf("read command: [%s]\n", command)
		} else {
			// Check if we are out of time
			if SecondsSince(start) > mf.Seconds {
				break
			}
		}

		// Play a game
		ending := mf.whitePlayer.Playout(mf.blackPlayer)
		mf.updateWinRate(ending.Winner)

		// Have the loser learn
		if ending.Winner == White {
			mf.whitePlayer.Learn(ending)
		} else {
			mf.blackPlayer.Learn(ending)
		}
	}

	if !mf.Quiet {
		log.Printf("played %d games. white: %.2f black: %.2f\n",
			mf.gamesPlayed, mf.whiteWinRate, mf.blackWinRate)
		mf.whitePlayer.Debug()
		mf.blackPlayer.Debug()
	}

	switch b.GetToMove() {
	case White:
		return mf.whitePlayer.ranking[0].Spot.ToSpot(), mf.whiteWinRate
	case Black:
		return mf.blackPlayer.ranking[0].Spot.ToSpot(), mf.blackWinRate
	default:
		panic("there is no player to move")
	}
}
