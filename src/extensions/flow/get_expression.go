package flow

import (
	"fmt"
	"strings"

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

func (self GetExpression) Print() {
	self.PrintIndent(0, "  ")
}

func (self GetExpression) PrintIndent(depth int, indent string) {
	fmt.Printf("%s[GetExpression]: name=\"%s\"\n", strings.Repeat(indent, depth), self.name.String())
	self.object.PrintIndent(depth+1, indent)
}
