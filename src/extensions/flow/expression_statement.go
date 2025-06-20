package flow

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/reflect"
	"github.com/thegogod/cmark/tokens"
)

func (self *Flow) ParseExpressionStatement(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	node, err := self.parseExpressionStatement(parser, NewScanner(ptr))

	if err != nil {
		return nil, err
	}

	return Html(node), nil
}

func (self *Flow) parseExpressionStatement(parser html.Parser, scan *Scanner) (Statement, error) {
	expr, err := self.parseExpression(parser, scan)

	if err != nil {
		return nil, err
	}

	if _, err = scan.Consume(SemiColon, "expected ';'"); err != nil {
		return nil, err
	}

	return ExpressionStatement{expr}, nil
}

type ExpressionStatement struct {
	expression Expression
}

func (self ExpressionStatement) Validate(scope *Scope) error {
	return self.expression.Validate(scope)
}

func (self ExpressionStatement) Evaluate(scope *Scope) (reflect.Value, error) {
	_, err := self.expression.Evaluate(scope)
	return reflect.NewNil(), err
}
