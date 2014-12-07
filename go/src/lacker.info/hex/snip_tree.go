package hex

import (
)

/*
A SnipTree is used to solve the question of, given two evolving
players, how can one alter its play to beat the second one?

Each node in the SnipTree represents a position in the game, with the
"evolving" player to move.

The key principle, which is unfortunately specific to Hex, is that if
an opponent can beat any move besides a move in the set S without ever
moving in the set S, then we don't need to investigate moves in the
set S to know that this position is lost.
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
