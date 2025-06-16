package markdown

import (
	"github.com/thegogod/cmark/ast"
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseStrike(parser ast.Parser, ptr *tokens.Pointer) (ast.Node, error) {
	return self.parseStrike(parser, NewScanner(ptr))
}

func (self *Markdown) parseStrike(parser ast.Parser, scan *_Scanner) (*html.StrikeElement, error) {
	if !scan.MatchCount(Tilde, 1) {
		return nil, scan.curr.Error("expected '~'")
	}

	strike := html.S()

	for !scan.Match(Tilde) {
		node, err := parser.ParseInline(scan.ptr)

		if node == nil {
			return strike, scan.curr.Error("expected closing '~'")
		}

		if err != nil {
			return strike, err
		}

		strike.Push(node)
	}

	return strike, nil
}
