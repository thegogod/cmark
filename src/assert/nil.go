package assert

import "fmt"

type NilExpression struct{}

func Nil() NilExpression {
	return NilExpression{}
}

func (self NilExpression) Evaluate(value any) error {
	if value != nil {
		return fmt.Errorf("expected: nil, received: %v", value)
	}

	return nil
}
