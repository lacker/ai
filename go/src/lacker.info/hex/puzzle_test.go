package hex

import (
	"testing"
)

type checker struct {
	Tester *testing.T
	Name string
	Player Player
}

func (c checker) check(puzzle Puzzle) {
	playerAnswer := c.Player.Play(puzzle.Board)
	if puzzle.CorrectAnswer != playerAnswer {
		c.Tester.Errorf("With puzzle: %s", puzzle.String)
		c.Tester.Errorf("%s gave incorrect answer: %s", c.Name, playerAnswer)
	}
}

func TestPuzzles(t *testing.T) {
	sr := checker{
		Tester: t,
		Name: "SR",
		Player: ShallowRave{Seconds:0.1, Quiet:true},
	}
	mcts := checker{
		Tester: t,
		Name: "MCTS",
		Player: MonteCarloTreeSearch{Seconds:0.1, Quiet:true},
	}

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

	sr.check(onePly)
	mcts.check(onePly)
}

