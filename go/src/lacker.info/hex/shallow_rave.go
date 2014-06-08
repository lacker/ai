package hex

/*
The shallow rave algorithm is that you do playouts from the given
position, and whichever move is in the most winning results, you
pick that one.
*/

type ShallowRave struct {
	NumPlayouts int
}

func (s *ShallowRave) Play(b *Board) Spot {
	panic("don't know how to play")
}

