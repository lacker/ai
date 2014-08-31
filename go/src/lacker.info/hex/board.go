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

func PrintInfo() {
	fmt.Printf("Playing hex on a size-%d board.\n", BoardSize)
}

type Spot struct {
	Row, Col int
}

func MakeSpot(row int, col int) Spot {
	return Spot{Row: row, Col: col}
}

func AllSpots() [NumSpots]Spot {
	var answer [NumSpots]Spot
	for r := 0; r < BoardSize; r++ {
		for c := 0; c < BoardSize; c++ {
			spot := MakeSpot(r, c)
			answer[spot.Index()] = spot
		}
	}
	return answer
}

func (s Spot) Index() int {
	return s.Col + BoardSize * s.Row
}

func (s Spot) String() string {
	return fmt.Sprintf("(%d, %d)", s.Row, s.Col)
}

func (s Spot) Transpose() Spot {
	return MakeSpot(s.Col, s.Row)
}

func (s Spot) ApplyToNeighbors(f func(Spot)) {
	if s.Row > 0 {
		f(MakeSpot(s.Row - 1, s.Col))
	}
	if s.Row + 1 < BoardSize {
		f(MakeSpot(s.Row + 1, s.Col))
		if s.Col > 0 {
			f(MakeSpot(s.Row + 1, s.Col - 1))
		}
	}
	if s.Col > 0 {
		f(MakeSpot(s.Row, s.Col - 1))
	}
	if s.Col + 1 < BoardSize {
		f(MakeSpot(s.Row, s.Col + 1))
		if s.Row > 0 {
			f(MakeSpot(s.Row - 1, s.Col + 1))
		}
	}
}

func (s Spot) Neighbors() []Spot {
	answer := make([]Spot, 0)
	possible := []Spot{
		Spot{s.Row - 1, s.Col},
		Spot{s.Row + 1, s.Col},
		Spot{s.Row, s.Col - 1},
		Spot{s.Row, s.Col + 1},
		Spot{s.Row + 1, s.Col - 1},
		Spot{s.Row - 1, s.Col + 1},
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
}
