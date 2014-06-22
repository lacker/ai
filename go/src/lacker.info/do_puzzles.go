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

	board, correctAnswer := hex.NewPuzzle(`
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
	playerAnswer := player.Play(board)
	if correctAnswer != playerAnswer {
		panic("wrong answer")
	}
}
