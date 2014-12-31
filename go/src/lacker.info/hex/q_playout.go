package hex

import (
)

// A playout between two QNets.

type QPlayout struct {
	// All of the actions that were taken during the game.
	actions []QAction

	// Which color won.
	winner Color
}

func (playout *QPlayout) AddAction(action QAction) {
	playout.actions = append(playout.actions, action)
}

func NewQPlayout(player1 *QNet, player2 *QNet) *QPlayout {
	playout := &QPlayout{
		actions: []QAction{},
		winner: Empty,
	}

	player1.Reset()
	player2.Reset()

	board := player1.StartingPosition().ToTopoBoard()
	for board.Winner == Empty {
		// player is the player whose move it is
		var player *QNet
		switch board.GetToMove() {
		case player1.Color():
			player = player1
		case player2.Color():
			player = player2
		default:
			panic("busted switch")
		}

		action := player.Act(board)
		playout.actions = append(playout.actions, action)

		feature := MakeQFeature(action.color, action.spot)
		player1.AddFeature(feature)
		player2.AddFeature(feature)
	}

	playout.winner = board.Winner
	return playout
}
