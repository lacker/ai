package main

import (
	"lacker.info/hex"
)

/*
Loads specific positions and tests to see if the algorithms can solve
them correctly.
*/

func main() {
	hex.Seed()

	puzzle := hex.MakePuzzle(`
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

	player := hex.GetPlayer("sr1")
	playerAnswer := player.Play(puzzle.Board)
	if puzzle.CorrectAnswer != playerAnswer {
		panic("wrong answer")
	}
}
