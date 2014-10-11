package hex

import (
	"log"
	"testing"
)

func TestDemocracyPlayerPlayout(t *testing.T) {
	board := NewTopoBoard()

	black1 := MakeLinearPlayer(board, Black)
	black2 := MakeLinearPlayer(board, Black)
	black3 := MakeLinearPlayer(board, Black)

	white1 := MakeLinearPlayer(board, White)
	white2 := MakeLinearPlayer(board, White)
	white3 := MakeLinearPlayer(board, White)

	black := MakeDemocracyPlayer(board, Black)
	black.Add(black1)
	black.Add(black2)
	black.Add(black3)
	
	white := MakeDemocracyPlayer(board, White)
	white.Add(white1)
	white.Add(white2)
	white.Add(white3)
	
	ending := Playout(black, white, false)
	if ending.Winner != Black {
		log.Fatal("expected Black to win because Black wins individuals")
	}
}
