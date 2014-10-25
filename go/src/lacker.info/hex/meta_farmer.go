package hex

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"time"
)

/*
The meta farmer keeps two opposing democracy players. It repeatedly
finds a way for the loser to slightly alter its ways to beat the
winner, with the hope that this converges towards the ideal way to
play.

One cycle is finding a linear player that can defeat the winner, and
then merging this new linear player into the loser hard enough so that
it now wins.

TODO: test and summarize here how well the meta farmer does.
(vs doomedx, ladder, needle, simpleBlock, manyBridges)
*/

type MetaFarmer struct {
	Seconds float64
	Quiet bool

	// The players we are farming
	whitePlayer *DemocracyPlayer
	blackPlayer *DemocracyPlayer

	// What you get when the white player and black player play each
	// other
	mainLine *TopoBoard

	// Whether the game is solved
	gameSolved bool
}

func (mf *MetaFarmer) init(b *TopoBoard) {
	mf.whitePlayer = NewDemocracyPlayer(b, White)
	mf.blackPlayer = NewDemocracyPlayer(b, Black)
	mf.mainLine = Playout(mf.whitePlayer, mf.blackPlayer, false)
}

func (mf *MetaFarmer) Debug() {
	mf.whitePlayer.Debug()
	mf.blackPlayer.Debug()
	log.Printf("Main line:\n")
	mf.mainLine.Debug()
}

func (mf *MetaFarmer) StartingPosition() *TopoBoard {
	return mf.whitePlayer.StartingPosition()
}

// Of the player to move
func (mf *MetaFarmer) WinProbability() float64 {
	toMove := mf.StartingPosition().GetToMove()
	if mf.gameSolved {
		if toMove == mf.mainLine.Winner {
			return 1.0
		} else {
			return 0.0
		}
	}

	// We could probably use a more intelligent heuristic
	return 0.5
}

func (mf *MetaFarmer) PlayOneCycle(debug bool) {
	if mf.gameSolved {
		if debug {
			log.Printf("No need to play one cycle when the game is solved.\n")
		}
		return
	}

	// The plan is to evolve the player who loses the main line.
	// Find who loses the main line
	var evolver *DemocracyPlayer
	var opponent *DemocracyPlayer
	if mf.mainLine.Winner == White {
		opponent = mf.whitePlayer
		evolver = mf.blackPlayer
	} else {
		opponent = mf.blackPlayer
		evolver = mf.whitePlayer
	}

	// Create a miniplayer that beats the opponent
	_, ending := FindWinningSnipList(evolver, opponent, mf.mainLine, debug)
	if ending == nil {
		if debug {
			log.Printf("It's over. %s is unbeatable.\n", opponent.Color().Name())
		}
		mf.gameSolved = true
		return
	}
	linear := NewLinearPlayerFromPlayout(
		evolver.startingPosition, evolver.Color(), ending)

	// Merge the miniplayer into the evolver to evolve it
	evolver.MergeForTheWin(linear, ending.History, debug)

	// At this point, evolver should defeat the opponent with the game
	// ending.History. If that isn't the case this algorithm will subtly
	// corrupt things, so we double-check here if we're in debug mode.
	if debug {
		log.Printf("%s evolved into:\n", evolver.Color().Name())
		evolver.Debug()

		ending2 := Playout(evolver, opponent, false)
		AssertHistoriesEqual(ending.History, ending2.History)
		if ending2.Winner != evolver.Color() {
			log.Fatal("sanity check failed; MergeForTheWin did not achieve win")
		}
	}

	// Update the main line.
	mf.mainLine = ending
}

func (mf MetaFarmer) Play(b Board) (NaiveSpot, float64) {
	start := time.Now()
	mf.init(b.ToTopoBoard())

	if Debug {
		keepPlaying := true

		log.Printf("Initial position:\n")
		mf.StartingPosition().Debug()

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
				mf.PlayOneCycle(true)
			case "10", "100", "1000", "10000", "100000", "1000000":
				// Run many cycles
				numCycles, err := strconv.ParseInt(command, 10, 32)
				if err != nil {
					panic("bad number")
				}
				i := 1
				for ; i <= int(numCycles); i++ {
					mf.PlayOneCycle(false)
					if mf.gameSolved {
						break
					}
				}
				log.Printf("ran %d cycles", i)
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
			mf.PlayOneCycle(false)
			if mf.gameSolved {
				break
			}
		}
	}

	if !mf.Quiet {
		mf.whitePlayer.Debug()
		mf.blackPlayer.Debug()
	}

	// Get the best move based on history
	alreadyMoved := len(mf.StartingPosition().History)
	bestMove := mf.mainLine.History[alreadyMoved]
	return bestMove.NaiveSpot(), mf.WinProbability()
}
