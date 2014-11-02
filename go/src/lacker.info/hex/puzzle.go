package hex

import (
	"log"
	"strings"
)

type Puzzle struct {
	String string
	Board *NaiveBoard
	CorrectAnswer NaiveSpot
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

	puzzleMap["doomed4"] = MakePuzzle(`
Black to move
. . . . . . . . . . B
 . B . . . . . . . B .
  . . . . . . . . B . .
   . . . . . . . B . W .
    . . . B B B B . W . .
     . . B . . . . W . . .
      . B . W W W . . . . .
       . W . . . B . . . . .
        . B B . B . . . . . .
         . . B B . . . . . . .
          . . B . . . . . . . .
`)

	puzzleMap["doomed6"] = MakePuzzle(`
Black to move
. . . . . . . . . . B
 . . . . . . . . . . B
  . . . . . . . . . . B
   . . . . . . . . . . B
    . . . . . . . . . . B
     . . . . . . . . . . B
      . . . . . . . . . . B
       . . . . . . . . . . B
        B B B B B B B B B B B
         . W . W . W . W . W .
          . W . W . W . W . W .
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
		log.Fatalf("no puzzle with name: %s", name)
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

type puzzleScorer struct {
	playerName string
	right int
	wrong int
	total int
}

func (s *puzzleScorer) score(success bool) bool {
	if success {
		s.right++
	} else {
		s.wrong++
	}
	s.total++
	return success
}

type PuzzleType int
const (
	DefiniteWin PuzzleType = iota
	DefiniteLoss
	ClearMove
) 

func (s *puzzleScorer) solve(puzzleName string, ptype PuzzleType) {
	player := GetPlayer(s.playerName)
	puzzle := GetPuzzle(puzzleName)
	playerAnswer, conf := player.Play(puzzle.Board)

	if ptype == DefiniteWin || ptype == ClearMove {
		// Score the move
		if s.score(puzzle.CorrectAnswer == playerAnswer) {
			log.Printf("%s: move OK", puzzleName)
		} else {
			log.Printf("%s:%s", puzzleName, puzzle.String)
			log.Printf("got wrong answer: %s", playerAnswer)
			if ptype == DefiniteWin {
				log.Printf("wrong answer means confidence doesn't matter")
				s.score(false)
				return
			}
		}
	}

	if ptype == DefiniteWin {
		// Score the confidence
		if s.score(conf > 0.999) {
			log.Printf("%s: conf OK", puzzleName)
		} else {
			log.Printf(puzzle.String)
			log.Printf("%s: confidence is only %.2f", puzzleName, conf)
		}
	}

	if ptype == DefiniteLoss {
		// Score the confidence
		if s.score(conf < 0.001) {
			log.Printf("%s: conf OK", puzzleName)
		} else {
			log.Printf(puzzle.String)
			log.Printf("%s: confidence is unwarranted at %.2f", puzzleName, conf)
		}
	}
}

// Runs the player through a series of puzzles.
func RunGauntlet(playerName string) {
	s := puzzleScorer{playerName:playerName}

	s.solve("doomed1", DefiniteLoss)
	s.solve("doomed2", DefiniteLoss)
	s.solve("doomed3", DefiniteLoss)
	s.solve("doomed4", DefiniteLoss)
	s.solve("triangleBlock", DefiniteWin)
	s.solve("ladder", DefiniteWin)

	// Leave out manyBridges until doomed and triangle work, because it
	// should be even harder than those.
	// s.solve("manyBridges", DefiniteWin)

	s.solve("needle", ClearMove)
	s.solve("simpleBlock", ClearMove)

	log.Printf("SCORE: %d / %d", s.right, s.total)
}
