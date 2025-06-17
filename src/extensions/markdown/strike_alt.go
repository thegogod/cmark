package markdown

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseStrikeAlt(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseStrikeAlt(parser, NewScanner(ptr))
}

func (self *Markdown) parseStrikeAlt(parser html.Parser, scan *_Scanner) (*html.StrikeElement, error) {
	strike := html.S()

	if !scan.MatchCount(Tilde, 2) {
		return strike, scan.Curr().Error("expected '~~'")
	}

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

	log.Debugln("strike_alt")
	return strike, nil
}
