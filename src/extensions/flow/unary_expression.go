package flow

import (
	"github.com/thegogod/cmark/reflect"
	"github.com/thegogod/cmark/tokens"
)

type UnaryExpression struct {
	op    tokens.Token
	right Expression
}

func Unary(op tokens.Token, right Expression) UnaryExpression {
	return UnaryExpression{
		op:    op,
		right: right,
	}
}

func (self UnaryExpression) Type() reflect.Type {
	return self.right.Type()
}

func (self UnaryExpression) Validate() error {
	return self.right.Validate()
}

func (self UnaryExpression) Evaluate(scope *Scope) (reflect.Value, error) {
	right, err := self.right.Evaluate(scope)

	if err != nil {
		return reflect.NewNil(), err
	}

	switch self.op.Kind() {
	case Not:
		return reflect.NewBool(!right.Truthy()), nil
	case Minus:
		right.Decrement()
		return right, nil
	}

	return right, nil
}
