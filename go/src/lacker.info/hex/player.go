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
		return ShallowRave{1}
	case "sr5":
		return ShallowRave{5}
	case "sr20":
		return ShallowRave{20}
	case "classic5":
		return MonteCarloTreeSearch{false, 5}
	case "classic20":
		return MonteCarloTreeSearch{false, 20}
	case "modern5":
		return MonteCarloTreeSearch{true, 5}
	case "modern20":
		return MonteCarloTreeSearch{true, 20}
	default:
		log.Fatalf("unknown player type: %s", s)
		return nil
	}
}
