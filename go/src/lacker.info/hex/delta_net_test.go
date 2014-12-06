package hex

import (
	// "log"
	"testing"
)

func TestDeltaNetBasicOperation(t *testing.T) {
	board := NewTopoBoard()
	whitePlayer := NewDeltaNet(board, White)
	blackPlayer := NewDeltaNet(board, Black)
	Playout(whitePlayer, blackPlayer, false)
}
