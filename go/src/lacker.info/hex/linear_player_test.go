package hex

import (
	"log"
	"testing"
)

func TestLinearPlayer(t *testing.T) {
	board := NewTopoBoard()
	black := MakeLinearPlayer(board, Black)
	white := MakeLinearPlayer(board, White)
	ending := Playout(black, white, false)
	if ending.Winner != Black {
		log.Fatal("expected Black to win default game among linear players")
	}
}
