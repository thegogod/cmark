package flow

import (
	"fmt"
	"strings"

	"github.com/thegogod/cmark/reflect"
)

type LiteralExpression struct {
	value reflect.Value
}

func (self LiteralExpression) Type() reflect.Type {
	return self.value.Type()
}

func (self LiteralExpression) Validate(scope *Scope) error {
	return nil
}

func (self LiteralExpression) Evaluate(scope *Scope) (reflect.Value, error) {
	return self.value, nil
}

func (self LiteralExpression) Print() {
	self.PrintIndent(0, "  ")
}

func (self LiteralExpression) PrintIndent(depth int, indent string) {
	fmt.Printf("%s[LiteralExpression]: value=%v\n", strings.Repeat(indent, depth), self.value.Any())
}
