package hex

/*
Monte Carlo Tree Search.
*/

import (
	"fmt"
	"log"
	"math"
	"time"
)

type TreeNode struct {
	BlackWins int
	WhiteWins int
	Board *Board
	NumPossibleMoves int
	Children map[Spot]*TreeNode
	Parent *TreeNode
}

func NewRoot(b *Board) *TreeNode {
	node := new(TreeNode)
	node.Board = b.Copy()
	node.Children = make(map[Spot]*TreeNode)
	node.NumPossibleMoves = len(node.Board.PossibleMoves())
	return node
}

func NewChild(parent *TreeNode, move Spot) *TreeNode {
	if parent == nil {
		panic("cannot create a child of nil")
	}
	if parent.Children[move] != nil {
		panic("cannot create a duplicate child")
	}
	node := new(TreeNode)
	if parent.Board == nil {
		panic("bad parent - board should not be nil")
	}
	node.Board = parent.Board.Copy()
	if !node.Board.MakeMove(move) {
		panic("cannot create new child with invalid move")
	}
	parent.Children[move] = node
	node.Children = make(map[Spot]*TreeNode)
	node.NumPossibleMoves = parent.NumPossibleMoves - 1
	node.Parent = parent
	return node
}

func (n *TreeNode) NumPlayouts() int {
	return n.BlackWins + n.WhiteWins
}

// The UCT formula for how promising this node is to investigate.
// The formula for a node should answer the question of, how good is
// it to make the move that *gets* to this node.
// Thus, it is optimizing for the player that is *not* to move.
func (n *TreeNode) UCT() float64 {
	if n.Parent == nil {
		// With no parent there are no alternative choices so this node
		// is infinitely promising
		return math.Inf(1)
	}
	var wins float64
	switch n.Board.ToMove {
	case Black:
		wins = float64(n.WhiteWins)
	case White:
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

// Finds the move from this node with the most MCTS simulations.
// Panics if it can't find any move.
func (n *TreeNode) MostSimulatedMove() Spot {
	var bestMove Spot
	numSims := -1

	for move, child := range n.Children {
		childSims := child.BlackWins + child.WhiteWins
		if childSims > numSims {
			bestMove = move
			numSims = childSims
		}
	}

	if numSims == -1 {
		panic("could not find any move")
	}

	return bestMove
}

// Selects a leaf node recursively from the provided tree according to UCT.
// A leaf node is defined as a node where either a new child could be added,
// or there are no possible children and the game is over.
func (n *TreeNode) SelectLeaf() *TreeNode {
	if n.NumPossibleMoves > len(n.Children) {
		return n
	}
	if n.NumPossibleMoves == 0 {
		return n
	}

	bestUCT := math.Inf(-1)
	var bestChild *TreeNode
	for _, child := range n.Children {
		uct := child.UCT()
		if uct > bestUCT {
			bestUCT = uct
			bestChild = child
		}
	}

	if bestChild == nil {
		panic("could not find a child")
	}

	return bestChild.SelectLeaf()
}

// Expands from the given leaf node if possible by choosing a new
// possible child randomly and creating it.
// Returns the child if expansion was possible, or nil if it was not
// possible.
func (n *TreeNode) Expand() *TreeNode {
	if n.NumPossibleMoves <= len(n.Children) {
		return nil
	}
	possibleMoves := n.Board.PossibleMoves()
	ShuffleSpots(possibleMoves)
	for _, move := range possibleMoves {
		_, ok := n.Children[move]
		if !ok {
			return NewChild(n, move)
		}
	}
	panic("children everywhere")
}

func (n *TreeNode) Depth() int {
	if n == nil {
		return 0
	}
	answer := 1
	for _, child := range n.Children {
		answer = Intmax(answer, child.Depth() + 1)
	}
	return answer
}

// Backpropagate a win, starting at this node and continuing through
// parents until we hit the root.
func (n *TreeNode) Backprop(c Color) {
	switch c {
	case Black:
		n.BlackWins++
	case White:
		n.WhiteWins++
	}
	if n.Parent != nil {
		n.Parent.Backprop(c)
	}
}

func (n *TreeNode) String() string {
	return fmt.Sprintf("(B:%d, W:%d, UCT:%.2f)",
		n.BlackWins, n.WhiteWins, n.UCT())
}

func (n *TreeNode) RunOneRoundOfMCTS() {
	leaf := n.SelectLeaf().Expand()
	winner := leaf.Board.Copy().Playout()
	leaf.Backprop(winner)
}

type MonteCarloTreeSearch struct {
	Seconds time.Duration
	Root *TreeNode
}

func (mcts MonteCarloTreeSearch) Play(b *Board) Spot {
	start := time.Now()
	mcts.Root = NewRoot(b)

	// Do playouts for a set amount of time
	for time.Since(start) < mcts.Seconds * time.Second {
		mcts.Root.RunOneRoundOfMCTS()
	}

	for _, move := range AllSpots() {
		child, ok := mcts.Root.Children[move]
		if ok {
			log.Printf("%s -- %s", move, child)			
		}
	}

	log.Printf("Overall: %s", mcts.Root)

	return mcts.Root.MostSimulatedMove()
}
