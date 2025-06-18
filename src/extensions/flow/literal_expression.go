package flow

import (
	"github.com/thegogod/cmark/reflect"
)

type LiteralExpression struct {
	value reflect.Value
}

func Literal(value reflect.Value) LiteralExpression {
	return LiteralExpression{value: value}
}

func (self LiteralExpression) Type() reflect.Type {
	return self.value.Type()
}

func (self LiteralExpression) Validate() error {
	return nil
}

func (self LiteralExpression) Evaluate(scope *Scope) (reflect.Value, error) {
	return self.value, nil
}
