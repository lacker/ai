package main

import (
	"flag"
	"fmt"
	// "lacker.info/hex"
)

func main() {
	// A board in json form should be passed as the first argument.
	flag.Parse()
	args := flag.Args()
	fmt.Printf("len args = %d\n", len(args))
}
