package flow

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/reflect"
	"github.com/thegogod/cmark/tokens"
)

type Statement interface {
	Validate(scope *Scope) error
	Evaluate(scope *Scope) (reflect.Value, error)
	Print()
	PrintIndent(depth int, indent string)
}

func (self *Flow) ParseStatement(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	node, err := self.parseStatement(parser, NewScanner(ptr))

	if err != nil {
		return nil, err
	}

	return Html(node), nil
}

func (self *Flow) parseStatement(parser html.Parser, scan *Scanner) (Statement, error) {
	if scan.Match(For) {
		return self.parseForStatement(parser, scan)
	} else if scan.Match(If) {
		return self.parseIfStatement(parser, scan)
	} else if scan.Match(LeftBrace) {
		return self.parseBlockStatement(parser, scan)
	}

	return self.parseExpressionStatement(parser, scan)
}
