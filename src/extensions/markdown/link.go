package markdown

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseLink(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseLink(parser, NewScanner(ptr))
}

func (self *Markdown) parseLink(parser html.Parser, scan *_Scanner) (*html.AnchorElement, error) {
	link := html.A()

	if !scan.Match(LeftBracket) {
		return link, scan.Curr().Error("expected '['")
	}

	for !scan.Match(RightBracket) {
		node, err := parser.ParseInline(scan.ptr)

		if node == nil || err != nil {
			return link, err
		}

		link.Push(node)
	}

	if _, err := scan.Consume(LeftParen, "expected '('"); err != nil {
		return link, err
	}

	node, err := self.parseTextUntil(RightParen, parser, scan)

	if node == nil || err != nil {
		return link, err
	}

	log.Debugln("link")
	link.WithHref(string(node))
	return link, nil
}
