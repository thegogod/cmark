package flow

import (
	"fmt"

	"github.com/thegogod/cmark/reflect"
	"github.com/thegogod/cmark/tokens"
)

type BinaryExpression struct {
	left  Expression
	op    tokens.Token
	right Expression
}

func (self BinaryExpression) Type() reflect.Type {
	return self.left.Type()
}

func (self BinaryExpression) Validate(scope *Scope) error {
	if err := self.left.Validate(scope); err != nil {
		return err
	}

	if err := self.right.Validate(scope); err != nil {
		return err
	}

	left := self.left.Type()
	right := self.right.Type()

	if !left.Equals(right) {
		return self.op.Error(fmt.Sprintf("expected type '%s', received '%s'", left.Name(), right.Name()))
	}

	return nil
}

func (self BinaryExpression) Evaluate(scope *Scope) (reflect.Value, error) {
	left, err := self.left.Evaluate(scope)

	if err != nil {
		return left, err
	}

	right, err := self.right.Evaluate(scope)

	if err != nil {
		return right, err
	}

	switch self.op.Kind() {
	case EqEq:
		return reflect.NewBool(left.Eq(right)), nil
	case NotEq:
		return reflect.NewBool(!left.Eq(right)), nil
	case Gt:
		return reflect.NewBool(left.Gt(right)), nil
	case GtEq:
		return reflect.NewBool(left.GtEq(right)), nil
	case Lt:
		return reflect.NewBool(left.Lt(right)), nil
	case LtEq:
		return reflect.NewBool(left.LtEq(right)), nil
	case Plus:
		if left.Numeric() {
			return left.Add(right), nil
		}

		return reflect.NewString(left.String() + right.String()), nil
	case Minus:
		return left.Subtract(right), nil
	case Star:
		return left.Multiply(right), nil
	case Slash:
		return left.Divide(right), nil
	default:
		panic("invalid binary expression")
	}
}
