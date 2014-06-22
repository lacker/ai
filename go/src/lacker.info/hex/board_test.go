package hex

import (
	"math/rand"
	"fmt"
	"testing"
)

func TestBlackWin(t *testing.T) {
	b := NewBoard()
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

func TestWhiteWin(t *testing.T) {
	b := NewBoard()
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
	b2 := NewBoardFromJSON(encoded)
	if b2.Winner() != White {
		t.Fatalf("something wacky happened with encoding")
	}
}

func TestPlayout(t *testing.T) {
	for i := 0; i < 10; i++ {
		b := NewBoard()
		b.Playout()
	}
}

func TestStringification(t *testing.T) {
	s := MakeSpot(2, 3)
	if fmt.Sprintf("%s", s) != "(2, 3)" {
		t.Fatalf("problems printf'ing %s", s)
	}
}

func BenchmarkPlayout(b *testing.B) {
	rand.Seed(1)

	for i := 0; i < b.N; i++ {
		board := NewBoard()
		board.Playout()
	}
}
