package main

import (
	"flag"
	"fmt"
	"log"

	"lacker.info/hex"
)

func main() {
	// Load a board position from args.
	// A board in json form should be passed as the first argument.
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		log.Fatal("expected exactly 1 arg to play_hex")
	}
	board := hex.NewBoardFromJSON(args[0])

	// Have a player figure out what move to make on this board.
	player := hex.ShallowRave{500}
	spot := player.Play(board)

	// Print out the move to make.
	fmt.Printf("%s\n", hex.ToJSON(spot))
}
