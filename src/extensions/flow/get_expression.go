package flow

import (
	"fmt"

	"github.com/thegogod/cmark/reflect"
	"github.com/thegogod/cmark/tokens"
)

type GetExpression struct {
	object Expression
	name   tokens.Token
}

func (self GetExpression) Type() reflect.Type {
	t := self.object.Type()
	return t.GetMember(self.name.String())
}

func (self GetExpression) Validate(scope *Scope) error {
	if err := self.object.Validate(scope); err != nil {
		return err
	}

	if !self.object.Type().HasMember(self.name.String()) {
		return fmt.Errorf("member '%s' not found", self.name.String())
	}

	return nil
}

func (self GetExpression) Evaluate(scope *Scope) (reflect.Value, error) {
	value, err := self.object.Evaluate(scope)

	if err != nil {
		return value, err
	}

	return value.GetMember(self.name.String()), nil
}
