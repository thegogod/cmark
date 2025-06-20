package flow

import (
	"fmt"
	"strings"

	"github.com/thegogod/cmark/reflect"
	"github.com/thegogod/cmark/tokens"
)

type LogicalExpression struct {
	left  Expression
	op    tokens.Token
	right Expression
}

func (self LogicalExpression) Type() reflect.Type {
	return self.left.Type()
}

func (self LogicalExpression) Validate(scope *Scope) error {
	if err := self.left.Validate(scope); err != nil {
		return err
	}

	if err := self.right.Validate(scope); err != nil {
		return err
	}

	left := self.left.Type()
	right := self.right.Type()

	if !left.Equals(right) {
		return self.op.Error(fmt.Sprintf("expected type '%s', received '%s'", left.Name(), right.Name()))
	}

	return nil
}

func (self LogicalExpression) Evaluate(scope *Scope) (reflect.Value, error) {
	left, err := self.left.Evaluate(scope)

	if err != nil {
		return reflect.NewNil(), err
	}

	if self.op.Kind() == Or {
		if left.Truthy() {
			return left, nil
		}
	} else {
		if !left.Truthy() {
			return left, nil
		}
	}

	return self.right.Evaluate(scope)
}

func (self LogicalExpression) Print() {
	self.PrintIndent(0, "  ")
}

func (self LogicalExpression) PrintIndent(depth int, indent string) {
	fmt.Printf("%s[LogicalExpression]: op=\"%s\"\n", strings.Repeat(indent, depth), self.op.String())
	self.left.PrintIndent(depth+1, indent)
	self.right.PrintIndent(depth+1, indent)
}
