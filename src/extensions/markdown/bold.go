package markdown

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseBold(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseBold(parser, NewScanner(ptr))
}

func (self *Markdown) parseBold(parser html.Parser, scan *_Scanner) (*html.StrongElement, error) {
	if !scan.MatchCount(Asterisk, 2) {
		return nil, scan.Curr().Error("expected '**'")
	}

	el := html.Strong()

	for !scan.MatchCount(Asterisk, 2) {
		node, err := parser.ParseInline(scan.ptr)

		if node == nil {
			return el, scan.Curr().Error("expected closing '**'")
		}

		if err != nil {
			return el, err
		}

		el.Push(node)
	}

	return el, nil
}
