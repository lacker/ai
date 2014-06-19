package hex

import (
	"testing"
)

func TestBlackWin(t *testing.T) {
	b := NewBoard()
	for r := 0; r < BoardSize; r++ {
		if r != 5 {
			b.Set(Spot{r, 3}, Black)
		}
	}
	if b.IsBlackTheWinner() {
		t.Fatalf("black is not supposed to be the winner because 5, 3 is missing")
	}
	b.Set(Spot{5, 3}, Black)
	if !b.IsBlackTheWinner() {
		t.Fatalf("black is supposed to be the winner because *, 3 is set")
	}
}

func TestWhiteWin(t *testing.T) {
	b := NewBoard()
	for c := 0; c < BoardSize; c++ {
		if c != 8 {
			b.Set(Spot{7, c}, White)
		}
	}
	if b.Winner() != Empty {
		t.Fatalf("expected empty")
	}
	b.Set(Spot{7, 8}, White)
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
