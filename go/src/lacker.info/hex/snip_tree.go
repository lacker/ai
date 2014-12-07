package hex

import (
)

/*
A SnipTree is used to solve the question of, given two evolving
players, how can one alter its play to beat the second one?

Each node in the SnipTree represents a position in the game, with the
"evolving" player to move.

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

	// The move index in the game of the next move to be made
	index int

	// The snip tree representing the game two ply before this one
	parent *SnipTree

	// TODO: track how many times particular spots are played
	// by the opponent in the tree rooted at this node?
	// Depends on the specific algorithm we want.
}
