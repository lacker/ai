package hex

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestBoardBlackWin(t *testing.T) {
	b := NewNaiveBoard()
	for r := 0; r < BoardSize; r++ {
		if r != 5 {
			b.Set(MakeSpot(r, 3), Black)
		}
	}
	if b.IsBlackTheWinner() {
		t.Fatalf("black is not supposed to be the winner because 5, 3 is missing")
	}
	b.Set(MakeSpot(5, 3), Black)
	if !b.IsBlackTheWinner() {
		t.Fatalf("black is supposed to be the winner because *, 3 is set")
	}
}

func TestBoardWhiteWin(t *testing.T) {
	b := NewNaiveBoard()
	for c := 0; c < BoardSize; c++ {
		if c != 8 {
			b.Set(MakeSpot(7, c), White)
		}
	}
	if b.Winner() != Empty {
		t.Fatalf("expected empty")
	}
	b.Set(MakeSpot(7, 8), White)
	if b.Winner() != White {
		t.Fatalf("expected white")
	}
	encoded := ToJSON(b)
	b2 := NewNaiveBoardFromJSON(encoded)
	if b2.Winner() != White {
		t.Fatalf("something wacky happened with encoding")
	}
}

func TestBoardPlayout(t *testing.T) {
	for i := 0; i < 10; i++ {
		b := NewNaiveBoard()
		b.Playout()
	}
}

func TestBoardStringification(t *testing.T) {
	s := MakeSpot(2, 3)
	if fmt.Sprintf("%s", s) != "(2, 3)" {
		t.Fatalf("problems printf'ing %s", s)
	}
}

func BenchmarkBoardPlayout(b *testing.B) {
	rand.Seed(1)

	for i := 0; i < b.N; i++ {
		board := NewNaiveBoard()
		board.Playout()
	}
}
