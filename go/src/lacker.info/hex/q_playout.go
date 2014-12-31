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

	panic("TODO")
}
