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
	default:
		log.Fatalf("unknown player type: %s", s)
		return nil
	}
}
