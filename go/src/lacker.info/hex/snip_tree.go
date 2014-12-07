package hex

import (
)

/*
A SnipTree is used to solve the question of, given two evolving
players, how can one alter its play to beat the second one?

The SnipTree itself is a tree of ways to diverge from the main line.

The key heuristic, which is unfortunately specific to Hex, is that if
there is a game tree in which an opponent always wins, but never
plays a particular move, then making that move cannot possibly be part
of a winning strategy for this player. Thus, we prioritize trying new
moves according to how frequently the opponent uses them in the game
tree descending from a particular position.
*/

type SnipTree struct {
	// The snips that define this tree node
	snipList []Snip

	// The node that is like this one, but does not have the last snip
	// in snipList.
	// We always investigate parents before children.
	parent *SnipTree

	// TODO: track how many times particular spots are played
	// by the opponent in the tree rooted at this node?
	// Depends on the specific algorithm we want.
}
