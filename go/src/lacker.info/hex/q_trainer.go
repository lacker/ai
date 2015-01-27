package hex

import (
	"bufio"
	"log"
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
}

func (trainer *QTrainer) init(b *TopoBoard) {
	trainer.whiteNet = NewQNet(b, White)
	trainer.blackNet = NewQNet(b, Black)
}

// Plays one game and accumulates the playout
func (trainer *QTrainer) PlayOneGame(debug bool) {
	playout := NewQPlayout(trainer.whiteNet, trainer.blackNet)
	if trainer.playouts == nil {
		trainer.playouts = []*QPlayout{}
	}
	trainer.playouts = append(trainer.playouts, playout)
}

const DefaultBatchSize int = 100

// Plays a batch, til we have batchSize games.
// This will complete any batch in progress.
func (trainer *QTrainer) PlayBatch(batchSize int) {
	for len(trainer.playouts) < batchSize {
		trainer.PlayOneGame(false)
	}
}

// Learns from a batch and resets for the next one.
func (trainer *QTrainer) LearnFromBatch() {
	trainer.whiteNet.LearnFromPlayouts(trainer.playouts, 0.1)
	trainer.blackNet.LearnFromPlayouts(trainer.playouts, 0.1)
	trainer.batches++
}

func (trainer *QTrainer) Play(b Board) (NaiveSpot, float64) {
	board := b.ToTopoBoard()
	trainer.init(board)

	if !Debug {
		start := time.Now()
		for SecondsSince(start) < trainer.Seconds {
			trainer.PlayBatch(DefaultBatchSize)
			trainer.LearnFromBatch()
		}
	} else {

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
				// Print what white is thinking
				trainer.whiteNet.Debug()
			case "1":
				trainer.PlayOneGame(true)
			case "p":
				trainer.PlayBatch(DefaultBatchSize)
			case "x":
				// finish
				keepPlaying = false

			default:
				log.Printf("unrecognized command")
			}
		}
	}

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

	log.Printf("played %d batches", trainer.batches)

	winRate := float64(winCount[bestMove]) / float64(moveCount[bestMove])
	return bestMove.NaiveSpot(), winRate
}
