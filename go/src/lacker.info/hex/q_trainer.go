package hex

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

// The QTrainer creates QNets for a particular position, and
// repeatedly trains them in order to approximate the best strategy
// with playouts and thus find the best move.
//
// QTrainer implements Player and in some ways is the next generation
// of MetaFarmer.

type QTrainer struct {
	Seconds float64
	Quiet bool

	// The nets we are training, one per player
	whiteNet *QNet
	blackNet *QNet

	// The playouts we have accumulated
	playouts []*QPlayout

	// How many batches we played
	batches int

	// How many games we played
	games int

	// A single logistic estimator for who wins.
	// Black (-1) is negative, White (+1) is positive.
	// Independent from the main neural net.
	winner float64

	// Which player, if any, we are handicapping.
	// This should only be altered right after a batch is processed, so
	// that this value is the same when we learn as it was for all the
	// playouts.
	handicap Color
}

func (trainer *QTrainer) init(b *TopoBoard) {
	trainer.whiteNet = NewQNet(b, White)
	trainer.blackNet = NewQNet(b, Black)
	trainer.playouts = []*QPlayout{}
}

func (trainer *QTrainer) GetToMove() Color {
	return trainer.whiteNet.StartingPosition().GetToMove()
}

func (trainer *QTrainer) NetToMove() *QNet {
	switch trainer.GetToMove() {
	case White:
		return trainer.whiteNet
	case Black:
		return trainer.blackNet
	}
	panic("unhandled switch fallthru")
}

// Plays one game and accumulates the playout
func (trainer *QTrainer) PlayOneGame(debug bool) {
	playout := NewQPlayout(trainer.whiteNet, trainer.blackNet)
	trainer.playouts = append(trainer.playouts, playout)

	// Update the winner
	calcProb := Logistic(trainer.winner)
	var probDiff float64
	if playout.winner == White {
		probDiff = 1.0 - calcProb
	} else {
		probDiff = 0.0 - calcProb
	}
	trainer.winner += 0.3 * probDiff

	trainer.games++
	
	if debug {
		playout.Debug()
	}
}

const DefaultBatchSize int = 1

// Plays a batch, til we have batchSize games.
// This will create a new batch if there is anything in progress.
func (trainer *QTrainer) PlayBatch(batchSize int, debug bool) {
	for len(trainer.playouts) < batchSize {
		trainer.PlayOneGame(false)
	}
}

// Finds the best move and win rate according to the neural net
func (trainer *QTrainer) BestMoveAndWinRate() (TopoSpot, float64) {
	net := trainer.NetToMove()
	net.Reset()
	action := net.IdealAction(net.StartingPosition(), false)
	return action.spot, Logistic(action.Q)
}

// Finds the best move and win rate in practice
func (trainer *QTrainer) BestMoveAndWinRateInPractice() (TopoSpot, float64) {
	// The most frequent move in the last batch should be the best
	var moveCount [NumTopoSpots]int
	var winCount [NumTopoSpots]int
	bestMove := NotASpot
	bestCount := 0
	for _, playout := range trainer.playouts {
		move := playout.FirstMove()
		moveCount[move]++
		if playout.FirstColor() == playout.winner {
			winCount[move]++
		}
		if moveCount[move] > bestCount {
			bestMove = move
			bestCount = moveCount[move]
		}
	}

	if bestMove == NotASpot {
		log.Fatal("empty batch")
	}

	winRate := float64(winCount[bestMove]) /
		float64(moveCount[bestMove])

	return bestMove, winRate
}

// Learns from a batch and resets for the next one.
func (trainer *QTrainer) LearnFromBatch(debug bool) {
	if trainer.handicap != White {
		trainer.whiteNet.LearnFromPlayouts(trainer.playouts, 0.1)
	} else if debug {
		log.Printf("White was handicapped; pausing learning")
	}
	if trainer.handicap != Black {
		trainer.blackNet.LearnFromPlayouts(trainer.playouts, 0.1)
	} else if debug {
		log.Printf("Black was handicapped; pausing learning")
	}
	trainer.batches++

	if debug {
		log.Printf("learned from batch #%d = %d playouts",
			trainer.batches, len(trainer.playouts))
		bestMove, winRate := trainer.BestMoveAndWinRate()
		log.Printf("best move was %v with Q win rate %.3f", bestMove, winRate)
		trainer.Debug()
	}

	// Determine handicapping for the next batch.
	trainer.handicap = Empty
	trainer.whiteNet.handicap = 0.0
	trainer.blackNet.handicap = 0.0
	probWhiteWins := Logistic(trainer.winner)
	// Handicapping starts at 0% handicap for 80% wins and goes to 100%
	// handicap for 100% wins.
	if probWhiteWins > 0.8 {
		trainer.handicap = White
		trainer.whiteNet.handicap = (probWhiteWins - 0.8) * 5.0
		if debug {
			log.Printf("Handicapping White by %.2f", trainer.whiteNet.handicap)
		}
	}
	if probWhiteWins < 0.2 {
		trainer.handicap = Black
		trainer.blackNet.handicap = (0.2 - probWhiteWins) * 5.0
		if debug {
			log.Printf("Handicapping Black by %.2f", trainer.blackNet.handicap)
		}
	}

	trainer.playouts = []*QPlayout{}
}

func (trainer *QTrainer) Debug() {
	log.Printf("played %d games. P(White wins) = %.3f",
		trainer.games, Logistic(trainer.winner))
}

func (trainer *QTrainer) Play(b Board) (NaiveSpot, float64) {
	board := b.ToTopoBoard()
	trainer.init(board)

	if !Debug {
		start := time.Now()
		for SecondsSince(start) < trainer.Seconds {
			trainer.PlayBatch(DefaultBatchSize, false)
			trainer.LearnFromBatch(false)
		}
	} else {
		rand.Seed(1)

		// Logic for debug mode
		keepPlaying := true

		log.Printf("Initial position:")
		board.Debug()

		for keepPlaying {
			// Read a debugger command
			log.Printf("enter command:")
			bio := bufio.NewReader(os.Stdin)
			line, _, _ := bio.ReadLine()
			command := string(line)
			log.Printf("read command: [%s]", command)

			// Handle the command
			switch command {
			case "b":
				// Print the black net
				trainer.blackNet.Debug()
			case "w":
				// Print the white net
				trainer.whiteNet.Debug()
			case "d":
				trainer.Debug()
			case "1":
				trainer.PlayOneGame(true)
			case "l":
				trainer.LearnFromBatch(true)
			case "p":
				trainer.PlayOneGame(true)
				trainer.LearnFromBatch(true)
			case "100":
				for i := 0; i < 100; i++ {
					trainer.PlayOneGame(true)
					trainer.LearnFromBatch(true)
				}
			case "x":
				// finish
				keepPlaying = false

			default:
				var row, col int
				_, err := fmt.Sscanf(command, "%d,%d", &row, &col)
				if err != nil {
					log.Printf("unrecognized command")
				} else {
					spot := MakeTopoSpot(row, col)
					log.Printf("*** White net:")
					trainer.whiteNet.DebugSpot(spot)
					log.Printf("*** Black net:")
					trainer.blackNet.DebugSpot(spot)
				}
			}
		}
	}

	trainer.Debug()

	bestMove, winRate := trainer.BestMoveAndWinRate()
	return bestMove.NaiveSpot(), winRate
}
