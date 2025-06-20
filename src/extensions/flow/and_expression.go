package flow

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Flow) ParseAndExpression(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	node, err := self.parseAndExpression(parser, NewScanner(ptr))

	if err != nil {
		return nil, err
	}

	return Html(node), nil
}

func (self *Flow) parseAndExpression(parser html.Parser, scan *Scanner) (Expression, error) {
	expr, err := self.parseEqualityExpression(parser, scan)

	if err != nil {
		return nil, err
	}

	for scan.Match(And) {
		op := scan.Prev()
		right, err := self.parseEqualityExpression(parser, scan)

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
