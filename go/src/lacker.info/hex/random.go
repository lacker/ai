package hex

import (
	"math/rand"
)

/*
The random player plays a random legal move.
*/

type Random struct {
}

func (r Random) Play(b *Board) Spot {
	moves := b.PossibleMoves()
	return moves[rand.Intn(len(moves))]
}
