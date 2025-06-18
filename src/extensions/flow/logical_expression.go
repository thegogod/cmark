package flow

import (
	"fmt"

	"github.com/thegogod/cmark/reflect"
	"github.com/thegogod/cmark/tokens"
)

type LogicalExpression struct {
	left  Expression
	op    tokens.Token
	right Expression
}

func Logical(left Expression, op tokens.Token, right Expression) LogicalExpression {
	return LogicalExpression{
		left:  left,
		op:    op,
		right: right,
	}
}

func (self LogicalExpression) Type() reflect.Type {
	return self.left.Type()
}

func (self LogicalExpression) Validate() error {
	if err := self.left.Validate(); err != nil {
		return err
	}

	if err := self.right.Validate(); err != nil {
		return err
	}

	left := self.left.Type()
	right := self.right.Type()

	if !left.Equals(right) {
		return self.op.Error(fmt.Sprintf("expected type '%s', received '%s'", left.Name(), right.Name()))
	}

	return nil
}

func (self LogicalExpression) Evaluate(scope *Scope) (reflect.Value, error) {
	left, err := self.left.Evaluate(scope)

	if err != nil {
		return reflect.NewNil(), err
	}

	if self.op.Kind() == Or {
		if left.Truthy() {
			return left, nil
		}
	} else {
		if !left.Truthy() {
			return left, nil
		}
	}

	return self.right.Evaluate(scope)
}
