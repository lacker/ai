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
	return b.Board[spot.Row][spot.Col]
}

func (b *Board) Set(spot Spot, color Color) {
	b.Board[spot.Row][spot.Col] = color
}

func (b *Board) Eprint() {
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

func (b *Board) PossibleMoves() []Spot {
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

func (b *Board) Copy() *Board {
	c := NewBoard()
	c.ToMove = b.ToMove
	for _, spot := range AllSpots() {
		c.Set(spot, b.Get(spot))
	}
	return c
}

// Makes moves repeatedly. When this stops the game is over.
// Returns the winner.
// This mutates the board.
func (b *Board) Playout() Color {
	moves := b.PossibleMoves()
	ShuffleSpots(moves)

	for _, move := range moves {
		if !b.MakeMove(move) {
			log.Fatal("a playout played an invalid move")
		}
	}

	winner := b.Winner()
	if winner == Empty {
		log.Fatal("no winner in a playout")
	}

	return winner
}

// Black wins if you can get from row 0 to row BoardSize - 1 with just
// black spots.
func (b *Board) IsBlackTheWinner() bool {
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

func (b *Board) Winner() Color {
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


