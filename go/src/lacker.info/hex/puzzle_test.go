package hex

import (
	"math/rand"
	"testing"
)

type checker struct {
	Tester *testing.T
	Name string
	Player Player
}

func (c checker) expectPass(puzzle Puzzle) {
	playerAnswer, _ := c.Player.Play(puzzle.Board)
	if puzzle.CorrectAnswer != playerAnswer {
		c.Tester.Errorf("With puzzle: %s", puzzle.String)
		c.Tester.Errorf("%s gave incorrect answer: %s", c.Name, playerAnswer)
	}
}

func (c checker) expectFail(puzzle Puzzle) {
	playerAnswer, _ := c.Player.Play(puzzle.Board)
	if puzzle.CorrectAnswer == playerAnswer {
		c.Tester.Errorf("With puzzle: %s", puzzle.String)
		c.Tester.Errorf("%s was supposed to fail but passed.", c.Name)
	}
}

func TestPuzzles(t *testing.T) {
	rand.Seed(1)
	
	sr := checker{
		Tester: t,
		Name: "SR",
		Player: ShallowRave{Seconds:0.1, Quiet:true},
	}
	mcts := checker{
		Tester: t,
		Name: "MCTS",
		Player: MonteCarloTreeSearch{Seconds:0.2, Quiet:true, V:0},
	}

	// Any reasonable method should be able to find a killer move.

	onePly := MakePuzzle(`
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

	sr.expectPass(onePly)
	mcts.expectPass(onePly)

	// Tree methods can figure out a block where shallow rave can't,
	// because they can figure out the bridges.
	// MCTS can figure this out consistently in 0.2s, but not in 0.1s.

	triangleBlock := MakePuzzle(`
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
	sr.expectFail(triangleBlock)
	mcts.expectPass(triangleBlock)

	// Tree methods still cannot understand a large amount of bridges.
	// MCTS can occasionally pass this but usually can't.

	manyBridges := MakePuzzle(`
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
	sr.expectFail(manyBridges)
	// mcts.expectFail(manyBridges)
}

