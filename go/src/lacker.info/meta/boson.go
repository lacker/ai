package meta

import (
)

// Boson is a programming language designed to contain the simplest
// possible way of representing functions. Even simpler than Lisp.
// TODO: is this description correct in light of lambda calculus?

type BExpression interface {
	String() string
}

// The actual content is always "false", and not used.
type Nil bool

func (n Nil) String() string {
	return "nil"
}

func MakeNil() Nil {
	return Nil(false)
}

// The main data type.
type Cons struct {
	car BExpression
	cdr BExpression
}

func (c Cons) String() string {
	return fmt.Sprintf("(cons %s %s)", c.car.String(), c.cdr.String())
}

func MakeCons(car BExpression, cdr BExpression) Cons {
	return Cons{car:car, cdr:cdr}
}

// An unbound variable.
type Variable struct {
	name string
}

func (v Variable) String() string {
	return v.name
}

func MakeX() Variable {
	return Variable{name:"x"}
}
