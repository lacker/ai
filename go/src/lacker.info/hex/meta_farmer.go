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

func (mf *MetaFarmer) PlayOneCycle(debug bool) {
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
	linear := NewLinearPlayerFromPlayout(
		evolver.startingPosition, evolver.Color(), ending)

	// Merge the miniplayer into the evolver to evolve it
	evolver.MergeForTheWin(linear, ending.History, debug)

	// At this point, evolver should defeat the opponent with the game
	// ending.History. If that isn't the case this algorithm will subtly
	// corrupt things, so we double-check here if we're in debug mode.
	if debug {
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
				log.Printf("about to run a cycle")
				mf.PlayOneCycle(true)
				log.Printf("ran a cycle")
			case "10", "100", "1000", "10000", "100000", "1000000":
				// Run many cycles
				numCycles, err := strconv.ParseInt(command, 10, 32)
				if err != nil {
					panic("bad number")
				}
				for i := 0; i < int(numCycles); i++ {
					mf.PlayOneCycle(false)
				}
				log.Printf("ran %d cycles", numCycles)
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
		}
	}

	if !mf.Quiet {
		mf.whitePlayer.Debug()
		mf.blackPlayer.Debug()
	}

	// Get the best move based on history
	alreadyMoved := len(mf.whitePlayer.StartingPosition().History)
	bestMove := mf.mainLine.History[alreadyMoved]
	return bestMove.NaiveSpot(), 0.1337
}
