package hex

import (
	"fmt"
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
	// The ply is how far deep in the game to apply this snip with.
	// 0 = the first move in the game
	// 1 = the second player's first move
	// 2 = th first player's second move
	// This is also an index into History. After playing a game with
	// this snip, checking the plyth element of History should reflect
	// this snip.
	ply int

	// The spot to move for this player
	spot TopoSpot
}

func (s Snip) String() string {
	return fmt.Sprintf("%d => %s", s.ply, s.spot)
}

// Finds a list of Snips in chronological order that will let player
// beat opponent.
// player and opponent both need to be deterministic for this to work.
// mainLine should be a board showing the position where player lost
// to opponent.
// Returns the winning snip list along with the ending position.
func FindWinningSnipList(
	player QuickPlayer, opponent QuickPlayer, mainLine *TopoBoard,
	debug bool) ([]Snip, *TopoBoard) {

	// Sanity checks
	if player.Color() == opponent.Color() {
		log.Fatal("both player and opponent are the same color")
	}
	board := player.StartingPosition()
	if board != opponent.StartingPosition() {
		log.Fatal("starting positions do not match")
	}
	if mainLine.Winner != opponent.Color() {
		log.Fatal("mainLine is supposed to have player losing to opponent")
	}

	// The frontier is a list of snip lists we haven't tried yet.
	frontier := make([][]Snip, 0)

	// Current is a snip list we tried.
	var current []Snip = make([]Snip, 0)

	// ending is the ending position we get with the current snip list.
	ending := mainLine

	// Every viable ply is at least beginPly a la STL iterators
	beginPly := len(player.StartingPosition().History)

	for {
		// The current snip list failed to defeat the opponent.

		// We want to add new snip lists to the frontier.
		// We use the heuristic that the only reasonable snips are the moves
		// that the opponent plays in a game after the snip point.
		// We use breadth-first search on top of this heuristic.
		// A more nuanced heuristic might be better.

		// Figure out the first ply to consider a snip at.
		// Snips must be in order in the snip list, so we can start at the
		// previous one.
		var startPly int
		if len(current) == 0 {
			// There are no snips in current, so the first ply to consider a
			// snip at is the player's first move after the starting
			// position.
			if player.StartingPosition().GetToMove() == player.Color() {
				startPly = beginPly
			} else {
				startPly = beginPly + 1
			}
		} else {
			startPly = current[len(current) - 1].ply + 2
		}

		// Figure out which ply to snip at
		for snipPly := startPly; snipPly < len(ending.History); snipPly += 2 {
			// Figure out which move to insert
			for oppoPly := snipPly + 1; oppoPly < len(ending.History); oppoPly += 2 {
				snip := Snip{ply: snipPly, spot: ending.History[oppoPly]}
				frontier = append(frontier, append(current, snip))
			}
		}

		// So we added new snip lists to the frontier. That means we are
		// done with current. It is time to play a new game with the next
		// snip list.
		if len(frontier) == 0 {
			log.Fatal("ran out of frontier")
		}
		current = frontier[0]
		frontier = frontier[1:]
		ending = PlayoutWithSnipList(player, opponent, current, false)

		if ending.Winner == player.Color() {
			// This snip list made player win!
			if debug {
				log.Printf("%s wins with snip list: %+v",
					player.Color().Name(), current)
				ending.Debug()
			}
			return current, ending
		}

		// This snip list also did not succeed. Just continue through to
		// the next iteration of the loop.
	}
}
