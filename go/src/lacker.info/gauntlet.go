package main

// Tests the performance of a hex bot on a particular set of challenges.
// This is useful as a sort of integration test to efficiently check
// that overall performance has remained high after some tweak.

import (
	"flag"
	"log"

	"lacker.info/hex"
)

func main() {
	hex.Seed()

	flag.Parse()
	args := flag.Args()
	if len(args) > 1 {
		log.Fatal("expected at most 1 arg to gauntlet")
	}

	panic("TODO: implement")
}
