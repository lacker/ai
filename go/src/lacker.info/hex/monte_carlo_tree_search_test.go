package hex

import (
	"math"
	"testing"
)

func TestUCT(t *testing.T) {
	board := NewBoard()
	root := NewRoot(board)
	if root.UCT() != math.Inf(1) {
		t.Fatalf("root.UCT() was not Inf")
	}

	middle := NewChild(root, Spot{1, 1})
	if middle.Board == nil {
		t.Fatalf("middle should have a non-nil board")
	}
	if middle.Board.Get(Spot{1, 1}) != Black {
		t.Fatalf("middle should have a black stone at 1,1")
	}
	t.Log("made middle ok")

	leaf := NewChild(middle, Spot{5, 5})
	t.Log("made leaf ok")

	if leaf.NumPossibleMoves != 119 {
		t.Fatalf("bad num possible moves")
	}
}

