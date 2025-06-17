package assert

import (
	"fmt"
	"reflect"
)

type LenExpression struct {
	length int
}

func Len(length int) *LenExpression {
	return &LenExpression{length}
}

func (self LenExpression) Evaluate(value any) error {
	v := reflect.ValueOf(value)
	kind := v.Kind()

	if kind != reflect.Array &&
		kind != reflect.Slice &&
		kind != reflect.String &&
		kind != reflect.Chan &&
		kind != reflect.Map {
		return fmt.Errorf("expected iterable type, received '%s'", kind.String())
	}

	if v.Len() != self.length {
		return fmt.Errorf("expected length %d, received %d", self.length, v.Len())
	}

	return nil
}
