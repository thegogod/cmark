package markdown

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseHighlight(parser ast.Parser, ptr *tokens.Pointer) (ast.Node, error) {
	return self.parseHighlight(parser, NewScanner(ptr))
}

func (self *Markdown) parseHighlight(parser ast.Parser, scan *_Scanner) (*html.MarkElement, error) {
	if !scan.MatchCount(EqualsEquals, 1) {
		return nil, scan.Curr().Error("expected '=='")
	}

	mark := html.Mark()

	for !scan.MatchCount(EqualsEquals, 1) {
		node, err := parser.ParseInline(scan.ptr)

		if node == nil {
			return mark, scan.Curr().Error("expected closing '=='")
		}

		if err != nil {
			return mark, err
		}

		mark.Push(node)
	}

	return mark, nil
}
