package hex

/*
Player is an interface for a hex player.
*/

import (
	"log"
)

type Player interface {
	Play(b *NaiveBoard) Spot
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
	case "mcts1":
		return MakeMCTS(1)
	case "mcts5":
		return MakeMCTS(5)
	case "mcts20":
		return MakeMCTS(20)
	case "bleeding":
		mcts := MakeMCTS(5)
		mcts.UseTopoBoards = true
		return mcts
	default:
		log.Fatalf("unknown player type: %s", s)
		return nil
	}
}
