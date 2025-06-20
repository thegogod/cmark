package flow

import (
	"fmt"
	"strings"

	"github.com/thegogod/cmark/reflect"
	"github.com/thegogod/cmark/tokens"
)

type VariableExpression struct {
	name  tokens.Token
	_type reflect.Type
}

func (self VariableExpression) Type() reflect.Type {
	return self._type
}

func (self VariableExpression) Validate(scope *Scope) error {
	if !scope.Has(self.name.String()) {
		return fmt.Errorf("identifier '%s' not found", self.name.String())
	}

	return nil
}

func (self VariableExpression) Evaluate(scope *Scope) (reflect.Value, error) {
	if err := self.Validate(scope); err != nil {
		return reflect.NewNil(), err
	}

	entry := scope.Get(self.name.String())

	if entry.Value.IsNil() {
		return reflect.NewNil(), fmt.Errorf("cannot reference type '%s' as a value", self.name.String())
	}

	return entry.Value, nil
}

func (self VariableExpression) Print() {
	self.PrintIndent(0, "  ")
}

func (self VariableExpression) PrintIndent(depth int, indent string) {
	fmt.Printf("%s[VariableExpression]: name=\"%s\" type=\"%s\"\n", strings.Repeat(indent, depth), self.name.String(), self._type.Name())
}
