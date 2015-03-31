package meta

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// A Lisp toolkit.
// See http://norvig.com/lispy.html

// A List, Symbol, Integer, Function, or Error.
type SExpression interface {
	String() string
	Eval(env *Environment) SExpression
}

type List struct {
	list []SExpression
}

type Symbol struct {
	symbol string
}

type Integer int

type Function struct {
	macro bool
	function func([]SExpression) SExpression
}

type Error struct {
	error string
}

func (list List) String() string {
	parts := make([]string, len(list.list))
	for i := 0; i < len(list.list); i++ {
		parts[i] = list.list[i].String()
	}
	return "(" + strings.Join(parts, " ") + ")"
}

func (list List) Eval(env *Environment) SExpression {
	if len(list.list) == 0 {
		return Error{error:"cannot eval empty list"}
	}
	f := list.list[0].Eval(env).(Function)
	rawArgs := list.list[1:]
	if f.macro {
		return f.function(rawArgs, env)
	}
	args := make([]SExpression, len(rawArgs))
	for i := 0; i < len(rawArgs); i++ {
		args[i] = rawArgs[i].Eval(env)
	}
	return f.function(args, env)
}

func (symbol Symbol) String() string {
	return symbol.symbol
}

func (symbol Symbol) Eval(env *Environment) SExpression {
	return env.Get(symbol.symbol)
}

func (i Integer) String() string {
	return fmt.Sprintf("%d", i)
}

func (i Integer) Eval(env *Environment) SExpression {
	return i
}

func (f Function) String() string {
	if f.macro {
		return "<macro>"
	}
	return "<func>"
}

func (f Function) Eval(env *Environment) SExpression {
	return Error{error:"functions cannot be eval'd on their own"}
}

func (e Error) String() string {
	return fmt.Sprintf("error(%s)", e.error)
}

func (e Error) Eval(env *Environment) SExpression {
	return e
}

// Some implementations for built-ins

func QuoteMacro(args []SExpression) SExpression {
	if len(args) != 1 {
		return Error{error:"quote must have exactly one argument"}
	}
	return args[0]
}

// TODO: more builtins. Does "define" need to live in eval?

// The way we track what variables refer to
type Environment struct {
	parent *Environment
	content map[string]*SExpression
}

func (env *Environment) Get(s string) SExpression {
	answer := env.content[s]
	if answer != nil {
		return *answer
	}
	if env.parent == nil {
		log.Fatalf("could not dereference '%s'", s)
	}
	return env.parent.Get(s)
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

	i, err := strconv.Atoi(token)
	if err != nil {
		return Symbol{symbol:token}
	}

	return Integer(i)
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

