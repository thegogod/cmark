package markdown

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseH4(parser ast.Parser, ptr *tokens.Pointer) (ast.Node, error) {
	return self.parseH4(parser, NewScanner(ptr))
}

func (self *Markdown) parseH4(parser ast.Parser, scan *_Scanner) (*html.HeadingElement, error) {
	if !scan.MatchCount(Hash, 4) || !scan.Match(Space) {
		return nil, scan.Curr().Error("expected '#### '")
	}

	heading := html.H4()

	for scan.Curr().Kind() != Eof && scan.Curr().Kind() != NewLine {
		node, err := parser.ParseInline(scan.ptr)

		if node == nil || err != nil {
			return heading, err
		}

		heading.Push(node)
	}

	return heading, nil
}
