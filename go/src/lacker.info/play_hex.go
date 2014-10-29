package main

import (
	"flag"
	"log"

	"lacker.info/hex"
)

func main() {
	hex.Seed()

	// Load a board position from args.
	// The first arg should be the player type to play.
	// A board in json form should be passed as the second argument.
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		log.Fatal("expected exactly 2 args to play_hex")
	}
	playerType := args[0]
	boardJSON := args[1]

	hex.PlayForJSON(playerType, boardJSON)
}
