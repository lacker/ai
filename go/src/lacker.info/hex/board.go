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

type NaiveSpot struct {
	Row, Col int
}

func MakeNaiveSpot(row int, col int) NaiveSpot {
	return NaiveSpot{Row: row, Col: col}
}

func AllSpots() [NumSpots]NaiveSpot {
	var answer [NumSpots]NaiveSpot
	for r := 0; r < BoardSize; r++ {
		for c := 0; c < BoardSize; c++ {
			spot := MakeNaiveSpot(r, c)
			answer[spot.Index()] = spot
		}
	}
	return answer
}

func (s NaiveSpot) Index() int {
	return s.Col + BoardSize * s.Row
}

func (s NaiveSpot) String() string {
	return fmt.Sprintf("(%d, %d)", s.Row, s.Col)
}

func (s NaiveSpot) Transpose() NaiveSpot {
	return MakeNaiveSpot(s.Col, s.Row)
}

func (s NaiveSpot) ApplyToNeighbors(f func(NaiveSpot)) {
	if s.Row > 0 {
		f(MakeNaiveSpot(s.Row - 1, s.Col))
	}
	if s.Row + 1 < BoardSize {
		f(MakeNaiveSpot(s.Row + 1, s.Col))
		if s.Col > 0 {
			f(MakeNaiveSpot(s.Row + 1, s.Col - 1))
		}
	}
	if s.Col > 0 {
		f(MakeNaiveSpot(s.Row, s.Col - 1))
	}
	if s.Col + 1 < BoardSize {
		f(MakeNaiveSpot(s.Row, s.Col + 1))
		if s.Row > 0 {
			f(MakeNaiveSpot(s.Row - 1, s.Col + 1))
		}
	}
}

func (s NaiveSpot) Neighbors() []NaiveSpot {
	answer := make([]NaiveSpot, 0)
	possible := []NaiveSpot{
		NaiveSpot{s.Row - 1, s.Col},
		NaiveSpot{s.Row + 1, s.Col},
		NaiveSpot{s.Row, s.Col - 1},
		NaiveSpot{s.Row, s.Col + 1},
		NaiveSpot{s.Row + 1, s.Col - 1},
		NaiveSpot{s.Row - 1, s.Col + 1},
	}
	for _, spot := range possible {
		if spot.Row < 0 || spot.Row >= BoardSize ||
			spot.Col < 0 || spot.Col >= BoardSize {
			continue
		}
		answer = append(answer, spot)
	}
	return answer
}

type Board interface {
	ToNaiveBoard() *NaiveBoard
	ToTopoBoard() *TopoBoard
	Copy() Board
	PossibleMoves() []NaiveSpot
	MakeMoveWithNaiveSpot(s NaiveSpot)
	MakeMove(s TopoSpot)
	GetToMove() Color
	Get(s NaiveSpot) Color

	// Returns a list of spots to count that contributed towards the
	// winner winning.
	GetWinningPathSpots() []NaiveSpot

	// Plays out the game randomly and tells you who won.
	Playout() Color
}
