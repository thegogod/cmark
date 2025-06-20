package flow

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Flow) ParseComparisonExpression(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	node, err := self.parseComparisonExpression(parser, NewScanner(ptr))

	if err != nil {
		return nil, err
	}

	return Html(node), nil
}

func (self *Flow) parseComparisonExpression(parser html.Parser, scan *Scanner) (Expression, error) {
	expr, err := self.parseTermExpression(parser, scan)

	if err != nil {
		return nil, err
	}

	for scan.Match(Gt) || scan.Match(GtEq) || scan.Match(Lt) || scan.Match(LtEq) {
		op := scan.Prev()
		right, err := self.parseTermExpression(parser, scan)

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
