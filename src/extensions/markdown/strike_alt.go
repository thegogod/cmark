package markdown

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseStrikeAlt(parser ast.Parser, ptr *tokens.Pointer) (ast.Node, error) {
	return self.parseStrikeAlt(parser, NewScanner(ptr))
}

func (self *Markdown) parseStrikeAlt(parser ast.Parser, scan *_Scanner) (*html.StrikeElement, error) {
	if !scan.MatchCount(Tilde, 2) {
		return nil, scan.curr.Error("expected '~~'")
	}

	strike := html.S()

	for !scan.MatchCount(Tilde, 2) {
		node, err := parser.ParseInline(scan.ptr)

		if node == nil {
			return strike, scan.curr.Error("expected closing '~~'")
		}

		if err != nil {
			return strike, err
		}

		strike.Push(node)
	}

	return strike, nil
}
