package hex

/*
Monte Carlo Tree Search.
*/

import (
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

// The UCT formula for how promising this node is to investigate
func (n *TreeNode) UCT() float64 {
	panic("TODO: implement")
}


type MonteCarloTreeSearch struct {
	Root *TreeNode
}

func (mcts MonteCarloTreeSearch) Play(b *Board) Spot {
	mcts.Root = NewRoot(b)
	panic("TODO: implement mcts algorithm")
}
