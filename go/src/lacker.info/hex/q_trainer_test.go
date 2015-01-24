package hex

import (
	"math/rand"
	"testing"
)

func TestQTrainerOnDoomed1(t *testing.T) {
	rand.Seed(1)
	board := PuzzleMap["doomed1"].Board.ToTopoBoard()
	qt := &QTrainer{Seconds:-1, Quiet:true}
	qt.init(board)
	qt.PlayOneGame(false)
	qt.PlayOneGame(false)
}
