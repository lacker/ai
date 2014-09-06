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
	//   go run solve_puzzles.go puzzlename playername
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		log.Fatal("expected exactly 2 args to solve_puzzles")
	}
	puzzleName := args[0]
	playerName := args[1]

	player := hex.GetPlayer(playerName)
	puzzle := hex.GetPuzzle(puzzleName)
	spot, odds := player.Play(puzzle.Board)

	// Print out the puzzle
	fmt.Printf("%s\n", puzzle.String)
	fmt.Printf("%s moved %s, estimating odds at %.3f\n\n",
		playerName, hex.ToJSON(spot), odds)
}
