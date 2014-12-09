package hex

import (
	"log"
	"sort"
	"testing"
)

func TestFindWinningSnipList(t *testing.T) {
	board := NewTopoBoard()
	black := NewLinearPlayer(board, Black)
	white := NewLinearPlayer(board, White)

	mainLine := Playout(black, white, false)

	// The default linear players should play a game that turns into
	// vertical lines.
	// Thus column zero is black and black wins.
	if mainLine.Winner != Black {
		log.Fatal("expected Black to win in game between default players")
	}
	if mainLine.GetByRowCol(10, 0) != Black {
		log.Fatal("expected Black to have spot 10, 0")
	}
	if mainLine.GetByRowCol(10, 1) != Empty {
		log.Fatal("expected 10, 1 to be empty")
	}

	snipList, ending := FindWinningSnipList(white, black, mainLine, 0, false)
	if len(snipList) != 1 || snipList[0].String() != "1 => (10, 0)" {
		log.Fatal("unexpected snip list")
	}
	if ending.Winner != White {
		log.Fatal("expected White to win with the new snip list")
	}
}

func TestSolvingDoubleBridgeViaSnipList(t *testing.T) {
	// This test is supposed to ensure that a double-bridge that a
	// linear player is attempting to block one bridge at a time can be
	// solved via snip list.
	board := PuzzleMap["doomed2"].Board.ToTopoBoard()
	
	// White is a semismart defender that just happens to prefer (5, 6)
	// and (6, 6) to (6, 2) and (7, 2).
	white := NewLinearPlayer(board, White)
	white.SetScore(5, 6, 400)
	white.SetScore(6, 6, 300)
	white.SetScore(6, 2, 200)
	white.SetScore(7, 2, 100)
	sort.Stable(white.ranking)

	// The only way to defeat white is by running through 6,2 and 7,2.
	// We should find that in a snip list.
	black := NewLinearPlayer(board, Black)
	mainLine := Playout(black, white, false)
	if mainLine.Winner != White {
		log.Fatal("expected White to defend the bridges")
	}
	
	snipList, _ := FindWinningSnipList(black, white, mainLine, 0, false)
	if len(snipList) != 2 {
		log.Fatal("expected two snips for a bridge")
	}
	if snipList[0].String() != "0 => (6, 2)" {
		log.Fatal("expected snipList[0] to be 0 => (6, 2)")
	}
	if snipList[1].String() != "2 => (7, 2)" {
		log.Fatal("expected snipList[1] to be 2 => (7, 2)")
	}
}
