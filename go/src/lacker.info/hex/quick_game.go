package hex

import (
	"log"
)

// A QuickGame object keeps around all the details of a single playout.
// This is kind of like an options object; there are just a lot of
// different ways to handle a playout.

type QuickGame struct {
	player1 QuickPlayer
	player2 QuickPlayer
	debug bool

	// An optional override to control what the players do.
	SnipList []Snip

	// An optional registry to notify when moves are made.
	Registry *SpotRegistry
}

func NewQuickGame(p1 QuickPlayer, p2 QuickPlayer, debug bool) *QuickGame {
	if p1.Color() == p2.Color() {
		log.Fatal("both players are the same color")
	}

	if p1.StartingPosition() != p2.StartingPosition() {
		log.Fatal("starting positions don't match")
	}

	return &QuickGame{
		player1: p1,
		player2: p2,
		debug: debug,
	}
}

// A helper for playouts
// TODO: refactor to be a member function
func MakeBestMove(player QuickPlayer, board *TopoBoard, debug bool) {
	if player.Color() != board.GetToMove() {
		log.Fatal("not the right player's turn")
	}
	spot, score := player.BestMove(board, debug)
	board.MakeMove(spot)
	if debug {
		log.Printf("%s moves %s with score %.2f",
			player.Color().Name(), spot.String(), score)
	}
}

// Plays out the game and returns the final board state.
// You are only supposed to call Playout once per QuickGame.
func (game *QuickGame) Playout() *TopoBoard {
	// Prepare for the game.
	// Run the playout on a copy so that we don't alter the original
	board := game.player1.StartingPosition().ToTopoBoard()
	game.player1.Reset()
	game.player2.Reset()
	snipListIndex := 0

	// Play the playout
	for board.Winner == Empty {
		if game.SnipList != nil && len(game.SnipList) > snipListIndex &&
			game.SnipList[snipListIndex].ply == len(board.History) {
			// The snip list overrides the player
			board.MakeMove(game.SnipList[snipListIndex].spot)
			snipListIndex++
		} else if game.player1.Color() == board.GetToMove() {
			MakeBestMove(game.player1, board, game.debug)
		} else {
			MakeBestMove(game.player2, board, game.debug)
		}
	}

	if game.debug {
		log.Printf("%s wins the playout", board.Winner.Name())
	}
	return board
}

// A helper for QuickGame.Playout
// Plays out a game and returns the final board state.
func Playout(
	player1 QuickPlayer, player2 QuickPlayer, debug bool) *TopoBoard {
	return PlayoutWithSnipList(player1, player2, nil, debug)
}

// A helper for QuickGame.Playout
// Plays out a game, overriding the players whenever mandated to by
// the snip list. Returns the final board state.
func PlayoutWithSnipList(
	player1 QuickPlayer, player2 QuickPlayer,
	snipList []Snip, debug bool) *TopoBoard {

	game := NewQuickGame(player1, player2, debug)
	game.SnipList = snipList
	return game.Playout()
}

