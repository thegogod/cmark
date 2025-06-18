package flow

import (
	"github.com/thegogod/cmark/reflect"
)

type ExpressionStatement struct {
	expression Expression
}

func Expr(expression Expression) ExpressionStatement {
	return ExpressionStatement{expression}
}

func (self ExpressionStatement) Validate() error {
	return self.expression.Validate()
}

func (self ExpressionStatement) Evaluate(scope *Scope) (reflect.Value, error) {
	_, err := self.expression.Evaluate(scope)
	return reflect.NewNil(), err
}

func (self ExpressionStatement) Render(scope *Scope) []byte {
	self.expression.Evaluate(scope)
	return []byte{}
}

func (self ExpressionStatement) RenderPretty(scope *Scope, indent string) []byte {
	self.expression.Evaluate(scope)
	return []byte{}
}
