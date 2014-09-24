package hex

import (
	"bufio"
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

func (mf *MetaFarmer) Debug() {
	log.Printf("played %d games. white: %.2f black: %.2f\n",
		mf.gamesPlayed, mf.whiteWinRate, mf.blackWinRate)
}

func (mf *MetaFarmer) PlayOneGame(debug bool) {
	// Play a game
	ending := mf.whitePlayer.Playout(mf.blackPlayer)
	mf.updateWinRate(ending.Winner)

	// Have the loser learn
	if ending.Winner == White {
		if debug {
			log.Printf("white is learning")
		}
		mf.whitePlayer.Learn(ending)
	} else {
		if debug {
			log.Printf("black is learning")
		}
		mf.blackPlayer.Learn(ending)
	}
}


func (mf MetaFarmer) Play(b Board) (Spot, float64) {
	start := time.Now()
	mf.init(b.ToTopoBoard())

	if Debug {
		for {
			// Read a debugger command
			log.Printf("enter command:")
			bio := bufio.NewReader(os.Stdin)
			line, _, _ := bio.ReadLine()
			command := string(line)
			log.Printf("read command: [%s]", command)

			// Handle the command
			switch command {
			case "b":
				// Print what black is thinking
				mf.blackPlayer.Debug()
			case "w":
				// Print what white is thinking
				mf.whitePlayer.Debug()
			case "s":
				// Print overall status
				mf.Debug()
			case "1":
				// Run one playout
				mf.PlayOneGame(true)
				log.Printf("ran a playout")
			case "x":
				// exit the loop and finish
				break
			default:
				log.Printf("unrecognized command")
			}
		}
	} else {
		for SecondsSince(start) < mf.Seconds {
			mf.PlayOneGame(false)
		}
	}

	if !mf.Quiet {
		mf.Debug()
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
