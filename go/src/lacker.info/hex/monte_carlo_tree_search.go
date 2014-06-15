package hex

/*
Monte Carlo Tree Search.
*/

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


type MonteCarloTreeSearch struct {
}

func (mcts MonteCarloTreeSearch) Play(b *Board) Spot {
	panic("TODO: implement mcts algorithm")
}
