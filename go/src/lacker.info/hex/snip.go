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
	// 2 = the first player's second move
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

// Figures out if it's possible to win from a given position.
// The position is defined by the snip list to get here, plus the
// index of the move we are considering alternatives for.
// moveIndex should always be player's turn.
// For convenience we also have a playout for this position.
//
// If we find a win, returns the snip list and topo board that
// correspond to the win.
// This always returns a count of the number of times each spot was
// used to defeat us in the search rooted at this position.
// This does a depth-first search so that it can be implemented
// recursively.
func FindWinFromPosition(
	player QuickPlayer, opponent QuickPlayer, playout *TopoBoard,
	snipList []Snip, moveIndex int) (
		[]Snip, *TopoBoard, [NumTopoSpots]int) {
	if playout.ColorForHistoryIndex(moveIndex) != player.Color() {
		log.Fatalf("moveIndex (%d) should always be player's (%s's) move",
			moveIndex, player.Color())
	}

	// Base case: if the game is already over at moveIndex, then there's
	// no win from this position.
	if moveIndex >= len(playout.History) {
		var allZero [NumTopoSpots]int
		return nil, nil, allZero
	}

	// See if we can solve this game by first playing what player
	// recommends.
	answer1, answer2, defeatCount := FindWinFromPosition(
		player, opponent, playout, snipList, moveIndex + 2)
	if answer1 != nil {
		// Found a win
		return answer1, answer2, defeatCount
	}

	// To make defeatCount correct for this node, we have to add in the
	// move that was used to immediately respond to our default move.
	defeatCount[playout.History[moveIndex + 1]]++

	// We will recurse with different alternatives for this move.
	// So, keep track of which positions we recursively try.
	// So far, we have only tried the default move.
	var tried [NumTopoSpots]bool
	tried[playout.History[moveIndex]] = true

	for {
		// Find an alternative move for this spot.
		// Pick the highest defeatCount move that we haven't tried yet.
		bestSpot, bestCount := NotASpot, 0
		for spot := TopLeftCorner; spot <= BottomRightCorner; spot++ {
			count := defeatCount[spot]
			if !tried[spot] && count > bestCount {
				bestSpot = spot
				bestCount = count
			}
		}

		if bestSpot == NotASpot {
			// There are no spots left to be tried.
			return nil, nil, defeatCount
		}
		
		// Try snipping to bestSpot at moveIndex.
		newSnipList := append(snipList,
			Snip{ply:moveIndex, spot:bestSpot})
		newPlayout := PlayoutWithSnipList(player, opponent, newSnipList,
			false)
		tried[bestSpot] = true
		if newPlayout.Winner == player.Color() {
			// Found a win
			return newSnipList, newPlayout, defeatCount
		}

		// Try recursing on this new snip list.
		answer1, answer2, newDefeatCount := FindWinFromPosition(
			player, opponent, newPlayout, newSnipList, moveIndex + 1)

		// Add in the defeats to make defeatCount correct.
		for spot := TopLeftCorner; spot <= BottomRightCorner; spot++ {
			defeatCount[spot] += newDefeatCount[spot]
		}
		// We also need the move that was used to immediately respond to
		// our default move.
		defeatCount[newPlayout.History[moveIndex + 1]]++

		if answer1 != nil {
			// Found a win
			return answer1, answer2, defeatCount
		}

		// There's no win with this snip, so continue through the loop to
		// look again.
	}
}

// Finds a list of Snips in chronological order that will let player
// beat opponent, using breadth-first search.
// player and opponent both need to be deterministic for this to work.
// mainLine should be a board showing the position where player lost
// to opponent.
// snipLimit is the maximum length of snip list to search for, or 0
// for no limit.
// One critical problem with this function is that it might miss an
// existing solution.
// If it's impossible to find a winning snip list, this returns nils.
// Returns the winning snip list along with the ending position.
func FindWinningSnipList(
	player QuickPlayer, opponent QuickPlayer, mainLine *TopoBoard,
	snipLimit int, debug bool) ([]Snip, *TopoBoard) {

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

	attempts := 0
	for {
		// The current snip list failed to defeat the opponent.

		if snipLimit == 0 || snipLimit > len(current) {
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
		}

		// So we added new snip lists to the frontier. That means we are
		// done with current. It is time to play a new game with the next
		// snip list.
		if len(frontier) == 0 {
			// We can't find a winning snip list. The opponent is unbeatable.
			return nil, nil
		}
		current = frontier[0]
		frontier = frontier[1:]
		ending = PlayoutWithSnipList(player, opponent, current, false)
		attempts++

		if ending.Winner == player.Color() {
			// This snip list made player win!
			if debug {
				log.Printf("after %d attempts, %s wins with snip list: %+v",
					attempts, player.Color().Name(), current)
				ending.Debug()
			}
			return current, ending
		}

		// This snip list also did not succeed. Just continue through to
		// the next iteration of the loop.
	}
}
