package hex

import (
	"testing"
)

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
		t.Fatalf("black is supposed to be the winner because *, 3 is set")
	}
}
