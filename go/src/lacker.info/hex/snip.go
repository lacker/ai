package hex

import (
	"log"
)

/*
A Snip is a single alteration to be made to a quickplayer's game.
The main goal of Snips is to find a small set of them that would lead
to a particular quickplayer winning instead of losing in a
particular matchup. Then they can be used for learning.
"Snip" is an allusion to a SNP = Single Nucleotide Polymorphism which
is a mutation that only hits a single spot in a DNA strand and also
pronounced "Snip".
*/

type Snip struct {
	// The moveNumber is how far in the future for this player to apply
	// the Snip. *Not* ply.
	// 0 = the next move this player makes
	// 1 = the one after that
	moveNumber int

	// The spot to move for this player
	spot TopoSpot
}

// Finds a list of Snips in chronological order that will let player
// beat opponent.
// player and opponent both need to be deterministic for this to work.
// mainLine should be a board showing the position where player lost
// to opponent.
func FindWinningSnipList(
	player QuickPlayer, opponent QuickPlayer, mainLine *TopoBoard) []Snip {
	// Sanity checks
	if player.Color() == opponent.Color() {
		log.Fatal("both player and opponent are the same color")
	}
	board := player.StartingPosition()
	if board != opponent.StartingPosition() {
		log.Fatal("starting positions do not match")
	}

	// Use breadth-first search because it seems like the easiest
	// reasonable thing.
	// The frontier is a list of snip lists we haven't tried yet.
	frontier := make([][]Snip, 0)

	// Current is a snip list we tried.
	var current []Snip = make([]Snip, 0)

	// ending is the ending position we get with the current snip list.
	ending := mainLine

	for {
		// The current snip list failed to defeat the opponent.
		// Add new snip lists to the frontier.
		// We use the heuristic that the only reasonable snips are the ones
		// that the opponent plays in a game.
		// TODO: we need game history for this really, add to TopoBoard
		log.Fatal("TODO", frontier, current, ending)
	}
}
