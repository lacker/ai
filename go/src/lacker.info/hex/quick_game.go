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

	// The board to run the playout on
	board *TopoBoard

	// An optional override to control what the players do.
	snipList []Snip

	// An optional registry to notify when moves are made.
	registry *SpotRegistry
}

func NewQuickGame(p1 QuickPlayer, p2 QuickPlayer, debug bool) *QuickGame {
	if p1.Color() == p2.Color() {
		log.Fatal("both players are the same color")
	}

	if p1.StartingPosition() != p2.StartingPosition() {
		log.Fatal("starting positions don't match")
	}

	game := QuickGame{
		player1: p1,
		player2: p2,
		debug: debug,
	}

	game.board = game.player1.StartingPosition().ToTopoBoard()
	game.player1.Reset()
	game.player2.Reset()

	return &game
}

// Makes the provided move and signals on the registry
func (game *QuickGame) MakeMove(spot TopoSpot) {
	game.board.MakeMove(spot)
	if game.registry != nil {
		game.registry.Notify(spot)
	}
}

// Makes the move that the current player would make
func (game *QuickGame) MakeMoveForCurrentPlayer() {
	var player QuickPlayer
	if game.player1.Color() == game.board.GetToMove() {
		player = game.player1
	} else {
		player = game.player2
	}

	spot, score := player.BestMove(game.board, game.debug)
	game.MakeMove(spot)
	if game.debug {
		log.Printf("%s moves %s with score %.2f",
			player.Color().Name(), spot.String(), score)
	}
}

// Plays out the game and returns the final board state.
// You are only supposed to call Playout once per QuickGame.
func (game *QuickGame) Playout() *TopoBoard {
	snipListIndex := 0

	// Play the playout
	for game.board.Winner == Empty {
		if game.snipList != nil && len(game.snipList) > snipListIndex &&
			game.snipList[snipListIndex].ply == len(game.board.History) {
			// The snip list overrides the player
			game.MakeMove(game.snipList[snipListIndex].spot)
			snipListIndex++
		} else {
			game.MakeMoveForCurrentPlayer()
		}
	}

	if game.debug {
		log.Printf("%s wins the playout", game.board.Winner.Name())
	}
	return game.board
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
	game.snipList = snipList
	return game.Playout()
}

