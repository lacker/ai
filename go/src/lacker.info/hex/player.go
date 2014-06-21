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
	case "sr1k":
		return ShallowRave{1000}
	case "sr2k":
		return ShallowRave{2000}
	case "sr10k":
		return ShallowRave{10000}
	case "sr25k":
		return ShallowRave{25000}
	case "sr100k":
		return ShallowRave{100000}
	case "mcts5":
		return MonteCarloTreeSearch{Seconds:5}
	case "mcts20":
		return MonteCarloTreeSearch{Seconds:20}
	default:
		log.Fatalf("unknown player type: %s", s)
		return nil
	}
}
