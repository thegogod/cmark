package flow

import (
	"github.com/thegogod/cmark/reflect"
)

type IfStatement struct {
	condition Expression
	then      Statement
	_else     Statement
}

func (self IfStatement) Validate() error {
	if err := self.condition.Validate(); err != nil {
		return err
	}

	if err := self.then.Validate(); err != nil {
		return err
	}

	if self._else != nil {
		return self._else.Validate()
	}

	return nil
}

func (self IfStatement) Evaluate(scope *Scope) (reflect.Value, error) {
	value, err := self.condition.Evaluate(scope)

	if err != nil {
		return reflect.NewNil(), err
	}

	if value.Truthy() {
		return self.then.Evaluate(scope)
	} else if self._else != nil {
		return self._else.Evaluate(scope)
	}

	return reflect.NewNil(), nil
}
