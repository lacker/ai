package hex

import (
	"log"
	"testing"
)

func TestDeltaNetBasicOperation(t *testing.T) {
	board := NewTopoBoard()
	whitePlayer := NewDeltaNet(board, White)
	blackPlayer := NewDeltaNet(board, Black)
	Playout(whitePlayer, blackPlayer, false)
}

func TestDeltaNetOverrideSpot(t *testing.T) {
	board := NewTopoBoard()
	blackPlayer := NewDeltaNet(board, Black)
	spot, _ := blackPlayer.BestMove(board, false)
	if spot != TopLeftCorner {
		log.Fatal("expected TopLeftCorner to win by default")
	}
	blackPlayer.overrideSpot = BottomRightCorner
	spot, _ = blackPlayer.BestMove(board, false)
	if spot != BottomRightCorner {
		log.Fatal("expected BottomRightCorner to win by override")
	}
}
