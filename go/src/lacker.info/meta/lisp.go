package meta

import (
	"log"
	"strings"
)

// A Lisp toolkit.
// See http://norvig.com/lispy.html

// Either a List or an Atom.
type SExpression interface {
}

type List struct {
	list []SExpression
}

type Atom struct {
	atom string
}

// Turns a list of tokens (from tokenize) into an SExpression.
// Starts at the given index;
// Returns the SExpression and the index of the first unread token.
func readFromTokens(tokens []string, index int) (SExpression, int) {
	if len(tokens) <= index {
		log.Fatalf("only %d tokens but need to read tokens[%d]",
			len(tokens), index)
	}
	token := tokens[index]
	index++

	if token == "(" {
		panic("TODO")
	}
}

// Turns a string into a list of Lisp tokens.
// White space and parentheses are the only separators.
func tokenize(s string) []string {
	s = strings.Replace(s, "(", " ( ", -1)
	s = strings.Replace(s, ")", " ) ", -1)
	return strings.Fields(s)
}

// This is just whatever run_meta runs. Feel free to muck around.
func Main() {
	log.Printf("%#v", tokenize("((arf bard (+  3 six)) ())"))
}

