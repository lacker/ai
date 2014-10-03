package hex

import (
	"testing"
)

func TestLinearPlayer(t *testing.T) {
	board := NewTopoBoard()
	black := MakeLinearPlayer(board, Black)
	white := MakeLinearPlayer(board, White)
	Playout(black, white, false)
}
