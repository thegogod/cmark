package flow

import (
	"fmt"

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
	entry := scope.Get(self.name.String())

	if entry.Kind == TypeScope {
		return reflect.NewNil(), fmt.Errorf("cannot reference type '%s' as a value", self.name.String())
	}

	return reflect.ValueOf(entry.Value), nil
}
