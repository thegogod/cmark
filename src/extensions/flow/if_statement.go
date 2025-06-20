package flow

import (
	"fmt"
	"strings"

	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/reflect"
	"github.com/thegogod/cmark/tokens"
)

func (self *Flow) ParseIfStatement(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	node, err := self.parseIfStatement(parser, NewScanner(ptr))

	if err != nil {
		return nil, err
	}

	return Html(node), nil
}

func (self *Flow) parseIfStatement(parser html.Parser, scan *Scanner) (Statement, error) {
	if _, err := scan.Consume(LeftParen, "expected '('"); err != nil {
		return nil, err
	}

	cond, err := self.parseExpression(parser, scan)

	if err != nil {
		return nil, err
	}

	if _, err = scan.Consume(RightParen, "expected ')'"); err != nil {
		return nil, err
	}

	then, err := self.parseStatement(parser, scan)

	if err != nil {
		return nil, err
	}

	var _else Statement = nil

	if scan.Match(Else) {
		_else, err = self.parseStatement(parser, scan)
	}

	return IfStatement{
		condition: cond,
		then:      then,
		_else:     _else,
	}, nil
}

type IfStatement struct {
	condition Expression
	then      Statement
	_else     Statement
}

func (self IfStatement) Validate(scope *Scope) error {
	if err := self.condition.Validate(scope); err != nil {
		return err
	}

	if err := self.then.Validate(scope); err != nil {
		return err
	}

	if self._else != nil {
		return self._else.Validate(scope)
	}

	return nil
}

func (self IfStatement) Evaluate(scope *Scope) (reflect.Value, error) {
	value, err := self.condition.Evaluate(scope)

	if err != nil {
		return value, err
	}

	if value.Truthy() {
		return self.then.Evaluate(scope)
	} else if self._else != nil {
		return self._else.Evaluate(scope)
	}

	return reflect.NewNil(), nil
}

func (self IfStatement) Print() {
	self.PrintIndent(0, "  ")
}

func (self IfStatement) PrintIndent(depth int, indent string) {
	fmt.Printf("%s[IfStatement]:\n", strings.Repeat(indent, depth))
	self.condition.PrintIndent(depth+1, indent)
	self.then.PrintIndent(depth+1, indent)

	if self._else != nil {
		self._else.PrintIndent(depth+1, indent)
	}
}
