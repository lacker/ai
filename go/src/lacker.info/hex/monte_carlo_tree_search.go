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

	// A win count for playouts that are descendants of this node and
	// contain a particular move. The particular move is defined by the
	// spot that is the index of the array, and the Board.ToMove color
	// making it. The move particularly does not have to be made
	// immediately from this position.
	RaveBlackWins [NumSpots]int
	RaveWhiteWins [NumSpots]int
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

// The expected win rate of a particular move.
// Hopefully this formula is more sophicated than UCT and thus
// superior.
// This uses a Dirichlet backoff from exact to rave to a constant.
func (n *TreeNode) ExpectedWinRate(move Spot, child *TreeNode) float64 {
	// Calculate a rave estimate with weak but win-slanted prior
	var raveWins int
	var raveLosses int
	switch n.Board.ToMove {
	case Black:
		raveWins = n.RaveBlackWins[move.Index()]
		raveLosses = n.RaveWhiteWins[move.Index()]
	case White:
		raveLosses = n.RaveBlackWins[move.Index()]
		raveWins = n.RaveWhiteWins[move.Index()]
	}
	raveWinRate := (1.0 + float64(raveWins)) /
		(1.0 + float64(raveWins + raveLosses))

	if child == nil {
		return raveWinRate
	}

	// Gather the specific win data
	var wins float64
	switch n.Board.ToMove {
	case Black:
		wins = float64(child.BlackWins)
	case White:
		wins = float64(child.WhiteWins)
	}
	sims := float64(child.NumPlayouts())
	if sims <= 0.0 {
		return raveWinRate
	}
	if wins > sims {
		panic("cannot have more wins than sims")
	}

	// Use the precise-node data to calculate the win rate, with the
	// rave-based calculation as a prior
	raveStrength := 20.0
	return (raveWinRate * raveStrength + wins) / (raveStrength + sims)
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
	total := n.Parent.NumPlayouts()
	return (wins / sims) + 0.5 * math.Sqrt(Fastlog(total) / sims)
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

// Uses ExpectedWinRate to figure out which move is expected to be the
// best.
func (n *TreeNode) ExpectedBestMove() (Spot, *TreeNode) {
	bestWinRate := math.Inf(-1)
	var bestMove Spot
	var bestChild *TreeNode
	for move, child := range n.Children {
		winRate := n.ExpectedWinRate(move, child)
		if winRate > bestWinRate {
			bestWinRate = winRate
			bestChild = child
			bestMove = move
		}
	}

	if bestChild == nil {
		panic("could not find a child")
	}

	return bestMove, bestChild
}

// Selects a leaf node recursively from the provided tree.
// A leaf node is defined as a node where either a new child could be added,
// or there are no possible children and the game is over.
func (n *TreeNode) SelectLeaf() *TreeNode {
	if n.NumPossibleMoves > len(n.Children) {
		return n
	}
	if n.NumPossibleMoves == 0 {
		return n
	}

	_, bestChild := n.ExpectedBestMove()

	return bestChild.SelectLeaf()
}

// Selects a leaf node recursively using UCT.
func (n *TreeNode) SelectLeafByUCT() *TreeNode {
	if n.NumPossibleMoves > len(n.Children) {
		return n
	}
	if n.NumPossibleMoves == 0 {
		return n
	}

	bestUCT := math.Inf(-1)
	var bestChild *TreeNode
	for _, child := range n.Children {
		UCT := child.UCT()
		if UCT > bestUCT {
			bestUCT = UCT
			bestChild = child
		}
	}

	if bestChild == nil {
		panic("could not find a child")
	}

	return bestChild.SelectLeafByUCT()
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

// Backpropagate a win, starting at this node and continuing
// through parents until we hit the root.
func (n *TreeNode) Backprop(winner Color, finalBoard *Board) {
	// Update regular win/loss stats
	switch winner {
	case Black:
		n.BlackWins++
	case White:
		n.WhiteWins++
	}

	// Update rave stats
	for index, move := range AllSpots() {
		if finalBoard.Get(move) != n.Board.ToMove {
			continue
		}
		if n.Board.Get(move) != Empty {
			continue
		}
		switch winner {
		case Black:
			n.RaveBlackWins[index]++
		case White:
			n.RaveWhiteWins[index]++
		}
	}

	if n.Parent != nil {
		n.Parent.Backprop(winner, finalBoard)
	}
}

func (n *TreeNode) String() string {
	return fmt.Sprintf("(P: %d, B:%d, W:%d, UCT:%.2f)",
		n.BlackWins + n.WhiteWins, n.BlackWins, n.WhiteWins, n.UCT())
}

func (n *TreeNode) RunOneClassicRound() {
	leaf := n.SelectLeafByUCT().Expand()
	board := leaf.Board.Copy()
	winner := board.Playout()
	leaf.Backprop(winner, board)
}

func (n *TreeNode) RunOneModernRound() {
	leaf := n.SelectLeaf().Expand()
	board := leaf.Board.Copy()
	winner := board.Playout()
	leaf.Backprop(winner, board)
}

type MonteCarloTreeSearch struct {
	Modern bool
	Seconds time.Duration
}

func (mcts MonteCarloTreeSearch) Play(b *Board) Spot {
	start := time.Now()
	root := NewRoot(b)

	// Do playouts for a set amount of time
	for time.Since(start) < mcts.Seconds * time.Second {
		if mcts.Modern {
			root.RunOneModernRound()
		} else {
			root.RunOneClassicRound()
		}
	}

	for _, move := range AllSpots() {
		child, ok := root.Children[move]
		if ok {
			log.Printf("%s -- %s", move, child)			
		}
	}

	log.Printf("MCTS: %s", root)

	if mcts.Modern {
		move, _ := root.ExpectedBestMove()
		return move
	} else {
		return root.MostSimulatedMove()
	}
}
