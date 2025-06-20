package flow

import (
	"fmt"
	"strings"

	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/reflect"
	"github.com/thegogod/cmark/tokens"
)

func (self *Flow) ParseUnaryExpression(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	node, err := self.parseUnaryExpression(parser, NewScanner(ptr))

	if err != nil {
		return nil, err
	}

	return Html(node), nil
}

func (self *Flow) parseUnaryExpression(parser html.Parser, scan *Scanner) (Expression, error) {
	if scan.Match(Not) || scan.Match(Minus) {
		op := scan.Prev()
		right, err := self.parseUnaryExpression(parser, scan)

		if err != nil {
			return nil, err
		}

		return UnaryExpression{op, right}, nil
	}

	return self.parseCallExpression(parser, scan)
}

type UnaryExpression struct {
	op    tokens.Token
	right Expression
}

func (self UnaryExpression) Type() reflect.Type {
	return self.right.Type()
}

func (self UnaryExpression) Validate(scope *Scope) error {
	return self.right.Validate(scope)
}

func (self UnaryExpression) Evaluate(scope *Scope) (reflect.Value, error) {
	right, err := self.right.Evaluate(scope)

	if err != nil {
		return reflect.NewNil(), err
	}

	switch self.op.Kind() {
	case Not:
		return reflect.NewBool(!right.Truthy()), nil
	case Minus:
		right.Decrement()
		return right, nil
	}

	return right, nil
}

func (self UnaryExpression) Print() {
	self.PrintIndent(0, "  ")
}

func (self UnaryExpression) PrintIndent(depth int, indent string) {
	fmt.Printf("%s[UnaryExpression]: op=\"%s\"\n", strings.Repeat(indent, depth), self.op.String())
	self.right.PrintIndent(depth+1, indent)
}
