package hex

import (
	"log"
	"testing"
)

func TestLinearPlayerPlayout(t *testing.T) {
	board := NewTopoBoard()
	black := NewLinearPlayer(board, Black)
	white := NewLinearPlayer(board, White)
	ending := Playout(black, white, false)
	if ending.Winner != Black {
		log.Fatal("expected Black to win default game among linear players")
	}
}

func TestLinearPlayerBestMove(t *testing.T) {
	board := NewTopoBoard()
	black := NewLinearPlayer(board, Black)
	spot := black.BestMove(board)
	if spot.Row() != 0 || spot.Col() != 0 {
		log.Fatal("expected (0, 0) to be the best move")
	}
}
