package markdown

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseBold(parser ast.Parser, ptr *tokens.Pointer) (ast.Node, error) {
	return self.parseBold(parser, NewScanner(ptr))
}

func (self *Markdown) parseBold(parser ast.Parser, scan *_Scanner) (*html.StrongElement, error) {
	if !scan.MatchCount(Asterisk, 2) {
		return nil, scan.curr.Error("expected '**'")
	}

	el := html.Strong()

	for !scan.MatchCount(Asterisk, 2) {
		node, err := parser.ParseInline(scan.ptr)

		if node == nil {
			return el, scan.curr.Error("expected closing '**'")
		}

		if err != nil {
			return el, err
		}

		el.Push(node)
	}

	return el, nil
}
