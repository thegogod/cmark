package ast

import "github.com/thegogod/cmark/reflect"

type Node interface {
	Validate(scope *Scope) error
	Evaluate(scope *Scope) (reflect.Value, error)
}
