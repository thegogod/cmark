package flow

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Flow) ParseFactorExpression(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	node, err := self.parseFactorExpression(parser, NewScanner(ptr))

	if err != nil {
		return nil, err
	}

	return Html(node), nil
}

func (self *Flow) parseFactorExpression(parser html.Parser, scan *Scanner) (Expression, error) {
	expr, err := self.parseUnaryExpression(parser, scan)

	if err != nil {
		return nil, err
	}

	for scan.Match(Star) || scan.Match(Slash) {
		op := scan.Prev()
		right, err := self.parseUnaryExpression(parser, scan)

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
