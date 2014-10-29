package hex

import (
	"fmt"
)

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

Black goes top to bottom; White goes left to right.
So Black could win with a single column; White could win with a single row.
*/

/*
Board is an interface which both NaiveBoard and TopoBoard implement.
*/

// If we raise BoardSize we also need to change the type of TopoSpot
const BoardSize = 11
const NumSpots = BoardSize * BoardSize

type Color int8
const Black Color = -1
const White Color = 1
const Empty Color = 0

var Debug bool = false

func (c Color) Name() string {
	switch c {
	case Black:
		return "Black"
	case White:
		return "White"
	case Empty:
		return "Empty"
	}
	panic("bad color")
}

func PrintInfo() {
	fmt.Printf("Playing hex on a size-%d board.\n", BoardSize)
}

type Board interface {
	ToNaiveBoard() *NaiveBoard
	ToTopoBoard() *TopoBoard
	Copy() Board
	PossibleMoves() []NaiveSpot
	MakeMove(s Spot)
	GetToMove() Color
	Get(s Spot) Color

	// Returns a list of spots to count that contributed towards the
	// winner winning.
	GetWinningPathSpots() []NaiveSpot

	// Plays out the game randomly and tells you who won.
	Playout() Color
}
