package markdown

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseH2(parser ast.Parser, ptr *tokens.Pointer) (ast.Node, error) {
	return self.parseH2(parser, NewScanner(ptr))
}

func (self *Markdown) parseH2(parser ast.Parser, scan *_Scanner) (*html.HeadingElement, error) {
	if !scan.MatchCount(Hash, 2) || !scan.Match(Space) {
		return nil, scan.curr.Error("expected '## '")
	}

	heading := html.H2()

	for scan.curr.Kind() != Eof && scan.curr.Kind() != NewLine {
		node, err := parser.ParseInline(scan.ptr)

		if node == nil || err != nil {
			return heading, err
		}

		heading.Push(node)
	}

	return heading, nil
}
