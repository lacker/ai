package meta

import (
	"log"
	"strings"
)

// A Lisp toolkit.
// See http://norvig.com/lispy.html

// Either a List or an Atom.
type SExpression interface {
	String() string
}

type List struct {
	list []SExpression
}

type Atom struct {
	atom string
}

func (list List) String() string {
	parts := make([]string, len(list.list))
	for i := 0; i < len(list.list); i++ {
		parts[i] = list.list[i].String()
	}
	return "(" + strings.Join(parts, " ") + ")"
}

func (atom Atom) String() string {
	return atom.atom
}

// Turns a list of tokens (from tokenize) into an SExpression.
// Starts at the provided index and moves it along.
func readFromTokensAtIndex(tokens []string, index *int) SExpression {
	if len(tokens) <= *index {
		log.Fatalf("only %d tokens but need to read tokens[%d]",
			len(tokens), *index)
	}
	token := tokens[*index]
	*index++

	if token == "(" {
		list := make([]SExpression, 0)
		for tokens[*index] != ")" {
			sexp := readFromTokensAtIndex(tokens, index)
			list = append(list, sexp)
		}
		*index++ // pop the ")"

		return List{list:list}
	}

	if token == ")" {
		log.Fatalf("unexpected ) at index %d", *index)
	}

	return Atom{atom:token}
}

// Turns a list of tokens (from tokenize) into an SExpression.
func readFromTokens(tokens []string) SExpression {
	var index int = 0
	answer := readFromTokensAtIndex(tokens, &index)
	if index != len(tokens) {
		log.Fatalf("we have %d tokens but only used %d of them",
			len(tokens), index)
	}
	return answer
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

