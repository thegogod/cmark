package markdown

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseStrikeAlt(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseStrikeAlt(parser, NewScanner(ptr))
}

func (self *Markdown) parseStrikeAlt(parser html.Parser, scan *_Scanner) (*html.StrikeElement, error) {
	if !scan.MatchCount(Tilde, 2) {
		return nil, scan.Curr().Error("expected '~~'")
	}

	strike := html.S()

	for !scan.MatchCount(Tilde, 2) {
		node, err := parser.ParseInline(scan.ptr)

		if node == nil {
			return strike, scan.Curr().Error("expected closing '~~'")
		}

		if err != nil {
			return strike, err
		}

		strike.Push(node)
	}

	return strike, nil
}
