package markdown

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseStrike(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseStrike(parser, NewScanner(ptr))
}

func (self *Markdown) parseStrike(parser html.Parser, scan *_Scanner) (*html.StrikeElement, error) {
	if !scan.MatchCount(Tilde, 1) {
		return nil, scan.Curr().Error("expected '~'")
	}

	strike := html.S()

	for !scan.Match(Tilde) {
		node, err := parser.ParseInline(scan.ptr)

		if node == nil {
			return strike, scan.Curr().Error("expected closing '~'")
		}

		if err != nil {
			return strike, err
		}

		strike.Push(node)
	}

	return strike, nil
}
