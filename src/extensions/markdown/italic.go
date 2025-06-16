package markdown

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseItalic(parser ast.Parser, ptr *tokens.Pointer) (ast.Node, error) {
	return self.parseItalic(parser, NewScanner(ptr))
}

func (self *Markdown) parseItalic(parser ast.Parser, scan *_Scanner) (*html.ItalicElement, error) {
	if !scan.MatchCount(Asterisk, 1) {
		return nil, scan.curr.Error("expected '*'")
	}

	italic := html.I()

	for !scan.Match(Asterisk) {
		node, err := parser.ParseInline(scan.ptr)

		if node == nil {
			return italic, scan.curr.Error("expected closing '*'")
		}

		if err != nil {
			return italic, err
		}

		italic.Push(node)
	}

	return italic, nil
}
