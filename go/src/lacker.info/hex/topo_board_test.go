package hex

import (
	"fmt"
	"testing"
)

func TestTopoBoardNumGroups(t *testing.T) {
	b := NewTopoBoard()

	if b.NumGroups() != 4 {
		t.Fatalf("board should start with four groups")
	}
	b.Set(0, 3, Black)
	if b.NumGroups() != 4 {
		t.Fatalf("board should still have four groups after one stone")
	}
	b.Set(1, 3, Black)
	if b.NumGroups() != 4 {
		t.Fatalf("board should still have four groups after two stones")
	}
}

func TestTopoBoardBlackWin(t *testing.T) {
	b := NewTopoBoard()
	for r := 0; r < BoardSize; r++ {
		if r != 5 {
			b.Set(r, 3, Black)
		}
	}
	if b.Winner == Black {
		t.Fatalf("black is not supposed to be the winner because 5, 3 is missing")
	}
	b.Set(5, 3, Black)

	if b.Winner != Black {
		fmt.Printf("b.GroupSpots: %v\n", b.GroupSpots)
		t.Fatalf("black is supposed to be the winner because *, 3 is set")
	}
}
