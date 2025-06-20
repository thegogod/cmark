package flow

import (
	"fmt"

	"github.com/thegogod/cmark/reflect"
	"github.com/thegogod/cmark/tokens"
)

type SelfExpression struct {
	keyword tokens.Token
	_type   reflect.Type
}

func (self SelfExpression) Type() reflect.Type {
	return self._type
}

func (self SelfExpression) Validate(scope *Scope) error {
	if !scope.Has("self") {
		return fmt.Errorf("identifier 'self' not found")
	}

	return nil
}

func (self SelfExpression) Evaluate(scope *Scope) (reflect.Value, error) {
	entry := scope.GetLocal("self")
	return entry.Value, nil
}
