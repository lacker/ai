package hex

/*
Player is an interface for a hex player.
*/

import (
	"fmt"
	"log"
)

type Player interface {
	// Returns the best move and an expected win rate
	Play(b Board) (NaiveSpot, float64)
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
	case "topo5":
		mcts := MakeMCTS(5)
		mcts.UseTopoBoards = true
		return mcts
	case "mcts1":
		return MakeMCTS(1)
	case "mcts5":
		return MakeMCTS(5)
	case "mcts20":
		return MakeMCTS(20)
	case "ss5":
		return SpotSorter{Seconds:5, Quiet:false}
	case "mf5":
		return MetaFarmer{Seconds:5, Quiet:false, QuickType:"democracy"}
	case "dn5":
		return MetaFarmer{Seconds:5, Quiet:false, QuickType:"deltanet"}
	case "qt":
		return &QTrainer{Seconds:5, Quiet:false}
	default:
		log.Fatalf("unknown player type: %s", s)
		return nil
	}
}

// Helper for script
func PlayForJSON(playerType string, boardJSON string) {
	board := NewNaiveBoardFromJSON(boardJSON)

	// Have a player figure out what move to make on this board.
	player := GetPlayer(playerType)
	spot, _ := player.Play(board)

	// Print out the move to make.
	fmt.Printf("{\"Row\":%d,\"Col\":%d}\n", spot.Row(), spot.Col())
}
