package flow

import (
	"github.com/thegogod/cmark/reflect"
)

type LiteralExpression struct {
	value reflect.Value
}

func (self LiteralExpression) Type() reflect.Type {
	return self.value.Type()
}

func (self LiteralExpression) Validate(scope *Scope) error {
	return nil
}

func (self LiteralExpression) Evaluate(scope *Scope) (reflect.Value, error) {
	return self.value, nil
}
