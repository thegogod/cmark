package flow

import (
	"fmt"
	"strings"

	"github.com/thegogod/cmark/reflect"
)

type SliceExpression struct {
	_type reflect.Type
	items []Expression
}

func (self SliceExpression) Type() reflect.Type {
	return self._type
}

func (self SliceExpression) Validate(scope *Scope) error {
	for _, exp := range self.items {
		if err := exp.Validate(scope); err != nil {
			return err
		}
	}

	return nil
}

func (self SliceExpression) Evaluate(scope *Scope) (reflect.Value, error) {
	slice := reflect.NewSlice(self._type, []reflect.Value{})

	for _, exp := range self.items {
		value, err := exp.Evaluate(scope)

		if err != nil {
			return reflect.NewNil(), err
		}

		slice.Push(value)
	}

	return slice, nil
}

func (self SliceExpression) Print() {
	self.PrintIndent(0, "  ")
}

func (self SliceExpression) PrintIndent(depth int, indent string) {
	fmt.Printf("%s[SliceExpression]: type=\"%s\"\n", strings.Repeat(indent, depth), self._type.Name())

	for _, expression := range self.items {
		expression.PrintIndent(depth+1, indent)
	}
}
