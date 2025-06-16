package markdown

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseBoldAlt(parser ast.Parser, ptr *tokens.Pointer) (ast.Node, error) {
	return self.parseBoldAlt(parser, NewScanner(ptr))
}

func (self *Markdown) parseBoldAlt(parser ast.Parser, scan *_Scanner) (*html.StrongElement, error) {
	if !scan.MatchCount(Underscore, 2) {
		return nil, scan.curr.Error("expected '__'")
	}

	el := html.Strong()

	for !scan.MatchCount(Underscore, 2) {
		node, err := parser.ParseInline(scan.ptr)

		if node == nil {
			return el, scan.curr.Error("expected closing '__'")
		}

		if err != nil {
			return el, err
		}

		el.Push(node)
	}

	return el, nil
}
