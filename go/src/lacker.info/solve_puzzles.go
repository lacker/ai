package main

import (
	"flag"
	"fmt"
	"log"

	"lacker.info/hex"
)

func main() {
	hex.Seed()

	// Usage:
	//   go run solve_puzzles.go [--debug] playerName puzzleName

	var debugp = flag.Bool("debug", false, "show debugging info")

	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		log.Fatal("usage: go run solve_puzzles.go [--debug] playerName puzzleName")
	}
	playerName := args[0]
	puzzleName := args[1]

	player := hex.GetPlayer(playerName)
	puzzle := hex.GetPuzzle(puzzleName)

	if (*debugp) {
		hex.Debug = true
	}

	spot, odds := player.Play(puzzle.Board)

	// Print out the puzzle
	fmt.Printf("%s\n", puzzle.String)
	fmt.Printf("%s moved %s, estimating odds at %.3f\n\n",
		playerName, hex.ToJSON(spot), odds)
}
