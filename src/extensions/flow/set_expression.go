package flow

import (
	"fmt"
	"strings"

	"github.com/thegogod/cmark/reflect"
	"github.com/thegogod/cmark/tokens"
)

type SetExpression struct {
	object Expression
	name   tokens.Token
	value  Expression
}

func (self SetExpression) Type() reflect.Type {
	return self.object.Type()
}

func (self SetExpression) Validate(scope *Scope) error {
	if err := self.object.Validate(scope); err != nil {
		return err
	}

	if err := self.value.Validate(scope); err != nil {
		return err
	}

	objectType := self.object.Type()
	valueType := self.value.Type()

	if !objectType.Equals(valueType) {
		return fmt.Errorf("expected type '%s', received '%s'", objectType.Name(), valueType.Name())
	}

	return nil
}

func (self SetExpression) Evaluate(scope *Scope) (reflect.Value, error) {
	return reflect.NewNil(), nil
}

func (self SetExpression) Print() {
	self.PrintIndent(0, "  ")
}

func (self SetExpression) PrintIndent(depth int, indent string) {
	fmt.Printf("%s[SetExpression]: name=\"%s\"\n", strings.Repeat(indent, depth), self.name.String())
	self.object.PrintIndent(depth+1, indent)
	self.value.PrintIndent(depth+1, indent)
}
