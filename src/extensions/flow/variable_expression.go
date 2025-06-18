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

func VarRef(name tokens.Token, _type reflect.Type) VariableExpression {
	return VariableExpression{name, _type}
}

func (self VariableExpression) Type() reflect.Type {
	return self._type
}

func (self VariableExpression) Validate() error {
	return nil
}

func (self VariableExpression) Evaluate(scope *Scope) (reflect.Value, error) {
	if !scope.Has(self.name.String()) {
		return reflect.NewNil(), fmt.Errorf("identifier '%s' not found", self.name.String())
	}

	entry := scope.Get(self.name.String())

	if entry.Kind == TypeScope {
		return reflect.NewNil(), fmt.Errorf("cannot reference type '%s' as a value", self.name.String())
	}

	return reflect.ValueOf(entry.Value), nil
}
