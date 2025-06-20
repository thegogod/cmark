package flow

import (
	"github.com/thegogod/cmark/reflect"
)

type GroupingExpression struct {
	expr Expression
}

func (self GroupingExpression) Type() reflect.Type {
	return self.expr.Type()
}

func (self GroupingExpression) Validate(scope *Scope) error {
	return self.expr.Validate(scope)
}

func (self GroupingExpression) Evaluate(scope *Scope) (reflect.Value, error) {
	return self.expr.Evaluate(scope)
}
