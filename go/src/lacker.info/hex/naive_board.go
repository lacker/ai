package hex

import (
	"encoding/json"
	"log"
	"strings"
)

/*
In the NaiveBoard, to represent a spot, we do row and column like it's
a matrix.
*/

type NaiveBoard struct {
	// Contents of the board
	// indices are Row, Col
	Board [BoardSize][BoardSize]Color

	// Whose move it is
	ToMove Color
}

func NewNaiveBoard() *NaiveBoard {
	return &NaiveBoard{ToMove: Black}
}

func (b *NaiveBoard) Get(spot Spot) Color {
	return b.Board[spot.Row][spot.Col]
}

func (b *NaiveBoard) Set(spot Spot, color Color) {
	b.Board[spot.Row][spot.Col] = color
}

func (b *NaiveBoard) Eprint() {
	Eprint("Board:\n")
	for r, col := range b.Board {
		Eprint(strings.Repeat(" ", r))
		for c, color := range col {
			if c > 0 {
				Eprint(" ")
			}
			switch color {
			case Black:
				Eprint("B")
			case White:
				Eprint("w")
			case Empty:
				Eprint(".")
			}
		}
		Eprint("\n")
	}
}

func (b *NaiveBoard) PossibleMoves() []Spot {
	answer := make([]Spot, 0)
	for r, col := range b.Board {
		for c, color := range col {
			if color == Empty {
				answer = append(answer, MakeSpot(r, c))
			}
		}
	}
	return answer
}

func (b *NaiveBoard) MakeMoveWithNaiveSpot(s Spot) {
	if b.ToMove == Empty {
		log.Fatal("this isn't a valid board, there is nobody to move")
	}
	if b.Get(s) != Empty {
		log.Fatal("cannot move on a non empty spot")
	}
	b.Set(s, b.ToMove)
	b.ToMove = -b.ToMove
}

func (b *NaiveBoard) MakeMove(s TopoSpot) {
	b.MakeMoveWithNaiveSpot(s.ToSpot())
}

func (b *NaiveBoard) GetToMove() Color {
	return b.ToMove
}

func (b *NaiveBoard) Transpose() *NaiveBoard {
	t := NewNaiveBoard()
	t.ToMove = -b.ToMove
	for _, spot := range AllSpots() {
		t.Set(spot, -b.Get(spot.Transpose()))
	}
	return t
}

func (b *NaiveBoard) ToNaiveBoard() *NaiveBoard {
	c := NewNaiveBoard()
	c.ToMove = b.ToMove
	for _, spot := range AllSpots() {
		c.Set(spot, b.Get(spot))
	}
	return c
}

func (b *NaiveBoard) ToTopoBoard() *TopoBoard {
	c := NewTopoBoard()
	c.ToMove = b.ToMove
	for _, spot := range AllSpots() {
		color := b.Get(spot)
		if color != Empty {
			c.Set(spot.Row, spot.Col, color)
		}
	}
	return c
}

func (b *NaiveBoard) Copy() Board {
	return b.ToNaiveBoard()
}

// Makes moves repeatedly. When this stops the game is over.
// Returns the winner.
// This mutates the board.
func (b *NaiveBoard) Playout() Color {
	moves := b.PossibleMoves()
	ShuffleSpots(moves)

	for _, move := range moves {
		b.MakeMoveWithNaiveSpot(move)
	}

	winner := b.Winner()
	if winner == Empty {
		log.Fatal("no winner in a playout")
	}

	return winner
}

// Black wins if you can get from row 0 to row BoardSize - 1 with just
// black spots.
func (b *NaiveBoard) IsBlackTheWinner() bool {
	// Frontier is black stones we haven't investigated yet.
	// Checked is any previously-frontier stone we already processed.
	// Start off with the frontier of all the row-zero black stones.
	frontier := make([]Spot, 0)
	var checked [NumSpots]bool
	for col, color := range(b.Board[0]) {
		if color == Black {
			frontier = append(frontier, MakeSpot(0, col))
		}
	}

	// The search loop continuously processes the frontier
	for len(frontier) > 0 {
		spot := frontier[0]
		frontier = frontier[1:]
		checked[spot.Index()] = true

		// Find all the neighboring black stones
		done := false
		spot.ApplyToNeighbors(func(neighbor Spot) {
			if b.Get(neighbor) != Black {
				return
			}
			if checked[neighbor.Index()] {
				return
			}
			// fmt.Printf("processing %d, %d\n", neighbor.Row, neighbor.Col)
			if neighbor.Row == BoardSize - 1 {
				done = true
				return
			}
			frontier = append(frontier, neighbor)
		})
		if done {
			return true
		}
	}
	
	return false
}

func (b *NaiveBoard) Winner() Color {
	if b.IsBlackTheWinner() {
		return Black
	}
	if b.Transpose().IsBlackTheWinner() {
		return White
	}
	return Empty
}

func NewNaiveBoardFromJSON(j string) *NaiveBoard {
	b := new(NaiveBoard)
	err := json.Unmarshal([]byte(j[:]), &b)
	if err != nil {
		log.Fatal("NewNaiveBoardFromJSON failed: ", err)
	}
	return b
}

func (b *NaiveBoard) GetWinningPathSpots() []Spot {
	panic("not implemented")
}
