package assert

import (
	"fmt"
	"reflect"
)

type EqualExpression struct {
	actual any
}

func Equal(actual any) *EqualExpression {
	return &EqualExpression{actual}
}

func (self EqualExpression) Evaluate(value any) error {
	if !reflect.DeepEqual(value, self.actual) {
		return fmt.Errorf("expected: %v, received: %v", self.actual, value)
	}

	return nil
}
