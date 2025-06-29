package flow

import (
	"fmt"
	"strings"

	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/reflect"
	"github.com/thegogod/cmark/tokens"
)

func (self *Flow) ParseAssignmentExpression(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	node, err := self.parseAssignmentExpression(parser, NewScanner(ptr))

	if err != nil {
		return nil, err
	}

	return Html(node), nil
}

func (self *Flow) parseAssignmentExpression(parser html.Parser, scan *Scanner) (Expression, error) {
	expr, err := self.parseOrExpression(parser, scan)

	if err != nil {
		return nil, err
	}

	varType := expr.Type()

	if scan.Match(Eq) {
		value, err := self.parseAssignmentExpression(parser, scan)

		if err != nil {
			return nil, err
		}

		assignType := value.Type()

		if callable, ok := assignType.(reflect.CallableType); ok {
			assignType = callable.ReturnType()
		}

		if !varType.Equals(assignType) {
			return nil, scan.Prev().Error(fmt.Sprintf(
				"expected type '%s', received '%s'",
				varType.Name(),
				assignType.Name(),
			))
		}

		switch v := expr.(type) {
		case VariableExpression:
			if !self.scope.Has(v.name.String()) {
				return nil, scan.Prev().Error("undefined identifier")
			}

			return AssignExpression{
				name:  v.name,
				value: value,
			}, nil
		case GetExpression:
			if !self.scope.Has(v.name.String()) {
				return nil, scan.Prev().Error("undefined identifier")
			}

			return SetExpression{
				object: v.object,
				name:   v.name,
				value:  value,
			}, nil
		}

		return nil, scan.Prev().Error("invalid assignment target")
	}

	return expr, nil
}

type AssignExpression struct {
	name  tokens.Token
	value Expression
}

func (self AssignExpression) Type() reflect.Type {
	return self.value.Type()
}

func (self AssignExpression) Validate(scope *Scope) error {
	return self.value.Validate(scope)
}

func (self AssignExpression) Evaluate(scope *Scope) (reflect.Value, error) {
	value, err := self.value.Evaluate(scope)

	if err != nil {
		return value, err
	}

	scope.Set(self.name.String(), &ScopeEntry{Value: value})
	return value, nil
}

func (self AssignExpression) Print() {
	self.PrintIndent(0, "  ")
}

func (self AssignExpression) PrintIndent(depth int, indent string) {
	fmt.Printf("%s[AssignExpression]: name=\"%s\"\n", strings.Repeat(indent, depth), self.name.String())
	self.value.PrintIndent(depth+1, indent)
}
