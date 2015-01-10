package hex

import (
	"math"
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
//
// The way this gradient is defined is fairly confusing so mentally
// prepare yourself to read complex stuff. I'll try to explain
// everything here, but it might help to hit up Wikipedia on Q-learning.
//
// The general Q-learning rule is applied to a single decision. You
// have some scalar that influences how much you learn from this
// particular case, and you move your neural network according to the
// gradient of a loss function. The loss function is defined by the
// Q-value you used to make the decision, and a "target" Q-value that
// is the ideal one you're moving towards.
//
// For Q-learning in general the target is  defined by the next
// Q-value in the sequence, with some future-discounting. In this case
// it seems like the right strategy is to not discount the future, and
// there are no mid-game rewards, so the target Q-value is just the
// next Q-value this playout had. The only distinction is that if we
// next chose a suboptimal move for exploration purposes, we should
// use the optimal for the purposes of learning. That gives us
// "off-policy" learning. (With the exception of our opponent's
// policy. That is fixable but I think we should ignore for now.)
//
// For our loss function we use cross-entropy and interpret Q as a
// logit of the probability of victory. Cross-entropy is best thought
// of as, if you predict a probability of an outcome is p, and you are
// correct, your reward is 0. If you are incorrect, your reward is
// -log(p). If an event is a weighted coin-flip and you are rewarded
// according to cross-entropy, your optimal strategy is to use the
// correct weight as your prediction. Yay information theory.
// 
// Now we can calculate what gradient this principle defines.
// (Insert real math here, left as an exercise for the reader.)
// 
// If the target probability was p_target and we predicted p_calc, then the
// gradient on the logit is just the probability difference:
// p_target - p_calc
// in every feature that the QNet sums to the logit for Q(s, a).
//
// Here it's convenient that the logit is just the sum of a bunch of
// weights. That means that, although the gradient is a vector over
// the weights on all the feature sets, it'll have the same magnitude
// in all of them and we can think of the gradient as just a single
// number. (As long as we're just talking about the update for a
// single move rather than the entire game.)
//
// So it's pretty cool that to get the gradient we just subtract two
// things that already have a sane meaning. This is a nice, fancy
// property of logits that IMO makes it worthwhile to use logistic
// stuff in the first place rather than some hacky more-comprehensible
// formula.
// By "the real probability" we'll take the next Q-value in the
// sequence, figuring that at least it'll be more accurate.
//
// Okay. Take a deep breath.
//
// Everything so far has been standard Q-learning stuff. Here is where
// we start to have custom tricks.
//
// The key is, since our neural net does incremental updates, a lot of
// the net's state is shared between subsequent game states. So we can
// actually apply the Q-learning rule to every stage of the game in
// one pass, and get the learning gradient for an entire playout
// rather than a single move at a time.
//
// To get the overall gradient for a whole playout, we need to do some
// dynamic programming. The Q-learning rule defines one update that
// can happen for each decision the provided color made. Since each
// feature becomes active at a single point during the playout and
// remains active for each successive Q-learning opportunity, we can
// apply each of the learning rules to each feature by keeping an
// accumulator of the gradient magnitude for each dimension.
//
// Good luck understanding. Maybe at this point just read the code.


// Each QLearningInstance contains the information we will need at a
// single step of this dynamic programming.
// learning for a particular action.
type QLearningInstance struct {
	// The Q-value we used when deciding to take this action.
	calculatedQ float64

	// The Q-value that would have been ideal, based on the Q-learning
	// rule.
	targetQ float64

	// The difference in probabilities, p_target - p_calc.
	probabilityDifference float64

	// The feature sets that were active for this decision, but not any
	// prior decision.
	newFeatureSets []QFeatureSet
}

func NewQLearningInstance() *QLearningInstance {
	return &QLearningInstance{
		newFeatureSets: []QFeatureSet{EmptyFeatureSet},
	}
}

// AddGradient adds scalar times the gradient to addend, using the
// gradient for the provided color's decisions.
// This uses dynamic programming on a list of QLearningInstances.
func (playout *QPlayout) AddGradient(color Color, scalar float64,
	addend *[NumFeatures]float64) {
	// In activeFeatures we accumulate all features that activate during
	// the game.
	activeFeatures := []QFeature{}

	// The newest instance that is being constructed. This accumulates
	// feature sets.
	instance := NewQLearningInstance()

	// We gather the data we will need to learn, in 'instances'.
	instances := []*QLearningInstance{}

	// This holds the last thing that 'instance' did, for updating
	// target Q.
	var lastInstance *QLearningInstance

	for _, action := range playout.actions {
		// Each new action also is a new feature.
		newFeature := action.Feature()

		// Add a singleton feature set for the new feature.
		instance.newFeatureSets = append(instance.newFeatureSets,
			MakeSingleton(newFeature))

		// Add a feature set for each feature pair you can make with the
		// new feature.
		for _, oldFeature := range activeFeatures {
			instance.newFeatureSets = append(instance.newFeatureSets,
				MakeDoubleton(oldFeature, newFeature))
		}

		// Accumulate features
		activeFeatures = append(activeFeatures, newFeature)

		if action.color == color {
			// 'instance' should apply to this action
			instance.calculatedQ = action.Q
			instances = append(instances, instance)

			// Figure out the target value for the last learning instance
			// from this action's data
			if lastInstance != nil {
				lastInstance.targetQ = action.Q + action.explorationCost
			}

			lastInstance = instance
			instance = NewQLearningInstance()
		}
	}

	if lastInstance == nil {
		// We never made a move this entire playout. We must just be
		// losing immediately to an opponent's move.
		return
	}

	// We accumulated some feature sets in 'instance' but there was
	// never another action we used them for. So we can forget about
	// 'instance' - we won't need to learn things about this feature
	// set. However, we can set the target Q for the lastInstance based
	// on the actual outcome.
	if playout.winner == color {
		lastInstance.targetQ = math.Inf(+1)
	} else {
		lastInstance.targetQ = math.Inf(-1)
	}

	// We do a backward pass to construct the gradient.
	panic("TODO")
}
