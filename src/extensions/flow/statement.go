package flow

import (
	"github.com/thegogod/cmark/reflect"
)

type Statement interface {
	Validate() error
	Evaluate(scope *Scope) (reflect.Value, error)
}
