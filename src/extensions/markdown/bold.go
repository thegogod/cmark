package markdown

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseBold(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseBold(parser, NewScanner(ptr))
}

func (self *Markdown) parseBold(parser html.Parser, scan *Scanner) (*html.StrongElement, error) {
	el := html.Strong()

	if !scan.MatchCount(Asterisk, 2) {
		return el, scan.Curr().Error("expected '**'")
	}

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

	log.Debugln("bold")
	return el, nil
}
