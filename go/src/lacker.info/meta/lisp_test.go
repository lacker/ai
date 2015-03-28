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
