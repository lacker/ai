package meta

import (
	"testing"
)

func TestTokenize(t *testing.T) {
	tokens := tokenize("((arf bard (+  3 six)) ())")
	AssertEq("(", tokens[0])
	AssertEq("(", tokens[1])
	AssertEq("arf", tokens[2])
	AssertEq("bard", tokens[3])
	AssertEq("(", tokens[4])
	AssertEq("+", tokens[5])
	AssertEq("3", tokens[6])
	AssertEq("six", tokens[7])
	AssertEq(")", tokens[8])
	AssertEq(")", tokens[9])
	AssertEq("(", tokens[10])
	AssertEq(")", tokens[11])
	AssertEq(")", tokens[12])
}

func TestReadFromTokens(t *testing.T) {
	s := read("((arf bard (+  3 six)) ())")
	AssertEq("((arf bard (+ 3 six)) ())", s.String())
}

func TestBuiltInFunction(t *testing.T) {
	env := DefaultEnvironment()
	s := read("(+ 2 2)")
	AssertEq("4", s.Eval(env).String())
}

func TestMultiFunction(t *testing.T) {
	env := DefaultEnvironment()
	s := read("(* (+ 1 2 3) (+ 4 5 6) (+ 7 8 9))")
	AssertEq("2160", s.Eval(env).String())
}

func TestQuote(t *testing.T) {
	env := DefaultEnvironment()
	s := read("(quote (+ 1 2 3))")
	AssertEq("(+ 1 2 3)", s.Eval(env).String())
}

func TestIf(t *testing.T) {
	env := DefaultEnvironment()
	s := read("(if (< 1 2) 3 4)")
	AssertEq("3", s.Eval(env).String())
}
