package meta

import (
	"strings"
)

// NOTE: this comment reflects an older version of the language with
// "lambda" and "recur".

// Boson is a programming language designed to contain the simplest
// possible way of representing functions.
// It has eight keywords, four "simple" and four "tricky":
// car, cdr, cons, nil
// if, lambda, this, recur

// Boson expressions are typically represented like Lisp lists.
// nil, this take no arguments
// car, cdr, recur, lambda take one argument
// cons takes two arguments
// if takes three arguments

// The funny part is how lambda/recur/this works. It's designed so
// that the programming language needs no variables. For each recur
// and this, what they refer to depends on the closest enclosing
// "lambda". When lambda creates a function, "recur" is bound to that
// function. When it gets called on an argument, "this" gets bound to
// that argument.

// In Boson, functions and data are different. Every value either has
// function type or data type. Every function takes something of data
// type and returns something of data type. So the only way to get a
// function is by a lambda returning it directly. So you can tell from
// an expression which type it has. You cannot cons a function onto
// something else. This essentially prevents functional programming.

// (We may want to change that later.)

// BValues are the data type. They are either like a Lisp list, or
// they can also be an error if we try to car or cdr a nil.
type BValue struct {
	error bool
	list []BValue
}

func (val BValue) String() string {
	if val.error {
		return "error"
	}
	parts := make([]string, len(val.list))
	for i := 0; i < len(val.list); i++ {
		parts[i] = val.list[i].String()
	}

	return "(" + strings.Join(parts, " ") + ")"
}

type BExpression interface {
	String() string
	IsFunctionType() bool
	IsDataType() bool
	IsBound() bool
}
