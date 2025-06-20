package flow

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Flow) ParseOrExpression(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	node, err := self.parseOrExpression(parser, NewScanner(ptr))

	if err != nil {
		return nil, err
	}

	return Html(node), nil
}

func (self *Flow) parseOrExpression(parser html.Parser, scan *Scanner) (Expression, error) {
	expr, err := self.parseAndExpression(parser, scan)

	if err != nil {
		return nil, err
	}

	for scan.Match(Or) {
		op := scan.Prev()
		right, err := self.parseAndExpression(parser, scan)

		if err != nil {
			return nil, err
		}

		expr = LogicalExpression{
			left:  expr,
			op:    op,
			right: right,
		}
	}

	return expr, nil
}
