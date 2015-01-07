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

// Each playout defines a gradient, of the direction Q should go in
// order to improve its accuracy according to Q-learning.
// See q_net.go for details on the Q function.
//
// In general, the gradient of the error on the logit can be
// determined from the actual probabilities.
// If the real probability was p_real and we predicted p, then the
// gradient on the logit is just:
// p_real - p
// in every feature that the QNet sums to the logit for Q(s, a).
//
// To get the overall gradient for a whole playout, we need to do some
// dynamic programming. The Q-learning rule defines one update that
// can happen for each decision the provided color made. Since each
// feature becomes active at a single point during the playout and
// remains active for each successive Q-learning opportunity, we can
// simultaneously apply each of the learning rules to each feature
// using dynamic programming. See the code for more detail here.
//
// AddGradient adds scalar times the gradient to addend, using the
// gradient for the provided color's decisions.
func (playout *QPlayout) AddGradient(color Color, scalar float64,
	addend *[NumFeatures]float64) {
	panic("TODO")
}
