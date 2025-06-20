package flow

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/reflect"
	"github.com/thegogod/cmark/tokens"
)

type Expression interface {
	Type() reflect.Type
	Validate(scope *Scope) error
	Evaluate(scope *Scope) (reflect.Value, error)
}

func (self *Flow) ParseExpression(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	node, err := self.parseExpression(parser, NewScanner(ptr))

	if err != nil {
		return nil, err
	}

	return Html(node), nil
}

func (self *Flow) parseExpression(parser html.Parser, scan *Scanner) (Expression, error) {
	return self.parseAssignmentExpression(parser, scan)
}
