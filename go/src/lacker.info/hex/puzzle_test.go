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

	onePly := PuzzleMap["onePly"]
	sr.expectPass(onePly)
	mcts.expectPass(onePly)

	triangleBlock := PuzzleMap["triangleBlock"]
	sr.expectFail(triangleBlock)
	mcts.expectPass(triangleBlock)

	manyBridges := PuzzleMap["manyBridges"]
	sr.expectFail(manyBridges)

	// This mostly fails but not always
	// mcts.expectFail(manyBridges)
}

