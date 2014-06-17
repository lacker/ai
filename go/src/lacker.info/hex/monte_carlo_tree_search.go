package hex

/*
Monte Carlo Tree Search.
*/

import (
	"math"
)

type TreeNode struct {
	BlackWins int
	WhiteWins int
	Board *Board
	Children map[Spot]*TreeNode
	Parent *TreeNode
}

func NewRoot(b *Board) *TreeNode {
	node := new(TreeNode)
	node.Board = b.Copy()
	node.Children = make(map[Spot]*TreeNode)
	return node
}

func (n *TreeNode) NumPlayouts() int {
	return n.BlackWins + n.WhiteWins
}

// The UCT formula for how promising this node is to investigate
func (n *TreeNode) UCT() float64 {
	if n.Parent == nil {
		// With no parent there are no alternative choices so this node
		// is infinitely promising
		return math.Inf(1)
	}
	var wins float64
	switch n.Board.ToMove {
	case White:
		wins = float64(n.WhiteWins)
	case Black:
		wins = float64(n.BlackWins)
	}
	sims := float64(n.NumPlayouts())
	if sims == 0 {
		// Always prefer an unexplored node
		return math.Inf(1)
	}
	total := float64(n.Parent.NumPlayouts())
	return (wins / sims) + 1.4 * math.Sqrt(math.Log(total) / sims)
}


type MonteCarloTreeSearch struct {
	Root *TreeNode
}

func (mcts MonteCarloTreeSearch) Play(b *Board) Spot {
	mcts.Root = NewRoot(b)
	panic("TODO: implement mcts algorithm")
}
