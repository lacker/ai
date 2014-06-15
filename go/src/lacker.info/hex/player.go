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
	case "shallowrave":
		return ShallowRave{1000}
	default:
		log.Fatalf("unknown player type: %s", s)
		return nil
	}
}
