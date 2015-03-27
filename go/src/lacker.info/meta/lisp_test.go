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
}
