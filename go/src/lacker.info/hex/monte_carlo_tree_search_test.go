package hex

import (
	"math"
	"testing"
)

func TestSimpleChain(t *testing.T) {
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

	if root.SelectLeaf() != root {
		t.Fatalf("the root should also be a leaf according to SelectLeaf")
	}
}

func TestExpansion(t *testing.T) {
	board := NewBoard()
	root := NewRoot(board)
	for i := 0; i < 121; i++ {
		if root.SelectLeaf() != root {
			t.Fatalf("root.SelectLeaf() should be root at iteration %d", i)
		}
		depth := root.Depth()
		if i > 0 && depth != 2 {
			t.Fatalf("on iteration %d got depth %d", i, depth)
		}
		if root.Expand() == nil {
			t.Fatalf("could not expand root on iteration %d", i)
		}
	}
	
	// Finally the root should be full
	leaf := root.SelectLeaf()
	if leaf == root {
		t.Fatalf("leaf should not be root when root is full")
	}
	leaf.Expand()
	if root.Depth() != 3 {
		t.Fatalf("root depth should be three after expanding a child")
	}
}

func TestMCTS(t *testing.T) {
	board := NewBoard()
	root := NewRoot(board)
	for i := 0; i < 5; i++ {
		root.RunOneRoundOfMCTS()
	}
	if root.BlackWins + root.WhiteWins != 5 {
		t.Fatalf("five mcts loops should lead to 5 win counts in the root")
	}
}
