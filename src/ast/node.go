package ast

import "github.com/thegogod/cmark/reflect"

type Visitor interface {
	Visit(node Node) (Visitor, error)
}

type Node interface {
	Validate(scope *Scope) error
	Evaluate(scope *Scope) (reflect.Value, error)
}
