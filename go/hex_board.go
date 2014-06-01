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

func (b *Board) Get(spot Spot) Color {
	return b.Board[spot.Row][spot.Col];
}

func (b *Board) Set(spot Spot, color Color) {
	b.Board[spot.Row][spot.Col] = color;
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

// Returns whether it was a possible move
func (b *Board) MakeMove(s Spot) bool {
	if b.ToMove == Empty {
		panic("this isn't a valid board, there is nobody to move")
	}
	if b.Get(s) != Empty {
		return false
	}
	b.Set(s, b.ToMove)
	b.ToMove = -b.ToMove
	return true
}

func main() {
	fmt.Printf("sup\n");
}
