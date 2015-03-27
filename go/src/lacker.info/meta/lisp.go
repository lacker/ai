package meta

import (
	"log"
)

// A Lisp toolkit.
// See http://norvig.com/lispy.html

func Test() {
	log.Printf("%v", tokenize("((arf bard (+ 3 six)) ())"))
}

func tokenize(s string) []string {
	panic("TODO: implement me")
}
