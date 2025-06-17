package markdown

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseHorizontalRule(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseHorizontalRule(parser, NewScanner(ptr))
}

func (self *Markdown) parseHorizontalRule(_ html.Parser, scan *_Scanner) (*html.HorizontalRuleElement, error) {
	if !scan.MatchCount(Dash, 3) {
		return nil, scan.Curr().Error("expected '---'")
	}

	return html.Hr(), nil
}
