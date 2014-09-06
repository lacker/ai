package hex

import (
	"math/rand"
)

/*
The random player plays a random legal move.
*/

type Random struct {
}

func (r Random) Play(b Board) (Spot, float64) {
	moves := b.PossibleMoves()
	return moves[rand.Intn(len(moves))], 0.5
}
