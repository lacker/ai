package hex

import (
	"log"
	"testing"
)

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

func TestDeltaNetFindNewMainLine(t *testing.T) {
	board := NewTopoBoard()
	whitePlayer := NewDeltaNet(board, White)
	blackPlayer := NewDeltaNet(board, Black)

	// black goes first
	ending := Playout(blackPlayer, whitePlayer, false)

	if ending.Winner != Black {
		log.Fatal("Black is supposed to win")
	}

	_, newMainLine := whitePlayer.FindNewMainLine(blackPlayer,
		ending, false)
	if newMainLine == nil {
		log.Fatal("newMainLine should not be nil")
	}
}
