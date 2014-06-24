package hex

/*
Player is an interface for a hex player.
*/

import (
	"log"
)

type Player interface {
	Play(b *Board) Spot
}

func GetPlayer(s string) Player {
	switch s {
	case "random":
		return Random{}
	case "sr1":
		return ShallowRave{Seconds:1, Quiet:false}
	case "sr5":
		return ShallowRave{Seconds:5, Quiet:false}
	case "sr20":
		return ShallowRave{Seconds:20, Quiet:false}
	case "uct5":
		return PureUCT{5}
	case "uct20":
		return PureUCT{20}
	case "mcts5":
		return MonteCarloTreeSearch{Seconds: 5, Quiet: false, V: 1000}
	case "mcts20":
		return MonteCarloTreeSearch{Seconds: 5, Quiet: false, V: 1000}
	case "bleeding":
		return MonteCarloTreeSearch{Seconds: 5, Quiet: false, V: 2000}
	default:
		log.Fatalf("unknown player type: %s", s)
		return nil
	}
}
