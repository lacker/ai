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

	// This ladder should be a win for the first player.
	// It requires looking 21 plies deep though.
	puzzleMap["ladder"] = MakePuzzle(`
Black to move
B . . . . . . . . . .
 B . . . . . . . . . .
  B . . . . . . . . . .
   B . . . . . . . . . .
    B . . . . . . . . . .
     B . . . . . . . . . .
      B . . . . . . . . . .
       B . . . . . . . . . .
        B W W W W W W W W W W
         * . . . . . . . . . .
          . . . . . . . . . . B
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

	// The right move should be obvious - there's one spot that will
	// make a critical connection for either side.
	//
	// "Thread the needle."
	//
	// The interesting thing with spot sorter is that adding a single
	// black stone to connect the bottom group to the bottom makes it
	// obvious what the correct answer is. So, the spot sorter is indeed
	// getting confused by an open bridge.
	//
	// The puzzle is reasonable for either black or white to move. With
	// black to move it should also be obvious that it's close to 100%
	// winning after moving (5, 6).

	puzzleMap["needle"] = MakePuzzle(`
White to move
. . . . . . . . . W B
 . B . . . . . . W B .
  . . . . . . . W B . .
   . . . . . . W B . . .
    . . . . . W B . . . .
     . . . . . B * W . . .
      . . . . . W B . . . .
       . . . . W B W . . . .
        . . . . B B B . . . .
         . . W B W W W . . . .
          . . . . . . . . . . .
`)

	// In the "doomedX" series, the player to move should realize that
	// we are doomed.
	puzzleMap["doomed1"] = MakePuzzle(`
Black to move
. . . . . . . . . . B
 . B . . . . . . . B .
  . . . . . . . . B . W
   . . . . . . . B . W .
    . . . . . . B . W . .
     . . . . . . . W . . .
      W W W W W W . . . . .
       . . . . . B . . . . .
        . . . . B . . . . . .
         . . . B . . . . . . .
          . . B . . . . . . . .
`)

	puzzleMap["doomed2"] = MakePuzzle(`
Black to move
. . . . . . . . . . B
 . B . . . . . . . B .
  . . . . . . . . B . W
   . . . . . . . B . W .
    . . . B B B B . W . .
     . . B . . . . W . . .
      . . . W W W . . . . .
       W W . . . B . . . . .
        . . B . B . . . . . .
         . . B B . . . . . . .
          . . B . . . . . . . .
`)

	puzzleMap["doomed3"] = MakePuzzle(`
Black to move
. . . . . . . . . . B
 . B . . . . . . . B .
  . . . . . . . . B . W
   . . . . . . . B . W .
    . . . B B B B . W . .
     . . B . . . . W . . .
      . B . W W W . . . . .
       . W . . . B . . . . .
        . B B . B . . . . . .
         . . B B . . . . . . .
          . . B . . . . . . . .
`)

	// This should be pretty straightforward - there's one obvious move
	// to block. You won't be winning but there's still some chance.
	puzzleMap["simpleBlock"] = MakePuzzle(`
Black to move
. . . . . . . . . . .
 . B . . . . . . . . .
  . . . . . . . . . . .
   . . . . . . . B . . .
    . . . . . B . . B B .
     . . . . . . W W W * .
      W W W W W W B . . . .
       . B B . . . . . . . .
        . . B . . . . . . . .
         . . . . . . . . . . .
          . . . . . . . . . . .
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
		log.Fatalf("cannot make puzzle from %d words", len(words))
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
