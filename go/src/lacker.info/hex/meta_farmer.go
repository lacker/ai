package hex

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"time"
)

/*
The meta farmer keeps a lot of quick players and does what the best
ones of those do.

With linear players:
mf handles doomed1, doomed2, doomed3 (mostly), ladder.
mf cannot handle needle. doesn't find (5, 6). the losing spots just
alternate.
mf does seem to handle simpleBlock, usually finding (5, 9) there.
but actually given the nature of the opponent, moves like (5, 6) in
needle might not be the best.
so the theory is that the linear player is just not
good enough.
*/

type MetaFarmer struct {
	Seconds float64
	Quiet bool

	// The players we are farming
	whitePlayer *LinearPlayer
	blackPlayer *LinearPlayer

	whiteWinRate float64
	blackWinRate float64
	gamesPlayed int
	lastWinner Color
}

func (mf *MetaFarmer) init(b *TopoBoard) {
	mf.whitePlayer = MakeLinearPlayer(b, White)
	mf.blackPlayer = MakeLinearPlayer(b, Black)
	mf.whiteWinRate = 0.5
	mf.blackWinRate = 0.5
	mf.gamesPlayed = 0
	mf.lastWinner = Empty
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

	mf.lastWinner = winner
}

func (mf *MetaFarmer) Debug() {
	log.Printf("played %d games. white: %.2f black: %.2f\n",
		mf.gamesPlayed, mf.whiteWinRate, mf.blackWinRate)
}

// Play a game
func (mf *MetaFarmer) PlayOneGame(debug bool) {
	ending := Playout(mf.whitePlayer, mf.blackPlayer, debug)
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
			case "10", "100", "1000", "10000", "100000", "1000000":
				// Run many playouts
				numPlayouts, err := strconv.ParseInt(command, 10, 32)
				if err != nil {
					panic("bad number")
				}
				for i := 0; i < int(numPlayouts); i++ {
					mf.PlayOneGame(false)
				}
				log.Printf("ran %d playouts", numPlayouts)
			case "x":
				// exit the loop and finish
				keepPlaying = false
			case "n":
				// Keep playouting until there is a new winner.
				initialWinner := mf.lastWinner

				numGames := 0
				for mf.lastWinner == initialWinner {
					mf.PlayOneGame(false)
					numGames += 1
				}

				log.Printf("ran %d games, until %s won.",
					numGames, mf.lastWinner.Name())

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
