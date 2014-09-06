package hex

import (
	"log"
	"strings"
)

type Puzzle struct {
	String string
	Board *NaiveBoard
	CorrectAnswer Spot
}

var PuzzleMap map[string]Puzzle = MakePuzzleMap()

// Create the library of interesting puzzles.
func MakePuzzleMap() map[string]Puzzle {
	puzzleMap := make(map[string]Puzzle)

	// Any reasonable method should be able to find a killer move.
	puzzleMap["onePly"] = MakePuzzle(`
Black to move
B . . . . . . . . . .
 B . . . . . . . . . .
  B . . . . . . . . . .
   B . . . . . . . . . .
    B . . . . . . . . . .
     B . . . . . . . . . .
      B . . . . . . . . . .
       B . . . . . . . . . .
        B . . . . . . . . . .
         B . . . . . . . . . .
          * W W W W W W W W W W
`)

	// Tree methods can figure out a block where shallow rave can't,
	// because they can figure out the bridges.
	// MCTS can figure this out consistently in 0.2s, but not in 0.1s.
	// Ideally this would be fast enough for a playouter to get it.
	puzzleMap["triangleBlock"] = MakePuzzle(`
Black to move
B . . . . . . . . . .
 B . . . . . . . . . .
  B . . . . . . . . . .
   B . . . . . . . . . .
    B . . . . . . . . . .
     B B . . . . . . . . .
      . . W W W W W W W W W
       * . . . . . . . . . .
        . B . . . . . . . . .
         B . . . . . . . . . .
          B . . . . . . . . . .
`)

	// Tree methods still cannot understand a large amount of bridges.
	// MCTS with 0.2s can occasionally pass this but usually can't.
	puzzleMap["manyBridges"] = MakePuzzle(`
Black to move
. . . . . . . . . . .
 . . B . . . . . . . .
  . . . . . . . . . . .
   . B . . . B B B B . .
    . . . B . . W . B B .
     B B . . W . . W . B B
      . . W . B B B . W . .
       * . W B . . B B . W .
        . B W . . . . B B B .
         . . . W . . W . . W .
          B . . . W . . W . . .
`)

	return puzzleMap
}

func GetPuzzle(name string) Puzzle {
	answer, ok := PuzzleMap[name]
	if !ok {
		log.Fatal("no puzzle with name: %s", name)
	}
	return answer
}

// The format is, the first three words are
// "x to move" where x is Black or White
// After that the non-white-space entries are B, ., or W
// The move you are supposed to make is a *
func MakePuzzle(s string) Puzzle {
	puzzle := Puzzle{String: s, Board: new(NaiveBoard)}
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

func (puzzle Puzzle) Test(player Player) bool {
	playerAnswer, _ := player.Play(puzzle.Board)
	if puzzle.CorrectAnswer != playerAnswer {
		log.Printf(puzzle.String)
		log.Printf("got wrong answer: %s", playerAnswer)
		return false
	}
	return true
}
