package flow

import (
	"github.com/thegogod/cmark/reflect"
)

type Expression interface {
	Type() reflect.Type
	Validate() error
	Evaluate(scope *Scope) (reflect.Value, error)
}
