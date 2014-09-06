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
	Board Board
	NumPossibleMoves int
	Children map[Spot]*TreeNode
	Parent *TreeNode

	// Holds parameters that may affect the search
	Strategy *MonteCarloTreeSearch

	// With classic MCTS, the rave counters track a win count for
	// playouts that are descendants of this node where the
	// Board.GetToMove() player moved at this particular spot at
	// some point in the playout.
	// With "topo" logic, this counts the number of times that this spot
	// was part of a winning path for each player.
	RaveBlackWins [NumSpots]int
	RaveWhiteWins [NumSpots]int
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
	node.Strategy = parent.Strategy
	if node.Strategy.UseTopoBoards {
		node.Board = parent.Board.ToTopoBoard()
	} else {
		node.Board = parent.Board.ToNaiveBoard()
	}
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

// The win rate just purely based on what this node has done before.
func (n *TreeNode) SimpleExpectedWinRate() float64 {
	var wins float64
	switch n.Board.GetToMove() {
	case Black:
		wins = float64(n.BlackWins)
	case White:
		wins = float64(n.WhiteWins)
	}
	return (1.0 + wins) / (2.0 + float64(n.BlackWins) + float64(n.WhiteWins))
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
	switch n.Board.GetToMove() {
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

func (n *TreeNode) ToMoveLetter() string {
	switch n.Board.GetToMove() {
	case Black:
		return "B"
	case White:
		return "W"
	}
	panic("bad tomove")
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
func (n *TreeNode) Backprop(winner Color, finalBoard Board) {
	// Update regular win/loss stats
	switch winner {
	case Black:
		n.BlackWins++
	case White:
		n.WhiteWins++
	}

	// Update rave stats.
	if n.Strategy.UseTopoBoards {
		// "Topo" scoring.
		// A spot counts towards rave if it was part of the winning path,
		// regardless of whose move it currently is.
		for _, move := range finalBoard.GetWinningPathSpots() {
			if n.Board.Get(move) != Empty {
				continue
			}
			switch winner {
			case Black:
				n.RaveBlackWins[move.Index()]++
			case White:
				n.RaveWhiteWins[move.Index()]++
			}
		}
	} else {
		// "Classic" scoring.
		// A spot counts towards rave if it was played on by the current
		// active player. Then, whatever side it won for gets counted.
		for index, move := range AllSpots() {
			if finalBoard.Get(move) != n.Board.GetToMove() {
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
	}

	if n.Parent != nil {
		n.Parent.Backprop(winner, finalBoard)
	}
}

func (n *TreeNode) String() string {
	return fmt.Sprintf("(P: %d, B:%d, W:%d, %sEV:%.2f)",
		n.BlackWins + n.WhiteWins, n.BlackWins, n.WhiteWins,
		n.ToMoveLetter(), n.SimpleExpectedWinRate())
}

type MonteCarloTreeSearch struct {
	Seconds float64
	Quiet bool

	// Parameter controlling rave mixing.
	V int

	// Whether to use topo boards
	UseTopoBoards bool
}

func MakeMCTS(seconds float64) MonteCarloTreeSearch {
	return MonteCarloTreeSearch{
		Seconds: seconds,
		Quiet: false,
		V: 1000,
		UseTopoBoards: false,
	}
}


func (mcts *MonteCarloTreeSearch) NewRoot(b Board) *TreeNode {
	node := new(TreeNode)
	if mcts.UseTopoBoards {
		node.Board = b.ToTopoBoard()
	} else {
		node.Board = b.ToNaiveBoard()
	}
	node.Children = make(map[Spot]*TreeNode)
	node.NumPossibleMoves = len(node.Board.PossibleMoves())
	node.Strategy = mcts
	return node
}

// The expected win rate of a particular move.
//
// If V is greater than 0, V used as the rave mixing parameter.
// That means rave stats are used to fill in until we have V real
// games.
//
// If V equals 0, this does something hacky, and uses
// a Dirichlet backoff from exact to rave to a constant.
// In general this is not as good as a V around 1000.
func (mcts *MonteCarloTreeSearch) ExpectedWinRate(
	parent *TreeNode, move Spot, child *TreeNode) float64 {
	var raveWins int
	var raveLosses int
	switch parent.Board.GetToMove() {
	case Black:
		raveWins = parent.RaveBlackWins[move.Index()]
		raveLosses = parent.RaveWhiteWins[move.Index()]
	case White:
		raveLosses = parent.RaveBlackWins[move.Index()]
		raveWins = parent.RaveWhiteWins[move.Index()]
	}

	var raveWinRate float64
	if mcts.UseTopoBoards {
		// The win rate estimation method here is, as far as I know, not
		// reflected in any published papers and is totally me making
		// things up.
		// First, we need a prior for win rate that predates the rave
		// data. Use the overall win rate for the parent.
		var parentWins int
		var parentLosses int
		switch parent.Board.GetToMove() {
		case Black:
			parentWins = parent.BlackWins
			parentLosses = parent.WhiteWins
		case White:
			parentLosses = parent.BlackWins
			parentWins = parent.WhiteWins
		}

		// There are three types of topo playout:
		// "neutral" playouts, where this spot isn't played
		// "win" playouts, where this spot is played by the current player
		// "loss" playouts, where this spot is played by the opponent. 
		parentPlayouts := parentWins + parentLosses
		if parentPlayouts <= 0 {
			panic("parentPlayouts <= 0")
		}
		neutralPlayouts := parentPlayouts - raveWins - raveLosses
		if neutralPlayouts < 0 {
			panic("neutralPlayouts < 0")
		}

		// Imagine that each topo playout reflects an infinite number of
		// playouts, each with the same winning path, but with spots other
		// than the winning path chosen randomly.
		// In this case, the topo playouts are just a representation of an
		// infinite number of theoretical playouts.
		// We want raveWinRate to be an estimate of what the real rave win
		// rate would be if we had all of these infinite playouts.
		//
		// There is a hidden variable, the neutralWinRate which is how
		// often does the player win for games where this spot is not
		// played.
		//
		// We can calculate neutralWinRate by observing that
		// parentWins = neutralWinRate * neutralPlayouts + raveWins
		// Thus:
		neutralWinRate := (
			float64(parentWins - raveWins) /
				float64(neutralPlayouts))

		if neutralWinRate < 0 {
			panic("neutralWinRate < 0")
		}

		// Let's say each topo playout reflects a very large number of
		// games with cardinality 2X.
		//
		// In games corresponding to the "neutral" topo playouts, about half
		// of the time this spot will be played.
		// That will be neutralPlayouts * X games.
		//
		// In games corresponding to the "win" topo playouts, all of the
		// time this spot will be played.
		// That will be 2X * raveWins games.
		//
		// We can use this plus the win rates for each type of playout to
		// get the rave win rate on these very large number of games.
		// We can also drop the X, it's just a mental convenience.
		numGames := neutralPlayouts + 2 * raveWins
		numWins := float64(neutralPlayouts) * neutralWinRate + float64(raveWins)
		raveWinRate = (0.1 + numWins) / (0.1 + float64(numGames))
	} else {
		// Calculate a rave estimate with weak but win-slanted prior.
		// The rave wins and losses account for all playouts, so we can
		// just use the win rate among them as an estimate for the real
		// win rate.
		raveWinRate = (1.0 + float64(raveWins)) /
			(1.0 + float64(raveWins + raveLosses))
	}

	if child == nil {
		return raveWinRate
	}

	// Gather the specific win data
	var wins float64
	switch parent.Board.GetToMove() {
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

	if mcts.V == 0 {
		// Use the precise-node data to calculate the win rate, with the
		// rave-based calculation as a prior
		raveStrength := 20.0
		return (raveWinRate * raveStrength + wins) / (raveStrength + sims)
	}

	// If we have less than V real games, use the rave stats to fill in.
	v := float64(mcts.V)
	if sims >= v {
		return wins / sims
	}
	return (wins + (v - sims) * raveWinRate) / v
}

// Uses ExpectedWinRate to figure out which move is expected to be the
// best.
func (mcts *MonteCarloTreeSearch) ExpectedBestMove(n *TreeNode) (
	Spot, *TreeNode, float64) {

	bestWinRate := math.Inf(-1)
	var bestMove Spot
	var bestChild *TreeNode
	for move, child := range n.Children {
		winRate := mcts.ExpectedWinRate(n, move, child)
		if winRate > bestWinRate {
			bestWinRate = winRate
			bestChild = child
			bestMove = move
		}
	}

	if bestChild == nil {
		panic("could not find a child")
	}

	return bestMove, bestChild, bestWinRate
}

// Selects a leaf node recursively from the provided tree.
// A leaf node is defined as a node where either a new child could be added,
// or there are no possible children and the game is over.
func (mcts *MonteCarloTreeSearch) SelectLeaf(n *TreeNode) *TreeNode {
	if n.NumPossibleMoves > len(n.Children) {
		return n
	}
	if n.NumPossibleMoves == 0 {
		return n
	}

	_, bestChild, _ := mcts.ExpectedBestMove(n)

	return mcts.SelectLeaf(bestChild)
}

func (mcts *MonteCarloTreeSearch) RunOneRound(n *TreeNode) {
	leaf := mcts.SelectLeaf(n).Expand()
	board := leaf.Board.Copy()
	winner := board.Playout()
	leaf.Backprop(winner, board)
}

func (mcts MonteCarloTreeSearch) Play(b Board) (Spot, float64) {
	start := time.Now()
	root := mcts.NewRoot(b)

	// Do playouts for a set amount of time
	for SecondsSince(start) < mcts.Seconds {
		mcts.RunOneRound(root)
	}

	for _, move := range AllSpots() {
		child, ok := root.Children[move]
		if ok && !mcts.Quiet && (child.WhiteWins + child.BlackWins >= 500) {
			log.Printf("%s -- %s", move, child)			
		}
	}

	if !mcts.Quiet {
		log.Printf("MCTS: %s", root)
	}

	move, _, score := mcts.ExpectedBestMove(root)
	return move, score
}
