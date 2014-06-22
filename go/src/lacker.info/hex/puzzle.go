package hex

import (
	"log"
	"strings"
)

type Puzzle struct {
	Board *Board
	CorrectAnswer Spot
}

// The format is, the first three words are
// "x to move" where x is Black or White
// After that the non-white-space entries are B, ., or W
// The move you are supposed to make is a *
func MakePuzzle(s string) Puzzle {
	puzzle := Puzzle{}
	puzzle.Board = new(Board)
	words := strings.Fields(s)
	if len(words) != 124 {
		log.Fatal("cannot make puzzle from %d words", len(words))
	}

	switch words[0] {
	case "Black":
		puzzle.Board.ToMove = Black
	case "White":
		puzzle.Board.ToMove = White
	default:
		log.Fatal("bad player name: %s", words[0])
	}

	for index, spot := range AllSpots() {
		word := words[index + 3]
		switch word {
		case "B":
			puzzle.Board.Set(spot, Black)
		case "W":
			puzzle.Board.Set(spot, White)
		case "*":
			puzzle.CorrectAnswer = spot
		}
	}

	return puzzle
}
