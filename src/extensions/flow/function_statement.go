package flow

import (
	"github.com/thegogod/cmark/reflect"
	"github.com/thegogod/cmark/tokens"
)

type FunctionStatement struct {
	name       tokens.Token
	params     []VariableStatement
	returnType reflect.Type
	body       []Statement
}

func (self FunctionStatement) Validate(scope *Scope) error {
	for _, stmt := range self.body {
		err := stmt.Validate(scope)

		if err != nil {
			return err
		}
	}

	return nil
}

func (self FunctionStatement) Evaluate(scope *Scope) (reflect.Value, error) {
	for _, stmt := range self.body {
		v, err := stmt.Evaluate(scope)

		if err != nil {
			return reflect.NewNil(), err
		}

		if v.IsNil() {
			return v, nil
		}
	}

	return reflect.NewNil(), nil
}
