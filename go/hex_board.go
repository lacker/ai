package main

import (
	"fmt"
)

const BoardSize = 11
const NumSpots = BoardSize * BoardSize

type Color int8
const Black Color = -1
const White Color = 1
const Empty Color = 0

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

func AllSpots() [NumSpots]Spot {
	var answer [NumSpots]Spot
	for r := 0; r < BoardSize; r++ {
		for c := 0; c < BoardSize; c++ {
			answer[r * BoardSize + c] = Spot{r, c};
		}
	}
	return answer
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
				answer = append(answer, Spot{r, c})
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

func (b *Board) Transpose() Board {
	t := NewBoard()
	t.ToMove = -b.ToMove
	for _, spot := range AllSpots() {
		t.Board[spot.Row][spot.Col] = -b.Board[spot.Col][spot.Row]
	}
	return t
}

func main() {
	fmt.Printf("sup\n");
}
