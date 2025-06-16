package markdown

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseItalicAlt(parser ast.Parser, ptr *tokens.Pointer) (ast.Node, error) {
	return self.parseItalicAlt(parser, NewScanner(ptr))
}

func (self *Markdown) parseItalicAlt(parser ast.Parser, scan *_Scanner) (*html.ItalicElement, error) {
	if !scan.MatchCount(Underscore, 1) {
		return nil, scan.curr.Error("expected '_'")
	}

	italic := html.I()

	for !scan.Match(Underscore) {
		node, err := parser.ParseInline(scan.ptr)

		if node == nil {
			return italic, scan.curr.Error("expected closing '_'")
		}

		if err != nil {
			return italic, err
		}

		italic.Push(node)
	}

	return italic, nil
}
