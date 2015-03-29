package main

import (
	"flag"
	"log"

	"lacker.info/meta"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) > 0 {
		log.Fatal("run_meta takes no args")
	}

	meta.Main()
}
