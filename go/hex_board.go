package main

import (
	"fmt"
)

const BoardSize = 11;

type Color int8
const Black Color = -1;
const White Color = 1;
const Empty Color = 0;

/*
The board is a grid. Each value is either BLACK, WHITE, or EMPTY.
In the external world, spots are typically referred to with a
letter plus a number.
Display would look like a diamond, e.g. for a 4x4 board:

ABCD
-----
\    \       1
 \    \      2
  \    \     3
   \    \    4
    -----

To represent a spot, we do row and column like it's a matrix.
Black goes top to bottom; White goes left to right.
So Black could win with a single column; White could win with a single row.
*/

type Spot struct {
	Row, Col int
}

type Board struct {
	// Contents of the board
	// indices are Row, Col
	Board [BoardSize][BoardSize]Color

	// Whose move it is
	ToMove Color
}

func NewBoard() Board {
	return Board{ToMove: Black}
}

func (b *Board) PossibleMoves() []Spot {
	answer := make([]Spot, 0);
	for r, col := range b.Board {
		for c, color := range col {
			if color == Empty {
				answer = append(answer, Spot{Row:r, Col:c})
			}
		}
	}
	return answer
}

func main() {
	fmt.Printf("sup\n");
}
