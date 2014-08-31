package hex

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestTopoBoardImplementsBoard(t *testing.T) {
	var b Board
	tb := NewTopoBoard()
	b = tb
	b.ToNaiveBoard()
}

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

func TestTopoBoardWhiteWin(t *testing.T) {
	b := NewTopoBoard()
	for c := 0; c < BoardSize; c++ {
		if c != 8 {
			b.Set(7, c, White)
		}
	}
	if b.Winner != Empty {
		t.Fatalf("expected empty")
	}
	b.Set(7, 8, White)
	if b.Winner != White {
		t.Fatalf("expected white")
	}
}

func TestTopoBoardPlayout(t *testing.T) {
	for i := 0; i < 10; i++ {
		b := NewTopoBoard()
		b.Playout()
	}
}

func BenchmarkTopoBoardPlayout(b *testing.B) {
	rand.Seed(1)

	for i := 0; i < b.N; i++ {
		board := NewTopoBoard()
		board.Playout()
	}
}
