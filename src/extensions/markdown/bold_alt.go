package markdown

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseBoldAlt(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseBoldAlt(parser, NewScanner(ptr))
}

func (self *Markdown) parseBoldAlt(parser html.Parser, scan *_Scanner) (*html.StrongElement, error) {
	el := html.Strong()

	if !scan.MatchCount(Underscore, 2) {
		return el, scan.Curr().Error("expected '__'")
	}

	for !scan.MatchCount(Underscore, 2) {
		node, err := parser.ParseInline(scan.ptr)

		if node == nil {
			return el, scan.Curr().Error("expected closing '__'")
		}

		if err != nil {
			return el, err
		}

		el.Push(node)
	}

	log.Debugln("bold_alt")
	return el, nil
}
