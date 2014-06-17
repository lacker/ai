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
}

