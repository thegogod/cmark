package flow

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Flow) ParseEqualityExpression(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	node, err := self.parseEqualityExpression(parser, NewScanner(ptr))

	if err != nil {
		return nil, err
	}

	return Html(node), nil
}

func (self *Flow) parseEqualityExpression(parser html.Parser, scan *Scanner) (Expression, error) {
	expr, err := self.parseComparisonExpression(parser, scan)

	if err != nil {
		return nil, err
	}

	for scan.Match(EqEq) || scan.Match(NotEq) {
		op := scan.Prev()
		right, err := self.parseComparisonExpression(parser, scan)

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
