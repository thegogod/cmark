package flow

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Flow) ParseTermExpression(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	node, err := self.parseTermExpression(parser, NewScanner(ptr))

	if err != nil {
		return nil, err
	}

	return Html(node), nil
}

func (self *Flow) parseTermExpression(parser html.Parser, scan *Scanner) (Expression, error) {
	expr, err := self.parseFactorExpression(parser, scan)

	if err != nil {
		return nil, err
	}

	for scan.Match(Plus) || scan.Match(Minus) {
		op := scan.Prev()
		right, err := self.parseFactorExpression(parser, scan)

		if err != nil {
			return nil, err
		}

		expr = BinaryExpression{
			left:  expr,
			op:    op,
			right: right,
		}
	}

	return expr, nil
}
