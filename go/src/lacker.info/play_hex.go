package main

import (
	"flag"
	"fmt"
	"log"

	"lacker.info/hex"
)

func main() {
	// A board in json form should be passed as the first argument.
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		log.Fatal("expected exactly 1 arg to play_hex")
	}
	board := hex.NewBoardFromJSON(args[0])
	fmt.Printf("%s\n", hex.ToJSON(board))
}
