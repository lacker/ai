package meta

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// A Lisp toolkit.
// See http://norvig.com/lispy.html

// A List, Symbol, Integer, Function, or Error.
type SExpression interface {
	String() string
	Eval(env *Environment) SExpression

	// Only zero and empty list are falsy
	Truthy() bool
}

type List struct {
	list []SExpression
}

type Symbol struct {
	symbol string
}

type Integer int

type Function struct {
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

	// Handle builtin macros

	rawFirst := list.list[0]
	switch rawFirst.(type) {
	case Symbol:
		switch rawFirst.(Symbol).symbol {
		case "quote":
			if len(list.list) != 2 {
				return Error{error:"quote must have exactly one argument"}
			}
			return list.list[1]
		case "if":
			if len(list.list) != 4 {
				return Error{error:"if must have exactly three arguments"}
			}
			if list.list[1].Eval(env).Truthy() {
				return list.list[2].Eval(env)
			} else {
				return list.list[3].Eval(env)
			}
		case "define":
			if len(list.list) != 3 {
				return Error{error:"define must have exactly two arguments"}
			}
			value := list.list[2].Eval(env)
			sym := list.list[1]
			switch sym.(type) {
			case Symbol:
				env.Set(sym.(Symbol).symbol, value)
				return value
			default:
				return Error{error:"define's first arg must be a symbol"}
			}
		}
	}

	// If it's not a builtin macro then it must be a function

	first := rawFirst.Eval(env)
	switch first.(type) {
	case Function:
		// This is assumed
	default:
		return Error{error:"first value in a list must be a function"}
	}
	f := first.(Function)
	rawArgs := list.list[1:]
	args := make([]SExpression, len(rawArgs))
	for i := 0; i < len(rawArgs); i++ {
		args[i] = rawArgs[i].Eval(env)
	}
	return f.function(args)
}

func (list List) Truthy() bool {
	return len(list.list) > 0
}

func (symbol Symbol) String() string {
	return symbol.symbol
}

func (symbol Symbol) Eval(env *Environment) SExpression {
	return env.Get(symbol.symbol)
}

func (symbol Symbol) Truthy() bool {
	return true
}

func (i Integer) String() string {
	return fmt.Sprintf("%d", i)
}

func (i Integer) Eval(env *Environment) SExpression {
	return i
}

func (i Integer) Truthy() bool {
	return i != 0
}

func (f Function) String() string {
	return "<func>"
}

func (f Function) Eval(env *Environment) SExpression {
	return Error{error:"functions cannot be eval'd on their own"}
}

func (f Function) Truthy() bool {
	return true
}

func (e Error) String() string {
	return fmt.Sprintf("error(%s)", e.error)
}

func (e Error) Eval(env *Environment) SExpression {
	return e
}

func (e Error) Truthy() bool {
	return true
}

func MakeIntFunction(f func([]Integer) SExpression) Function {
	wrapped := func(args []SExpression) SExpression {
		ints := make([]Integer, len(args))
		for i := 0; i < len(args); i++ {
			switch args[i].(type) {
			case Integer:
				ints[i] = args[i].(Integer)
			default:
				return Error{error:"function requires Integer args"}
			}
		}
		return f(ints)
	}
	return Function{function:wrapped}
}

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
		return Error{error:fmt.Sprintf("could not dereference %s", s)}
	}
	return env.parent.Get(s)
}

func (env *Environment) Set(s string, val SExpression) {
	env.content[s] = &val
}

func EmptyEnvironment() *Environment {
	return &Environment{
		content: make(map[string]*SExpression),
	}
}

func DefaultEnvironment() *Environment {
	env := EmptyEnvironment()
	env.Set("+", MakeIntFunction(func(ints []Integer) SExpression {
		sum := Integer(0)
		for i := 0; i < len(ints); i++ {
			sum += ints[i]
		}
		return sum
	}))
	env.Set("*", MakeIntFunction(func(ints []Integer) SExpression {
		prod := Integer(1)
		for i := 0; i < len(ints); i++ {
			prod *= ints[i]
		}
		return prod
	}))
	env.Set("<", MakeIntFunction(func(ints []Integer) SExpression {
		if len(ints) != 2 {
			return Error{error:"< expects two args"}
		}
		if (ints[0] < ints[1]) {
			return Integer(1)
		} else {
			return Integer(0)
		}
	}))
	env.Set(">", MakeIntFunction(func(ints []Integer) SExpression {
		if len(ints) != 2 {
			return Error{error:"> expects two args"}
		}
		if (ints[0] > ints[1]) {
			return Integer(1)
		} else {
			return Integer(0)
		}
	}))
	return env
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
		for {
			if len(tokens) <= *index {
				log.Fatalf("ran off the end of tokens")
			}
			if tokens[*index] == ")" {
				break
			}
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

func read(s string) SExpression {
	return readFromTokens(tokenize(s))
}

// Turns a string into a list of Lisp tokens.
// White space and parentheses are the only separators.
func tokenize(s string) []string {
	s = strings.Replace(s, "(", " ( ", -1)
	s = strings.Replace(s, ")", " ) ", -1)
	return strings.Fields(s)
}

// Runs a REPL
func Main() {
	env := DefaultEnvironment()
	for {
		// Show a prompt
		fmt.Printf("> ")

		// Read a line
		bio := bufio.NewReader(os.Stdin)
		line, hasMoreInLine, err := bio.ReadLine()
		if hasMoreInLine || err != nil {
			// This happens on a control-D
			fmt.Printf("\n")
			break
		}
		
		// Evaluate it
		s := read(string(line))
		out := s.Eval(env).String()

		// Print the result
		fmt.Printf("%s\n", out)
	}
}

