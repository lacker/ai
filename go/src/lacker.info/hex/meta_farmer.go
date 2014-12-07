package hex

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"time"
)

/*
The meta farmer keeps two opposing quickplayers. It repeatedly
finds a way for the loser to slightly alter its ways to beat the
winner, with the hope that this converges towards the ideal way to
play.

One cycle is finding a quick player that can defeat the winner, and
then merging this new quick player into the loser hard enough so that
it now wins.

The meta farmer with democracy players:
Solves doomed1, doomed2, doomed3
Cannot solve doomed4, triangleBlock
Moves correctly on ladder, manyBridges but does not totally solve
Moves correctly on needle, simpleBlock

The meta farmer with delta nets:
TODO
*/

type MetaFarmer struct {
	Seconds float64
	Quiet bool

	// The players we are farming
	whitePlayer EvolvingPlayer
	blackPlayer EvolvingPlayer

	// What you get when the white player and black player play each
	// other
	mainLine *TopoBoard

	// The type of quickplayer to use.
	QuickType string

	// Whether the game is solved
	gameSolved bool

	// The number of evolution cycles that have run
	cycles int
}

func (mf *MetaFarmer) init(b *TopoBoard) {
	switch mf.QuickType {
	case "democracy":
		mf.whitePlayer = NewDemocracyPlayer(b, White)
		mf.blackPlayer = NewDemocracyPlayer(b, Black)
	case "deltanet":
		mf.whitePlayer = NewDeltaNet(b, White)
		mf.blackPlayer = NewDeltaNet(b, Black)
	default:
		log.Fatalf("invalid QuickType: %s", mf.QuickType)
	}
	mf.mainLine = Playout(mf.whitePlayer, mf.blackPlayer, false)
}

func (mf *MetaFarmer) Debug() {
	mf.whitePlayer.Debug()
	mf.blackPlayer.Debug()
	log.Printf("%d cycles have been played. Main line:", mf.cycles)
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
	mf.cycles++
	if mf.gameSolved {
		if debug {
			log.Printf("The game is solved but we can try anyway.")
		}
	}

	// The plan is to evolve the player who loses the main line.
	// Find who loses the main line
	var evolver EvolvingPlayer
	var opponent EvolvingPlayer
	if mf.mainLine.Winner == White {
		opponent = mf.whitePlayer
		evolver = mf.blackPlayer
	} else {
		opponent = mf.blackPlayer
		evolver = mf.whitePlayer
	}

	// Create a miniplayer that beats the opponent
	var ending *TopoBoard
	_, ending = FindWinningSnipList(evolver, opponent, mf.mainLine, debug)
	if ending == nil {
		if debug {
			log.Printf("It's over. %s is unbeatable.\n", opponent.Color().Name())
		}
		mf.gameSolved = true
		return
	}

	if debug {
		log.Printf("New main line:")
		for i := len(evolver.StartingPosition().History);
		i < len(ending.History); i++ {
			log.Printf("%v %d: %v", ending.ColorForHistoryIndex(i),
				i, ending.History[i])
		}
	}

	evolver.EvolveToPlay(ending, debug)

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
	mf.init(b.ToTopoBoard())

	if Debug {
		keepPlaying := true

		log.Printf("Initial position:")
		mf.StartingPosition().Debug()
		log.Printf("Initial main line:")
		mf.mainLine.Debug()

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
				start := time.Now()
				mf.PlayOneCycle(true)
				log.Printf("one cycle took %.2f seconds", SecondsSince(start))
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
			case "5s":
				// Run for five seconds
				start := time.Now()
				i := 0
				for SecondsSince(start) < mf.Seconds {
					mf.PlayOneCycle(false)
					i++
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
		start := time.Now()
		for SecondsSince(start) < mf.Seconds {
			mf.PlayOneCycle(false)
			if mf.gameSolved {
				break
			}
		}
	}

	// Get the best move based on history
	alreadyMoved := len(mf.StartingPosition().History)
	bestMove := mf.mainLine.History[alreadyMoved]
	return bestMove.NaiveSpot(), mf.WinProbability()
}
