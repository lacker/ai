package hex

import (
	"encoding/json"
	"fmt"
	"log"
)

const BoardSize = 11
const NumSpots = BoardSize * BoardSize

type Color int8
const Black Color = -1
const White Color = 1
const Empty Color = 0

func PrintInfo() {
	fmt.Printf("Playing hex on a size-%d board.\n", BoardSize)
}

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

func (s *Spot) Transpose() Spot {
	return Spot{Row:s.Col, Col:s.Row}
}

func (s *Spot) Neighbors() []Spot {
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

type Board struct {
	// Contents of the board
	// indices are Row, Col
	Board [BoardSize][BoardSize]Color

	// Whose move it is
	ToMove Color
}

func NewBoard() *Board {
	return &Board{ToMove: Black}
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
		log.Fatal("this isn't a valid board, there is nobody to move")
	}
	if b.Get(s) != Empty {
		return false
	}
	b.Set(s, b.ToMove)
	b.ToMove = -b.ToMove
	return true
}

func (b *Board) Transpose() *Board {
	t := NewBoard()
	t.ToMove = -b.ToMove
	for _, spot := range AllSpots() {
		t.Set(spot, -b.Get(spot.Transpose()))
	}
	return t
}

// Black wins if you can get from row 0 to row BoardSize - 1 with just
// black spots.
func (b *Board) IsBlackTheWinner() bool {
	// Frontier is black stones we haven't investigated yet.
	// Checked is any previously-frontier stone we already processed.
	// Start off with the frontier of all the row-zero black stones.
	frontier := make([]Spot, 0)
	checked := make(map[Spot]bool)
	for col, color := range(b.Board[0]) {
		if color == Black {
			frontier = append(frontier, Spot{Row:0, Col:col});
		}
	}

	// The search loop continuously processes the frontier
	for len(frontier) > 0 {
		spot := frontier[0]
		frontier = frontier[1:]
		checked[spot] = true

		// Find all the neighboring black stones
		for _, neighbor := range spot.Neighbors() {
			if b.Get(neighbor) != Black {
				continue
			}
			if checked[neighbor] {
				continue
			}
			// fmt.Printf("processing %d, %d\n", neighbor.Row, neighbor.Col)
			if neighbor.Row == BoardSize - 1 {
				return true
			}
			frontier = append(frontier, neighbor)
		}
	}
	
	return false
}

func (b *Board) winner() Color {
	if b.IsBlackTheWinner() {
		return Black
	}
	if b.Transpose().IsBlackTheWinner() {
		return White
	}
	return Empty
}

func NewBoardFromJSON(j string) *Board {
	b := new(Board)
	err := json.Unmarshal([]byte(j[:]), &b)
	if err != nil {
		log.Fatal("NewBoardFromJSON failed: ", err)
	}
	return b
}


