package markdown

import (
	"github.com/thegogod/cmark/html"
	"github.com/thegogod/cmark/tokens"
)

func (self *Markdown) ParseHorizontalRule(parser html.Parser, ptr *tokens.Pointer) (html.Node, error) {
	return self.parseHorizontalRule(parser, NewScanner(ptr))
}

func (self *Markdown) parseHorizontalRule(_ html.Parser, scan *Scanner) (*html.HorizontalRuleElement, error) {
	el := html.Hr()

	if !scan.MatchCount(Dash, 3) {
		return el, scan.Curr().Error("expected '---'")
	}

	log.Debugln("horizontal_rule")
	return el, nil
}
