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
	ending := mf.whitePlayer.Playout(mf.blackPlayer, debug)
	mf.updateWinRate(ending.Winner)

	// Have the loser learn and the winner celebrate
	if ending.Winner == White {
		mf.whitePlayer.LearnFromWin(ending, debug)
		mf.blackPlayer.LearnFromLoss(ending, debug)
	} else {
		mf.blackPlayer.LearnFromWin(ending, debug)
		mf.whitePlayer.LearnFromLoss(ending, debug)
	}
}


func (mf MetaFarmer) Play(b Board) (Spot, float64) {
	start := time.Now()
	mf.init(b.ToTopoBoard())

	if Debug {
		keepPlaying := true
		for keepPlaying {
			// Read a debugger command
			log.Printf("enter command:")
			bio := bufio.NewReader(os.Stdin)
			line, _, _ := bio.ReadLine()
			command := string(line)
			log.Printf("read command: [%s]", command)

			// Handle the command
			endGame := false
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
			case "1000":
				for i := 0; i < 1000; i++ {
					mf.PlayOneGame(false)
				}
				log.Printf("ran a thousand playouts")
			case "x":
				// exit the loop and finish
				keepPlaying = false
			default:
				log.Printf("unrecognized command")
			}
			if endGame {
				break
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
