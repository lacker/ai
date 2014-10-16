package hex

import (
	"log"
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

	snipList, ending := FindWinningSnipList(white, black, mainLine, false)
	if len(snipList) != 1 || snipList[0].String() != "1 => (10, 0)" {
		log.Fatal("unexpected snip list")
	}
	if ending.Winner != White {
		log.Fatal("expected White to win with the new snip list")
	}
}
