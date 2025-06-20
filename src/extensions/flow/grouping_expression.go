package flow

import (
	"fmt"
	"strings"

	"github.com/thegogod/cmark/reflect"
)

type GroupingExpression struct {
	expr Expression
}

func (self GroupingExpression) Type() reflect.Type {
	return self.expr.Type()
}

func (self GroupingExpression) Validate(scope *Scope) error {
	return self.expr.Validate(scope)
}

func (self GroupingExpression) Evaluate(scope *Scope) (reflect.Value, error) {
	return self.expr.Evaluate(scope)
}

func (self GroupingExpression) Print() {
	self.PrintIndent(0, "  ")
}

func (self GroupingExpression) PrintIndent(depth int, indent string) {
	fmt.Printf("%s[GroupExpression]:\n", strings.Repeat(indent, depth))
	self.expr.PrintIndent(depth+1, indent)
}
