package meta

import (
	"log"
	"strings"
)

// A Lisp toolkit.
// See http://norvig.com/lispy.html

func Test() {
	log.Printf("%#v", tokenize("((arf bard (+  3 six)) ())"))
}

func tokenize(s string) []string {
	s = strings.Replace(s, "(", " ( ", -1)
	s = strings.Replace(s, ")", " ) ", -1)
	return strings.Fields(s)
}
